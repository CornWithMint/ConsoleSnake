package main

import (
	"fmt"
	"os"
	"strconv"
)

func WriteScore(score int) {

	StrScore := strconv.Itoa(score)
	file, err := os.OpenFile("Score.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {

	}
	defer file.Close()

	file.WriteString(StrScore)
}

func GetScore(score int) int {
	_, err := os.Stat("Score.txt")
	if err == nil {
		bytes, err := os.ReadFile("Score.txt")
		if err != nil {
			fmt.Println("ad")
		}
		return int(bytes[0] - 48)
	} else {
		WriteScore(score)
		return score
	}

}
