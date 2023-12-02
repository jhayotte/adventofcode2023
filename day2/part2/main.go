package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bag struct {
	Red   int
	Green int
	Blue  int
}
type Game struct {
	Number int
	Bags   []Bag
}

func main() {
	readFile("input_full.txt")
}

func readFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	b := Bag{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	sumIds := 0
	for scanner.Scan() {
		g := GetCubeConfiguration(scanner.Text())

		battleValid, powerNumber := IsValid(b, g)
		if battleValid {
			sumIds += powerNumber
		}
	}
	fmt.Println(sumIds)
}

// Example: Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func GetCubeConfiguration(game string) Game {
	g := Game{}

	//Game Number
	gameString := strings.Split(game, ":")[0]
	g.Number, _ = strconv.Atoi(strings.Split(gameString, " ")[1])

	// Iteration
	iterations := strings.Split(strings.Split(game, ":")[1], ";")
	battles := []Bag{}
	for _, battle := range iterations {

		colors := strings.Split(battle, ",")
		b := Bag{}
		for _, color := range colors {
			unitColor := strings.Split(color, " ")

			if unitColor[2] == "red" {
				b.Red, _ = strconv.Atoi(unitColor[1])
			}
			if unitColor[2] == "blue" {
				b.Blue, _ = strconv.Atoi(unitColor[1])
			}
			if unitColor[2] == "green" {
				b.Green, _ = strconv.Atoi(unitColor[1])
			}
		}
		battles = append(battles, b)
	}
	g.Bags = battles
	return g
}

func IsValid(b Bag, g Game) (valid bool, sum int) {
	maxBlue := 0
	maxRed := 0
	maxGreen := 0
	for _, battle := range g.Bags {
		if maxBlue < battle.Blue {
			maxBlue = battle.Blue
		}
		if maxRed < battle.Red {
			maxRed = battle.Red
		}
		if maxGreen < battle.Green {
			maxGreen = battle.Green
		}
	}
	fmt.Println("maxNumber::", maxBlue, maxGreen, maxRed, "power: ", (maxBlue * maxGreen * maxRed))
	return true, (maxBlue * maxGreen * maxRed)
}
