package main

import (
	"fmt"
)

type GradeCalculator struct {
	grades map[string]float64
	averageGrade float64
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		grades: make(map[string]float64),
	}
}

func (g *GradeCalculator) AcceptGrades(){
	var n int
	var subjectName string
	var subjectGrade int

	for {
		fmt.Print("Enter number of Subjects: ")
		_, err := fmt.Scan(&n)
		if err != nil {
			fmt.Println("Invalid Input. Please insert a valid number of subjects.")
			continue
		}
		break
	}

	for i := 0; i < n; i++ {
		for {
			fmt.Printf("Enter Subject%d name: ", i + 1)
			_, subjectErr := fmt.Scan(&subjectName)
			if subjectErr != nil {
				fmt.Println("Invalid Subject Name. Please enter the subject again.")
				continue
			}
			break
		}

		for {
			fmt.Printf("Enter Subject%d grade: ", i + 1)
			_ , gradeErr := fmt.Scan(&subjectGrade)
			if gradeErr != nil || subjectGrade < 0 || subjectGrade > 100 {
				fmt.Println("Invalid Subject Grade.Please enter a number between 0 and 100.")
				continue
			}
			break
		}
		
		g.grades[subjectName] = float64(subjectGrade)
	}
}

func (g *GradeCalculator) CalculateGrade() {
	n := float64(len(g.grades))
	var total float64 = 0

	for _, grade := range g.grades {
		total += grade
	}

	g.averageGrade = total / n
}


func main() {
	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AcceptGrades()
	gradeCalculator.CalculateGrade()
	fmt.Printf("Your Average Grade is: %v", gradeCalculator.averageGrade);
}