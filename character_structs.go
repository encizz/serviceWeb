package main

import (
	"github.com/mitchellh/mapstructure"
)

func GetCharacter(integer int) (*Character, error) {

	options := map[string]interface{}{
		"endpoint": endpointCharacter,
		"params": map[string]int{
			"integer": integer,
		},
	}

	data, err := MakePetition(options)
	if err != nil {
		return &Character{}, err
	}

	character := new(Character)

	if err := mapstructure.Decode(data, &character); err != nil {
		return &Character{}, err
	}

	return character, nil
}

func GetCharactersArray(integers []int) (*MultipleCharacters, error) {

	options := map[string]interface{}{
		"endpoint": endpointCharacter,
		"integers": integers,
	}

	data, err := MakePetition(options)
	if err != nil {
		return &MultipleCharacters{}, err
	}

	characters := new(MultipleCharacters)

	if err := mapstructure.Decode(data, &characters); err != nil {
		return &MultipleCharacters{}, err
	}

	return characters, nil
}
