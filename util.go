package service

import (
	"encoding/json"
	"io/ioutil"
)

// a data structure is a collection of related values, the relationship among them,
// and the functions or operations that can be applied to the data.

type Data struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	Metascore  string
	imdbRating string
	imdbVotes  string
	imdbID     string
	Type       string
	Images     []string
}

// Payload structure refers to the way data is organized and formatted when it is sent over a network.

type payload struct {
	Data []Data
}

func raw() ([]Data, error) {
	r, err := ioutil.ReadFile("data.json")
	if err != nil {
		return nil, err
	}

	// here we creae the new variable called payload of type payload with structure

	var payload payload
	err = json.Unmarshal(r, &payload.Data)
	if err != nil {
		return nil, err
	}
	return payload.Data, nil

}

func GetAll() ([]Data, error) {
	data, err := raw()
	if err != nil {
		return nil, err
	}
	return data, nil

}

func GetById(idx int) (any, error) {
	data, err := raw()
	if err != nil {
		return nil, err
	}
	if idx > len(data) {
		res := make([]string, 0)
		return res, nil
	}

	return data[idx], nil

}
