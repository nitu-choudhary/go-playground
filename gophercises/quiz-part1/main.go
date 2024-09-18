package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func calculateScore(input int, answer int, score int) int {
	if input == answer {
		score++
	}
	return score
}

func main() {

	// score variable
	score := 0
	// open the csv file
	file, err := os.Open("gophercises/quiz-part1/problems.csv")
	if err != nil {
		panic(err)
	}

	// read the csv data
	reader := csv.NewReader(file)

	i := 0
	for {
		var input int
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("#Problem #%d: %v = ", i+1, record[0])
		_, err = fmt.Scanf("%d", &input)
		if err != nil {
			panic(err)
		}
		num, _ := strconv.Atoi(record[1])
		score = calculateScore(input, num, score)
		i++
	}

	fmt.Printf("Total : %d / %d\n", score, i)

	defer file.Close()
}
