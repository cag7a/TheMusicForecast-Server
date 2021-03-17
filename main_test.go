package main

import (
	"fmt"
	"log"
	"math"
	"testing"

	"github.com/joho/godotenv"
)

func TestGeolocation(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	//1600 Pennsylvania Avenue NW, Washington, DC 20500
	//lat = 38.8977  lon = -77.0365
	lat, lon := GetLocation("1600 Pennsylvania Avenue NW, Washington, DC 20500")

	if math.Round(lat) != 39 {
		t.Error(fmt.Sprintf("FUCK, Wrong Lat.\nExpected: 39.\nGot: %v", int(lat)))
	}
	if math.Round(lon) != -77 {
		t.Error(fmt.Sprintf("FUCK, Wrong Lon.\nExpected: -77.\nGot: %v", int(lon)))
	}

	//The White House
	//lat = 38.8977  lon = -77.0365
	lat, lon = GetLocation("The White House")

	if math.Round(lat) != 39 {
		t.Error(fmt.Sprintf("FUCK, Wrong Lat.\nExpected: 39.\nGot: %v", int(lat)))
	}
	if math.Round(lon) != -77 {
		t.Error(fmt.Sprintf("FUCK, Wrong Lon.\nExpected: -77.\nGot: %v", int(lon)))
	}

}
