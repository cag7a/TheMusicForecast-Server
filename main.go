package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

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

type Geolocation struct {
	Authenticationresultcode string `json:"authenticationResultCode"`
	Brandlogouri             string `json:"brandLogoUri"`
	Copyright                string `json:"copyright"`
	Resourcesets             []struct {
		Estimatedtotal int `json:"estimatedTotal"`
		Resources      []struct {
			Type  string    `json:"__type"`
			Bbox  []float64 `json:"bbox"`
			Name  string    `json:"name"`
			Point struct {
				Type        string    `json:"type"`
				Coordinates []float64 `json:"coordinates"`
			} `json:"point"`
			Address struct {
				Admindistrict    string `json:"adminDistrict"`
				Admindistrict2   string `json:"adminDistrict2"`
				Countryregion    string `json:"countryRegion"`
				Formattedaddress string `json:"formattedAddress"`
				Locality         string `json:"locality"`
			} `json:"address"`
			Confidence    string `json:"confidence"`
			Entitytype    string `json:"entityType"`
			Geocodepoints []struct {
				Type              string    `json:"type"`
				Coordinates       []float64 `json:"coordinates"`
				Calculationmethod string    `json:"calculationMethod"`
				Usagetypes        []string  `json:"usageTypes"`
			} `json:"geocodePoints"`
			Matchcodes []string `json:"matchCodes"`
		} `json:"resources"`
	} `json:"resourceSets"`
	Statuscode        int       `json:"statusCode"`
	Statusdescription string    `json:"statusDescription"`
	Traceid           time.Time `json:"traceId"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
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
	WeatherAPI(35.8486, -86.3649)

	http.HandleFunc("/api/playlist", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		search := query.Get("search")
		fmt.Println(search)
		w.Write([]byte(`{"playlist": "37i9dQZF1DX2Nc3B70tvx0", "weather": "rain"}`))

		GetLocation(search)
	})

	fmt.Println("Now forecasting on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// Gets a location using a string
// Returns latitude and longitude as float64
func GetLocation(location string) (float64, float64) {
	// API call to the Bing Maps API
	resp, err := http.Get(fmt.Sprintf("http://dev.virtualearth.net/REST/v1/Locations/query=%v?maxResults=1&key=%v", location, os.Getenv("BING_MAPS_KEY")))

	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var geolocation Geolocation
	json.Unmarshal(bodyBytes, &geolocation)
	return geolocation.Resourcesets[0].Resources[0].Point.Coordinates[0], geolocation.Resourcesets[0].Resources[0].Point.Coordinates[1]

}

func WeatherAPI(lat float64, lon float64) {
	fmt.Printf("WEATHER_API_KEY = %v\n", os.Getenv("WEATHER_API_KEY"))
	fmt.Printf("http://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%v\n", lat, lon, os.Getenv("WEATHER_API_KEY"))
	resp, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%v&units=imperial", lat, lon, os.Getenv("WEATHER_API_KEY")))
	if err != nil {
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
