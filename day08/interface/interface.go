package main

import (
	"fmt"
)

type student interface {
	getStudentDetails() string
	getGPAX() float64
}

type highschoolStudent struct {
	name   string
	school string
	grade  int
	gpax   float64
}

func (s highschoolStudent) getStudentDetails() string {
	return fmt.Sprintf("Name: %s, Grade: %d, %s School", s.name, s.grade, s.school)
}

func (s highschoolStudent) getGPAX() float64 {
	return s.gpax
}

type universityStudent struct {
	name       string
	university string
	year       string
	major      string
	gpax       float64
}

func (s universityStudent) getStudentDetails() string {
	return fmt.Sprintf("Name: %s, %s Year %s, Major, %s University", s.name, s.year, s.major, s.university)
}

func (s universityStudent) getGPAX() float64 {
	return s.gpax
}

func printDetails(s student) {
	fmt.Println(s.getStudentDetails(), "\nGPAX:", s.getGPAX())
}

func main() {
	// If struct has all of the method required by interface,
	// it is that interface!

	h := highschoolStudent{"Bob", "Bobby", 12, 3.42}
	printDetails(h)

	u := universityStudent{"Pawat", "Kasetsart", "Sophomore", "Software Engineering", 3.88}
	printDetails(u)
}
