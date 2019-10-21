package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func getStartDate() string {
	dt := time.Now()
	dt = dt.AddDate(0, 0, -7)
	return dt.Format("2006-01-02")
}

func getEndDate() string {
	dt := time.Now()
	return dt.Format("2006-01-02")
}

func RequestAsteroidData() ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf(RestURL, getStartDate(), getEndDate(), APIKey))
	if err != nil {
		fmt.Println("Error getting asteroid data")
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error parsing asteroid data")
		return []byte{}, err
	}

	return body, nil
}

func GetAsteroidData() {
	strData, err := RequestAsteroidData()
	if err != nil {
		panic(err)
	}
	err = ParseAsteroidData(strData)
	if err != nil {
		panic(err)
	}
}

type dataJSON struct {
	Objects map[string]*json.RawMessage `json:"near_earth_objects"`
}

func ParseAsteroidData(data []byte) error {
	var objmap dataJSON
	err := json.Unmarshal(data, &objmap)
	if err != nil {
		return err
	}

	var asteroids []asteroidData
	for key := range objmap.Objects {
		err = json.Unmarshal(*objmap.Objects[key], &asteroids)
		if err != nil {
			fmt.Println(err)
			continue
		}
		AsteroidData.Objects = append(AsteroidData.Objects, asteroids...)
	}
	return nil
}
