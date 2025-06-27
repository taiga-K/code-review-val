package education

import (
	"time"
)

// EducationLevel は教育段階を表す構造体
type EducationLevel struct {
	Name           string // 学校名
	StartAge       int    // 入学年齢
	Duration       int    // 修業年数
	GraduationYear int    // 卒業年
}

// GraduationCalculator は卒業年計算の結果を保持する構造体
type GraduationCalculator struct {
	BirthYear int
	Levels    []EducationLevel
}

// NewGraduationCalculator は新しい卒業年計算機を作成します
func NewGraduationCalculator(birthDate time.Time) *GraduationCalculator {
	birthYear := birthDate.Year()

	calculator := &GraduationCalculator{
		BirthYear: birthYear,
		Levels: []EducationLevel{
			{Name: "小学校", StartAge: 6, Duration: 6},
			{Name: "中学校", StartAge: 12, Duration: 3},
			{Name: "高校", StartAge: 15, Duration: 3},
			{Name: "高専", StartAge: 15, Duration: 5},
			{Name: "短大", StartAge: 18, Duration: 2},
			{Name: "専門学校（2年制）", StartAge: 18, Duration: 2},
			{Name: "専門学校（3年制）", StartAge: 18, Duration: 3},
			{Name: "大学", StartAge: 18, Duration: 4},
		},
	}

	// 各教育段階の卒業年を計算
	for i := range calculator.Levels {
		calculator.Levels[i].GraduationYear = birthYear + calculator.Levels[i].StartAge + calculator.Levels[i].Duration
	}

	return calculator
}

// GetElementaryGraduation は小学校卒業年を返します
func (gc *GraduationCalculator) GetElementaryGraduation() int {
	return gc.BirthYear + 6 + 6 // 6歳入学 + 6年間
}

// GetMiddleSchoolGraduation は中学校卒業年を返します
func (gc *GraduationCalculator) GetMiddleSchoolGraduation() int {
	return gc.BirthYear + 12 + 3 // 12歳入学 + 3年間
}

// GetHighSchoolGraduation は高校卒業年を返します
func (gc *GraduationCalculator) GetHighSchoolGraduation() int {
	return gc.BirthYear + 15 + 3 // 15歳入学 + 3年間
}

// GetTechnicalCollegeGraduation は高専卒業年を返します
func (gc *GraduationCalculator) GetTechnicalCollegeGraduation() int {
	return gc.BirthYear + 15 + 5 // 15歳入学 + 5年間
}

// GetJuniorCollegeGraduation は短大卒業年を返します
func (gc *GraduationCalculator) GetJuniorCollegeGraduation() int {
	return gc.BirthYear + 18 + 2 // 18歳入学 + 2年間
}

// GetVocationalSchoolGraduation は専門学校卒業年を返します（2年制）
func (gc *GraduationCalculator) GetVocationalSchoolGraduation() int {
	return gc.BirthYear + 18 + 2 // 18歳入学 + 2年間
}

// GetUniversityGraduation は大学卒業年を返します
func (gc *GraduationCalculator) GetUniversityGraduation() int {
	return gc.BirthYear + 18 + 4 // 18歳入学 + 4年間
}

// GetAllGraduations はすべての教育段階の卒業年を返します
func (gc *GraduationCalculator) GetAllGraduations() []EducationLevel {
	return gc.Levels
}
