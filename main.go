package main

import (
	"log"
	"net/http"
)

func main() {

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

	http.HandleFunc("/api/playlist", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		lat := query.Get("lat")
		long := query.Get("long")
		w.WriteHeader(200)
		w.Write([]byte(lat + long))
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
