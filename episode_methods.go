package main

import (
	"github.com/mitchellh/mapstructure"
)

func GetEpisode(integer int) (*Episode, error) {

	options := map[string]interface{}{
		"endpoint": endpointEpisode,
		"params": map[string]int{
			"integer": integer,
		},
	}

	data, err := MakePetition(options)
	if err != nil {
		return &Episode{}, err
	}

	episode := new(Episode)

	if err := mapstructure.Decode(data, &episode); err != nil {
		return &Episode{}, err
	}

	return episode, nil
}

func GetEpisodesArray(integers []int) (*MultipleEpisodes, error) {

	options := map[string]interface{}{
		"endpoint": endpointEpisode,
		"integers": integers,
	}

	data, err := MakePetition(options)
	if err != nil {
		return &MultipleEpisodes{}, err
	}

	episodes := new(MultipleEpisodes)

	if err := mapstructure.Decode(data, &episodes); err != nil {
		return &MultipleEpisodes{}, err
	}

	return episodes, nil
}
