package main

import (
	"encoding/json"
	"log"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

//Creation of Struct
type BusStop struct {
	BusStopCode  string `json:"BusStopCode"`
	Description string `json:"Description"`
}

var busStops map[string]BusStop

func busStop(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-type") == "application/json" {

		if r.Method == "POST" {

			var newBusStop BusStop
			reqBody, err := ioutil.ReadAll(r.Body)
			if err == nil{
				json.Unmarshal(reqBody, &newBusStop)
				fmt.Println(newBusStop)
				if newBusStop.BusStopCode != ""{
					busStops[newBusStop.BusStopCode] = newBusStop
					w.WriteHeader(201)
					w.Write([]byte(
                        "201"))
					return
				}else{
					w.WriteHeader(404)
					w.Write([]byte(
                        "404"))
					return
				}
			}
			return
		}

		if r.Method == "PUT" {
			var newBusStop BusStop
			reqBody, err := ioutil.ReadAll(r.Body)
			if err == nil{
				json.Unmarshal(reqBody, &newBusStop)
				fmt.Println(newBusStop)
				if _, ok := busStops[newBusStop.BusStopCode]; !ok {
                    busStops[newBusStop.BusStopCode] = newBusStop
                    w.WriteHeader(http.StatusCreated)
                    w.Write([]byte("201 - BusStop added: " + newBusStop.BusStopCode))
                } else {
                    // update course
                    busStops[newBusStop.BusStopCode] = newBusStop
                    w.WriteHeader(http.StatusAccepted)
					w.Write([]byte("202 - BusStop updated: " + newBusStop.BusStopCode))
				}
			return
		}
	}
	}
	return

}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/BusStops", busStop).Methods(
		"PUT", "POST")

	fmt.Println("Listening at port 5040")
	log.Fatal(http.ListenAndServe(":5040", router))

}
