package education

import (
	"testing"
	"time"
)

func TestNewGraduationCalculator(t *testing.T) {
	birthDate := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	calculator := NewGraduationCalculator(birthDate)

	if calculator.BirthYear != 1990 {
		t.Errorf("BirthYear = %v, 期待値 1990", calculator.BirthYear)
	}

	if len(calculator.Levels) != 8 {
		t.Errorf("Levels の数 = %v, 期待値 8", len(calculator.Levels))
	}
}

func TestGraduationCalculations(t *testing.T) {
	birthDate := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	calculator := NewGraduationCalculator(birthDate)

	tests := []struct {
		name     string
		method   func() int
		expected int
	}{
		{"小学校卒業", calculator.GetElementaryGraduation, 2002},
		{"中学校卒業", calculator.GetMiddleSchoolGraduation, 2005},
		{"高校卒業", calculator.GetHighSchoolGraduation, 2008},
		{"高専卒業", calculator.GetTechnicalCollegeGraduation, 2010},
		{"短大卒業", calculator.GetJuniorCollegeGraduation, 2010},
		{"専門学校卒業", calculator.GetVocationalSchoolGraduation, 2010},
		{"大学卒業", calculator.GetUniversityGraduation, 2012},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.method()
			if result != tt.expected {
				t.Errorf("%s = %v, 期待値 %v", tt.name, result, tt.expected)
			}
		})
	}
}

func TestGetAllGraduations(t *testing.T) {
	birthDate := time.Date(2000, 4, 1, 0, 0, 0, 0, time.UTC)
	calculator := NewGraduationCalculator(birthDate)

	levels := calculator.GetAllGraduations()

	expectedGraduations := map[string]int{
		"小学校":       2012,
		"中学校":       2015,
		"高校":        2018,
		"高専":        2020,
		"短大":        2020,
		"専門学校（2年制）": 2020,
		"専門学校（3年制）": 2021,
		"大学":        2022,
	}

	for _, level := range levels {
		expected, exists := expectedGraduations[level.Name]
		if !exists {
			t.Errorf("予期しない教育段階: %s", level.Name)
			continue
		}

		if level.GraduationYear != expected {
			t.Errorf("%s の卒業年 = %v, 期待値 %v", level.Name, level.GraduationYear, expected)
		}
	}
}
