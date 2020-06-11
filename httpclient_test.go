package main

import (
	"encoding/json"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const coordinateValidResp = `{
		"features": [
			{
				"geometry": {
					"type": "Point",
					"coordinates": [
						11.544467,
						48.152126
					]
				}
			}
		]
	}`

const distanceValidResp = `{
		"distances": [
			[
				0.0,
				585.53
			],
			[
				585.53,
				0.0
			]
    	]
	}`

func TestGetCoordinate(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	//////////////////////////
	//test case 1, valid response
	//////////////////////////
	validResponse := new(PointInfoResponse)
	json.Unmarshal([]byte(coordinateValidResp), validResponse)
	resp, _ := httpmock.NewJsonResponder(200, validResponse)
	httpmock.RegisterResponder("GET", HOST+"/geocode/search",
		resp)

	coord, err := GetCoordinate("test")

	assert.NoError(t, err, "there should be no error")
	assert.Equal(t, []float64{11.544467, 48.152126}, coord)

	//////////////////////////
	//test case 2, server error
	//////////////////////////
	httpmock.RegisterResponder(http.MethodGet, HOST+"/geocode/search",
		httpmock.NewStringResponder(501, ""))

	_, err = GetCoordinate("test")

	assert.EqualError(t, err, "Cannot get coordinates of test", "expecting error here")

}

func TestGetDistance(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	//////////////////////////
	//test case 1, valid response
	//////////////////////////
	validResponse := new(DistanceResponse)
	json.Unmarshal([]byte(distanceValidResp), validResponse)
	resp, _ := httpmock.NewJsonResponder(200, validResponse)
	httpmock.RegisterResponder(http.MethodPost, HOST+"/v2/matrix/driving-car",
		resp)

	distance, err := GetDistance([]float64{1.0, 1.0}, []float64{2.0, 2.0})

	assert.NoError(t, err, "there should be no error")
	assert.Equal(t, 585.53, *distance)

	//////////////////////////
	//test case 2, server error
	//////////////////////////
	httpmock.RegisterResponder(http.MethodPost, HOST+"/v2/matrix/driving-car",
		httpmock.NewStringResponder(501, ""))

	_, err = GetDistance([]float64{1.0, 1.0}, []float64{2.0, 2.0})

	assert.EqualError(t, err, "Cannot get distance", "expecting error here")
}
