package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Moustafa-Moustafa/equation"

	"github.com/gorilla/mux"
)

// Solve the linear equation
func Solve(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	parser := *equation.NewParser()
	e1 := parser.Parse(params["equation1"])
	e2 := parser.Parse(params["equation1"])

	solver := *equation.NewSolver()
	sol := solver.Solve([]equation.Equation{e1, e2})

	fmt.Fprintf(w, "Solution: x = %f, y = %f\n", sol.FirstUnknown, sol.SecondUnknown)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/solve/{equation1},{equation2}", Solve).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
