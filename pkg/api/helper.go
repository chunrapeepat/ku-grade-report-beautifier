package api

import (
	"strconv"
)

// CalculateGPA calculate GPA
// ref: http://gpacalculator.net/how-to-calculate-gpa/
func CalculateGPA(grade []*CourseInfo) (float32, error) {
	var sumCredit float32
	var sumScore float32
	for _, c := range grade {
		credit, err := strconv.Atoi(c.Credit)
		if err != nil {
			return 0, err
		}
		if c.Grade[0] >= 65 && c.Grade[0] <= 68 || c.Grade == "F" {
			sumCredit += float32(credit)
			sumScore += gradeToNumber(c.Grade) * float32(credit)
		}
	}
	return sumScore / sumCredit, nil
}

// A 4.0, B 3.0, C 2.0, D 1.0, F 0.0
func gradeToNumber(grade string) float32 {
	if grade == "F" || grade[0] < 65 || grade[0] > 68 {
		return 0.0
	}
	gradeNum := float32(69 - grade[0])
	if len(grade) == 2 {
		gradeNum += 0.5
	}
	return gradeNum
}
