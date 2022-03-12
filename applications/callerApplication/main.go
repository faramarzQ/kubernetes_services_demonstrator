package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	fmt.Println("Running caller server...")
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		listenerHost := os.Getenv("GO_APP_LISTENER_CLUSTERIP_SERVICE_HOST")
		listenerPort := os.Getenv("GO_APP_LISTENER_CLUSTERIP_SERVICE_PORT")
		listenerAddr := "http://" + listenerHost + ":" + listenerPort

		fmt.Printf("Calling %v\n" + listenerAddr)

		_, err := http.Get(listenerAddr)
		if err != nil {
			log.Fatal(err)
		}

		w.Write([]byte("Called."))
	})
	http.ListenAndServe(":"+os.Getenv("GO_APP_CALLER_CLUSTERIP_SERVICE_PORT"), nil)

}
