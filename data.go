package main

type PointInfoResponse struct {
	Features []struct {
		Geometry struct {
			Coordinates []float64 `json:"coordinates"`
			Type        string    `json:"type"`
		} `json:"geometry"`
	} `json:"features"`
}

type DistanceRequest struct {
	Locations [][]float64 `json:"locations"`
	Metrics   []string    `json:"metrics"`
	Units     string      `json:"units"`
}

type DistanceResponse struct {
	Distances [][]float64 `json:"distances"`
}
