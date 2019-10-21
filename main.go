package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/subosito/gotenv"
)

var (
	APIKey		 string
	RestURL      = "https://api.nasa.gov/neo/rest/v1/feed?detailed=true&start_date=%s&end_date=%s&api_key=%s"
	AsteroidData responseData
)

func ServeTemplate(w http.ResponseWriter, filename string, data interface{}) {
	t, err := template.ParseFiles(filepath.Join("templates", filename))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	gotenv.Load()
	APIKey = os.Getenv("NASA_API_KEY")
	GetAsteroidData()

	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/list/size", listSizeHandler)
	http.HandleFunc("/list/dangerous", listDangerousHandler)
	http.HandleFunc("/list/speed", listSpeedHandler)
	http.HandleFunc("/asteroid/", asteroidHandler)
	http.HandleFunc("/map/", mapHandler)
	http.Handle("/", http.FileServer(http.Dir("public")))
	//http.ListenAndServe("192.168.30.92:8080", nil)
	http.ListenAndServe("localhost:8080", nil)
}
