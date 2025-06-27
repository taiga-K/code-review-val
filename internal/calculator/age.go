package calculator

import (
	"time"
)

// CalculateAge は生年月日から現在の年齢を計算します
func CalculateAge(birthDate time.Time) int {
	now := time.Now()
	age := now.Year() - birthDate.Year()

	// 誕生日がまだ来ていない場合は1を引く
	if now.Month() < birthDate.Month() ||
		(now.Month() == birthDate.Month() && now.Day() < birthDate.Day()) {
		age--
	}

	return age
}

// CalculateAgeAtDate は指定された日付での年齢を計算します
func CalculateAgeAtDate(birthDate, targetDate time.Time) int {
	age := targetDate.Year() - birthDate.Year()

	// 誕生日がまだ来ていない場合は1を引く
	if targetDate.Month() < birthDate.Month() ||
		(targetDate.Month() == birthDate.Month() && targetDate.Day() < birthDate.Day()) {
		age--
	}

	return age
}
