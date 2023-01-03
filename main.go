package main

import (
	"fmt"
	"go-elk/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func ping(w http.ResponseWriter, r *http.Request) {

	utils.CreateLog(&utils.Logs{
		Method:     r.Method,
		URL:        r.RequestURI,
		Event:      "FORWARD-LOG",
		StatusCode: 200,
		Message:    "Success forward log with elk",
	}, "info")

	fmt.Println("testing")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", ping).GetMethods()

	http.ListenAndServe(":6060", router)
}
