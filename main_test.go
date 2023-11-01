package main

import ("testing")

func TestStudentGrades(t *testing.T) {
	studentGrades := map[string]float64{
		"John": 85.5,
		"Jane": 90.0,
		"Mike": 76.5,
	}

	johnGrade := studentGrades["John"]
	janeGrade := studentGrades["Jane"]
	mikeGrade := studentGrades["Mike"]

	if johnGrade != 85.5 {
		t.Errorf("Expected John's grade to be 85.5, but got %.2f", johnGrade)
	}

	if janeGrade != 90.0 {
		t.Errorf("Expected Jane's grade to be 90.0, but got %.2f", janeGrade)
	}

	if mikeGrade != 76.5 {
		t.Errorf("Expected Mike's grade to be 76.5, but got %.2f", mikeGrade)
	}
}