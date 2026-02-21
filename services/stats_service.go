package services

import "student-grade-api/models"

// CalculateCourseAverage returns average grade for a course
func CalculateCourseAverage(grades []models.Grade) float64 {
    if len(grades) == 0 {
        return 0
    }

    var total float64
    for _, g := range grades {
        total += g.Grade
    }

    return total / float64(len(grades))
}