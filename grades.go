package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func formHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	score := r.FormValue("score")
	re := regexp.MustCompile(`^(\d)+(.)?(\d)*$`)
	val := re.Match([]byte(score))
	if val == true {
		res := getGrade(score)
		if res == "Invalid Input" {
			http.Error(w, "422 Invalid Input", http.StatusUnprocessableEntity)
			return
		}
		fmt.Fprintf(w, res)

	} else {
		http.Error(w, "422 Invalid Input", http.StatusUnprocessableEntity)
		return
	}
}

func getGrade(input string) string {

	sc, err := strconv.ParseFloat(input, 64)
	if err != nil || (sc > 100 || sc < 0) {
		return "Invalid Input"
	}
	if sc >= 90 {
		return ("You got an A!")
	} else if sc >= 80 {
		return ("You got a B.")
	} else if sc >= 70 {
		return ("You got a C.")
	} else if sc >= 60 {
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
