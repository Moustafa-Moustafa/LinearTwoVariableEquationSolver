package main

import (
	"testing"
)

func TestSolveEquations(t *testing.T) {
	t.Run("Valid", validEquation)
	t.Run("Invalid", invalidEquation)
	t.Run("NoSolution", noSolutionEquation)
}

func validEquation(t *testing.T) {
	sol, _ := solveEquations("x+-y+6=2x-2", "-x=-y-3")

	if sol.FirstUnknown != 5.5 || sol.SecondUnknown != 2.5 {
		t.Fail()
	}
}

func invalidEquation(t *testing.T) {
	_, err := solveEquations("x+y=6", "x+y=3")

	if err == nil {
		t.Fail()
	}
}

func noSolutionEquation(t *testing.T) {
	_, err := solveEquations("x=3", "x=-3")

	if err == nil {
		t.Fail()
	}
}
