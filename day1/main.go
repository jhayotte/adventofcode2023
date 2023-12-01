package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Part 2
func main() {

	//file, err := os.Open("input_example_part2.txt")
	file, err := os.Open("input_full.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	sum := 0

	for scanner.Scan() {

		s := scanner.Text()
		fmt.Println(s)

		orderWords := ReplaceWordByNum(s)
		firstDigit := FirstChar(orderWords)
		reversedStr := Reverse(s)
		reversedWords := ReplaceReversedWordByNum(reversedStr)
		lastDigit := FirstChar(reversedWords)

		digits := fmt.Sprintf("%v%v", firstDigit, lastDigit)
		fmt.Println(digits)
		d, _ := strconv.Atoi(digits)
		sum += d
	}
	fmt.Println(sum)
}

func ReplaceWordByNum(s string) (result string) {
	for index, _ := range s {
		if strings.Contains(s[:index+1], "zero") {
			s = strings.Replace(s, "zero", "0", -1)
			return ReplaceWordByNum(s)
		}
		if strings.Contains(s[:index+1], "one") {
			s = strings.Replace(s, "one", "1", -1)
			return ReplaceWordByNum(s)
		}
		if strings.Contains(s[:index+1], "two") {
			s = strings.Replace(s, "two", "2", -1)
			return ReplaceWordByNum(s)
		}
		if strings.Contains(s[:index+1], "three") {
			s = strings.Replace(s, "three", "3", -1)
			return ReplaceWordByNum(s)
		}
		if strings.Contains(s[:index+1], "four") {
			s = strings.Replace(s, "four", "4", -1)
			return ReplaceWordByNum(s)
		}
		if strings.Contains(s[:index+1], "five") {
			s = strings.Replace(s, "five", "5", -1)
			return ReplaceWordByNum(s)
		}
		if strings.Contains(s[:index+1], "six") {
			s = strings.Replace(s, "six", "6", -1)
			return ReplaceWordByNum(s)
		}
		if strings.Contains(s[:index+1], "seven") {
			s = strings.Replace(s, "seven", "7", -1)
			return ReplaceWordByNum(s)
		}
		if strings.Contains(s[:index+1], "eight") {
			s = strings.Replace(s, "eight", "8", -1)
			return ReplaceWordByNum(s)
		}
		if strings.Contains(s[:index+1], "nine") {
			s = strings.Replace(s, "nine", "9", -1)
			return ReplaceWordByNum(s)
		}
	}
	return s
}

func ReplaceReversedWordByNum(s string) (result string) {
	for index, _ := range s {
		if strings.Contains(s[:index+1], "orez") {
			s = strings.Replace(s, "orez", "0", -1)
			return ReplaceReversedWordByNum(s)
		}
		if strings.Contains(s[:index+1], "eno") {
			s = strings.Replace(s, "eno", "1", -1)
			return ReplaceReversedWordByNum(s)
		}
		if strings.Contains(s[:index+1], "owt") {
			s = strings.Replace(s, "owt", "2", -1)
			return ReplaceReversedWordByNum(s)
		}
		if strings.Contains(s[:index+1], "eerht") {
			s = strings.Replace(s, "eerht", "3", -1)
			return ReplaceReversedWordByNum(s)
		}
		if strings.Contains(s[:index+1], "ruof") {
			s = strings.Replace(s, "ruof", "4", -1)
			return ReplaceReversedWordByNum(s)
		}
		if strings.Contains(s[:index+1], "evif") {
			s = strings.Replace(s, "evif", "5", -1)
			return ReplaceReversedWordByNum(s)
		}
		if strings.Contains(s[:index+1], "xis") {
			s = strings.Replace(s, "xis", "6", -1)
			return ReplaceReversedWordByNum(s)
		}
		if strings.Contains(s[:index+1], "neves") {
			s = strings.Replace(s, "neves", "7", -1)
			return ReplaceReversedWordByNum(s)
		}
		if strings.Contains(s[:index+1], "thgie") {
			s = strings.Replace(s, "thgie", "8", -1)
			return ReplaceReversedWordByNum(s)
		}
		if strings.Contains(s[:index+1], "enin") {
			s = strings.Replace(s, "enin", "9", -1)
			return ReplaceReversedWordByNum(s)
		}
	}
	return s
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func FirstChar(s string) (result int) {
	for _, v := range s {
		d, err := strconv.Atoi(string(v))
		if err == nil {
			return d
		}
	}
	return -1
}
