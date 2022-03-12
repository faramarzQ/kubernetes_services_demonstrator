package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Listening ...")
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Called!")
	})
	http.ListenAndServe(":"+os.Getenv("GO_APP_LISTENER_CLUSTERIP_SERVICE_PORT"), nil)
}
