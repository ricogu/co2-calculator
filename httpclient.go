package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

const HOST = "https://api.openrouteservice.org"

var openRouterToken string

func GetCoordinate(city string) ([]float64, error) {
	client := &http.Client{}
	pointInfo := new(PointInfoResponse)
	req, err := http.NewRequest(http.MethodGet, HOST+"/geocode/search", nil)

	if err != nil {
		return nil, err
	}

	//add the query and header to request
	q := req.URL.Query()
	q.Add("api_key", openRouterToken)
	q.Add("text", city)
	q.Add("layers", "locality")
	q.Add("size", "1")
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("Cannot get coordinates of " + city)
	}

	//decode response
	json.NewDecoder(resp.Body).Decode(pointInfo)

	//return coordinate
	return pointInfo.Features[0].Geometry.Coordinates, nil
}

func GetDistance(start []float64, end []float64) (*float64, error) {
	client := &http.Client{}

	distanceResponse := new(DistanceResponse)

	//construct request body
	distanceRequest := DistanceRequest{
		Locations: [][]float64{start, end},
		Metrics:   []string{"distance"},
		Units:     "km",
	}

	requestBody, err := json.Marshal(distanceRequest)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, HOST+"/v2/matrix/driving-car", bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, err
	}

	//add headers
	req.Header.Add("Authorization", openRouterToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("Cannot get distance")
	}

	//decode response
	json.NewDecoder(resp.Body).Decode(distanceResponse)

	return &distanceResponse.Distances[0][1], nil
}
