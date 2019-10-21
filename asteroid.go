package main

import (
	"net/http"
)

func asteroidHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/asteroid/"):]
	for _, asteroid := range AsteroidData.Objects {
		if asteroid.ID == id {
			ServeTemplate(w, "asteroid.html", asteroid)
			return
		}
	}
	w.Write([]byte("asteroid not found"))
}

func mapHandler(w http.ResponseWriter, r *http.Request) {	
	id := r.URL.Path[len("/map/"):]
	for _, asteroid := range AsteroidData.Objects {
		if asteroid.ID == id {
			ServeTemplate(w, "map.html", asteroid)
			return
		}
	}
	w.Write([]byte("asteroid not found"))
}
