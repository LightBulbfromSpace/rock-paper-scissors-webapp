package main

import (
	"encoding/json"
	rps_web "github.com/LightBulbfromSpace/rock-paper-scissors-webapp/rps_web"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func playRound(w http.ResponseWriter, r *http.Request) {

	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps_web.PlayRound(playerChoice)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func main() {
	http.HandleFunc("/play", playRound)
	http.HandleFunc("/", homePage)

	http.ListenAndServe("localhost:8080", nil)
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
