package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func formHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	score := r.FormValue("score")
	sc, err := strconv.ParseFloat(score, 64)
	if err != nil || (sc > 100 || sc < 0) {
		http.Error(w, "422 Invalid Input", http.StatusUnprocessableEntity)
		return
	}
	res := getGrade(sc)

	fmt.Fprintf(w, res)

}

func getGrade(input float64) string {
	if input >= 90 {
		return ("You got an A!")
	} else if input >= 80 {
		return ("You got a B.")
	} else if input >= 70 {
		return ("You got a C.")
	} else if input >= 60 {
		return ("You got a D.")
	} else {
		return ("You got an E.")
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at port 8086\n")
	err := http.ListenAndServe(":8086", nil)
	if err != nil {
		log.Fatal(err)
	}

}
