package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"code-review-val/internal/calculator"
	"code-review-val/pkg/education"
)

func main() {
	fmt.Println("=== 年齢・卒業年計算プログラム ===")
	fmt.Println("生年月日を入力してください")

	// 生年月日の入力
	birthDate, err := getBirthDate()
	if err != nil {
		log.Fatalf("生年月日の入力でエラーが発生しました: %v", err)
	}

	// 現在の年齢を計算
	age := calculator.CalculateAge(birthDate)

	// 卒業年を計算
	gradCalc := education.NewGraduationCalculator(birthDate)

	// 結果の表示
	displayResults(birthDate, age, gradCalc)
}

// getBirthDate はユーザーから生年月日を入力してもらいます
func getBirthDate() (time.Time, error) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("生年月日を入力してください (YYYY-MM-DD形式): ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			fmt.Println("生年月日を入力してください。")
			continue
		}

		// 日付の解析
		birthDate, err := time.Parse("2006-01-02", input)
		if err != nil {
			fmt.Printf("正しい形式で入力してください (YYYY-MM-DD): %v\n", err)
			continue
		}

		// 未来の日付チェック
		if birthDate.After(time.Now()) {
			fmt.Println("未来の日付は入力できません。")
			continue
		}

		return birthDate, nil
	}
}

// displayResults は計算結果を表示します
func displayResults(birthDate time.Time, age int, gradCalc *education.GraduationCalculator) {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Printf("生年月日: %s\n", birthDate.Format("2006年01月02日"))
	fmt.Printf("現在の年齢: %d歳\n", age)
	fmt.Println(strings.Repeat("=", 50))

	fmt.Println("\n【各教育段階の卒業年】")
	fmt.Println(strings.Repeat("-", 40))

	levels := gradCalc.GetAllGraduations()
	for _, level := range levels {
		fmt.Printf("%-15s: %d年卒業\n", level.Name, level.GraduationYear)
	}

	fmt.Println(strings.Repeat("-", 40))
	fmt.Println("\n※ 一般的な進学パターンに基づいた計算です")
	fmt.Println("※ 飛び級、留年、浪人等は考慮されていません")
	fmt.Println("※ 4月入学を前提としています")
}
