package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"os"
	"github.com/joho/godotenv"
)

type Weather struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil{
		fmt.Print(err.Error())
	}
	http.Handle("/", http.FileServer(http.Dir("client/build")))

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	net := New(4, 4, 4, 4, 0.005, Sigmoid, SigmoidPrime)

	// 	inData := [][]float64{{0.1, 0.4, 0.6, 0.5}, {0.1, 0.4, 0.6, 0.5}, {0.1, 0.4, 0.6, 0.5}, {0.1, 0.4, 0.6, 0.5}, {0.1, 0.4, 0.6, 0.5}}
	// 	targets := [][]float64{{0.1, 0.4, 0.6, 0.5}, {0.1, 0.4, 0.6, 0.5}, {0.1, 0.4, 0.6, 0.5}, {0.1, 0.4, 0.6, 0.5}, {0.1, 0.4, 0.6, 0.5}}

	// 	rand.Seed(time.Now().UTC().UnixNano())

	// 	for epochs := 0; epochs < 5000; epochs++ {
	// 		// Train the net on each example
	// 		for i, ix := range inData {
	// 			net.Train(ix, targets[i])
	// 		}
	// 	}

	// 	pred := net.Predict(inData[0])

	// 	matrixPrint(pred)
	// 	fmt.Println(pred.Dims())

	// 	http.ServeFile(w, r, "client/build/index.html")
	// })
	WeatherAPI(35.8486, -86.3649);

	http.HandleFunc("/api/playlist", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		search := query.Get("search")
		fmt.Println(search)
		w.Write([]byte(`{"playlist": "37i9dQZF1DX4dyzvuaRJ0n"}`))
	})

	fmt.Println("Now forecasting on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}


func WeatherAPI(lat float64, lon float64){
	fmt.Printf("WEATHER_API_KEY = %v\n", os.Getenv("WEATHER_API_KEY"))
	fmt.Printf("http://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%v", lat, lon, os.Getenv("WEATHER_API_KEY"))
	resp, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%v&units=imperial", lat, lon, os.Getenv("WEATHER_API_KEY")))
	if err != nil{
		fmt.Print(err.Error())
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var weatherData Weather
 	json.Unmarshal(bodyBytes, &weatherData)
 	fmt.Printf("API Response as struct %+v\n", weatherData)
}