package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	checkORSToken()

	flags := collectFlag()

	fmt.Println("start:", flags[0])
	fmt.Println("end:", flags[1])
	fmt.Println("method:", flags[2])

	consumptionRate, err := GetConsumption(flags[2])
	if err != nil {
		panic(err)
	}

	startCoord := getPointCoordinate(flags[0])
	endCoord := getPointCoordinate(flags[1])

	distance, err := GetDistance(startCoord, endCoord)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Your trip caused %.2fkg of CO2-equivalent\n", *distance*float64(consumptionRate)/1000)

}

func collectFlag() []string {
	startPtr := flag.String("start", "munich", "starting location")
	endPtr := flag.String("end", "munich", "ending location")
	carPtr := flag.String("transportation-method", "small-diesel-car", "way of transportation")

	flag.Parse()

	return []string{*startPtr, *endPtr, *carPtr}
}

func checkORSToken() {
	openRouterToken = os.Getenv("ORS_TOKEN")
	if openRouterToken == "" {
		panic("Env var ORS_TOKEN Not set")
	}
}

func getPointCoordinate(name string) []float64 {
	coord, err := GetCoordinate(name)

	if err != nil {
		panic(err)
	}

	return coord
}
