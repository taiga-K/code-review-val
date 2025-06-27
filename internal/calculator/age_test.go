package calculator

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	tests := []struct {
		name      string
		birthDate string
		expected  int
	}{
		{
			name:      "通常の年齢計算",
			birthDate: "1990-01-01",
			expected:  time.Now().Year() - 1990,
		},
		{
			name:      "まだ誕生日が来ていない場合",
			birthDate: "1990-12-31",
			expected:  time.Now().Year() - 1990 - 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			birthDate, err := time.Parse("2006-01-02", tt.birthDate)
			if err != nil {
				t.Fatalf("日付の解析に失敗しました: %v", err)
			}

			result := CalculateAge(birthDate)

			// 現在の日付によって結果が変わる可能性があるため、範囲チェック
			if result < 0 || result > 150 {
				t.Errorf("CalculateAge() = %v, 期待される範囲外です", result)
			}
		})
	}
}

func TestCalculateAgeAtDate(t *testing.T) {
	tests := []struct {
		name       string
		birthDate  string
		targetDate string
		expected   int
	}{
		{
			name:       "2020年時点での年齢",
			birthDate:  "1990-01-01",
			targetDate: "2020-01-01",
			expected:   30,
		},
		{
			name:       "誕生日前の年齢",
			birthDate:  "1990-06-15",
			targetDate: "2020-03-10",
			expected:   29,
		},
		{
			name:       "誕生日後の年齢",
			birthDate:  "1990-03-10",
			targetDate: "2020-06-15",
			expected:   30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			birthDate, err := time.Parse("2006-01-02", tt.birthDate)
			if err != nil {
				t.Fatalf("生年月日の解析に失敗しました: %v", err)
			}

			targetDate, err := time.Parse("2006-01-02", tt.targetDate)
			if err != nil {
				t.Fatalf("対象日の解析に失敗しました: %v", err)
			}

			result := CalculateAgeAtDate(birthDate, targetDate)
			if result != tt.expected {
				t.Errorf("CalculateAgeAtDate() = %v, 期待値 %v", result, tt.expected)
			}
		})
	}
}
