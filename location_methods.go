package main

import (
	"github.com/mitchellh/mapstructure"
)

func GetLocation(integer int) (*Location, error) {

	options := map[string]interface{}{
		"endpoint": endpointLocation,
		"params": map[string]int{
			"integer": integer,
		},
	}

	data, err := MakePetition(options)
	if err != nil {
		return &Location{}, err
	}

	location := new(Location)

	if err := mapstructure.Decode(data, &location); err != nil {
		return &Location{}, err
	}

	return location, nil
}

func GetLocationsArray(integers []int) (*MultipleLocations, error) {

	options := map[string]interface{}{
		"endpoint": endpointLocation,
		"integers": integers,
	}

	data, err := MakePetition(options)
	if err != nil {
		return &MultipleLocations{}, err
	}

	locations := new(MultipleLocations)

	if err := mapstructure.Decode(data, &locations); err != nil {
		return &MultipleLocations{}, err
	}

	return locations, nil
}
