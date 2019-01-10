package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Moustafa-Moustafa/equation"

	"github.com/gorilla/mux"
)

// Solve the system of linear equations
func solve(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	sol, err := solveEquations(params["equation1"], params["equation2"])

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprintf(w, "Solution: x = %f, y = %f", sol.FirstUnknown, sol.SecondUnknown)
}

func solveEquations(firstEquation string, secondEquation string) (equation.Solution, error) {
	parser := *equation.NewParser()
	e1, err := parser.Parse(firstEquation)
	if err != nil {
		return equation.Solution{0, 0}, err
	}

	e2, err := parser.Parse(secondEquation)
	if err != nil {
		return equation.Solution{0, 0}, err
	}

	solver := *equation.NewSolver()
	sol, err := solver.Solve([]equation.Equation{e1, e2})
	if err != nil {
		return equation.Solution{0, 0}, err
	}

	return sol, nil
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/solve/{equation1},{equation2}", solve).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
