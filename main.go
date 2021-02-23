package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "The Music Forecast")
		net := New(4, 4, 4, 4, 0.005, Sigmoid, SigmoidPrime)

		inData := [][]float64{{0.1, 0.4, 0.6, 0.5}, {0.1, 0.4, 0.6, 0.5}, {0.1, 0.4, 0.6, 0.5}, {0.1, 0.4, 0.6, 0.5}, {0.1, 0.4, 0.6, 0.5}}
		targets := [][]float64{{0.1, 0.4, 0.6, 0.5}, {0.1, 0.4, 0.6, 0.5}, {0.1, 0.4, 0.6, 0.5}, {0.1, 0.4, 0.6, 0.5}, {0.1, 0.4, 0.6, 0.5}}

		rand.Seed(time.Now().UTC().UnixNano())

		for epochs := 0; epochs < 5000; epochs++ {
			// Train the net on each example
			for i, ix := range inData {
				net.Train(ix, targets[i])
			}
		}

		pred := net.Predict(inData[0])

		matrixPrint(pred)
		fmt.Println(pred.Dims())

		// http.ServeFile(w, r, 'client/index.html')
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
