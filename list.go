package main

import (
	"net/http"
	"sort"
)

func listHandler(w http.ResponseWriter, r *http.Request) {
	ServeTemplate(w, "list.html", AsteroidData)
}

func listSpeedHandler(w http.ResponseWriter, r *http.Request) {
	var asteroidVelocity []asteroidData
	for _, asteroid := range AsteroidData.Objects {
		if len(asteroid.ApproachData) > 0 {
			asteroidVelocity = append(asteroidVelocity, asteroid)
		}
	}
	sort.Slice(asteroidVelocity, func(i, j int) bool {
		return asteroidVelocity[i].ApproachData[0].RelativeVelocity.KilometersPerSecond >
			asteroidVelocity[j].ApproachData[0].RelativeVelocity.KilometersPerSecond
	})
	ServeTemplate(w, "list.html", responseData{asteroidVelocity})
}

func listSizeHandler(w http.ResponseWriter, r *http.Request) {
	var asteroidVelocity []asteroidData
	for _, asteroid := range AsteroidData.Objects {
		if asteroid.Diameter.Kilometers.Min > 0 {
			asteroidVelocity = append(asteroidVelocity, asteroid)
		}
	}
	sort.Slice(asteroidVelocity, func(i, j int) bool {
		if asteroidVelocity[i].Diameter.Kilometers.Min >
			asteroidVelocity[j].Diameter.Kilometers.Min {
			return true
		}
		return asteroidVelocity[i].Diameter.Kilometers.Max >
			asteroidVelocity[j].Diameter.Kilometers.Max
	})
	ServeTemplate(w, "list.html", responseData{asteroidVelocity})
}

func listDangerousHandler(w http.ResponseWriter, r *http.Request) {
	var asteroidVelocity []asteroidData
	for _, asteroid := range AsteroidData.Objects {
		if asteroid.IsHazardous {
			asteroidVelocity = append(asteroidVelocity, asteroid)
		}
	}
	ServeTemplate(w, "list.html", responseData{asteroidVelocity})
}
