package main

import (
	"errors"
	"fmt"
)

var consumptionMap map[string]int

func init() {
	consumptionMap = make(map[string]int)
	//small cars
	consumptionMap["small-diesel-car"] = 142
	consumptionMap["small-petrol-car"] = 154
	consumptionMap["small-plugin-hybrid-car"] = 73
	consumptionMap["small-electric-car"] = 50

	//medium cars
	consumptionMap["medium-diesel-car"] = 171
	consumptionMap["medium-petrol-car"] = 192
	consumptionMap["medium-plugin-hybrid-car"] = 110
	consumptionMap["medium-electric-car"] = 58

	//large cars
	consumptionMap["large-diesel-car"] = 209
	consumptionMap["large-petrol-car"] = 282
	consumptionMap["large-plugin-hybrid-car"] = 126
	consumptionMap["large-electric-car"] = 73

	consumptionMap["bus"] = 27
	consumptionMap["train"] = 6
}

func GetConsumption(transporation string) (int, error) {
	consumption := consumptionMap[transporation]
	if consumption == 0 {
		return 0, errors.New(fmt.Sprintf("transportation method %s does not exist", transporation))
	}
	return consumption, nil
}
