package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/montybeatnik/lab"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpls, err := template.ParseFiles("./cmd/web/templates/home.html")
	if err != nil {
		log.Println(err)
	}
	if err := tmpls.ExecuteTemplate(w, "home", nil); err != nil {
		log.Println(err)
	}
}

func handleDevices(w http.ResponseWriter, r *http.Request) {
	tmpls, err := template.ParseFiles("./cmd/web/templates/devices.html")
	if err != nil {
		log.Println(err)
	}
	store := lab.NewStore()
	devices, err := store.GetDevices()
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, "error")
		return
	}
	if err := tmpls.ExecuteTemplate(w, "devices", devices); err != nil {
		log.Println(err)
	}
}

func main() {

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/devices", handleDevices)
	http.ListenAndServe(":8080", nil)
}
