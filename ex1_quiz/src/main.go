package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type config struct {
	problemsFilePath string
}

type problem struct {
	question string
	answer   string
}

type gameResult struct {
	correctAnswers    uint
	incorrectAnswers  uint
	numberOfQuestions uint
}

func checkErr(err error) {
	if err != nil {
		// fmt.Println(err)
		panic(err)
	}
}

func parseData(data []byte) []problem {
	reader := csv.NewReader(strings.NewReader(string(data)))

	readRows, err := reader.ReadAll()
	var parsedProblems = make([]problem, len(readRows))
	if err != nil {
		log.Fatal(err)
	}

	for i, row := range readRows {
		parsedProblems[i] = problem{
			question: row[0],
			answer:   row[1],
		}
	}

	return parsedProblems
}

func playGame(problems []problem) (gameResult, error) {
	gameResult := gameResult{
		correctAnswers:    0,
		incorrectAnswers:  0,
		numberOfQuestions: uint(len(problems)),
	}

	for i, problem := range problems {
		var userAnswer string
		_, err := fmt.Printf("Problem #%d: %s = ", i+1, problem.question)
		if err != nil {
			return gameResult, err
		}
		_, err = fmt.Scanln(&userAnswer)
		if err != nil {
			return gameResult, err
		}
		if userAnswer == problem.answer {
			gameResult.correctAnswers++
		} else {
			gameResult.incorrectAnswers++
		}
	}

	return gameResult, nil
}

func main() {
	var config config

	flag.StringVar(&config.problemsFilePath, "file_path", "data/problems.csv", "Specify path to the file with questions")
	flag.Parse()

	absPath, err := filepath.Abs(config.problemsFilePath)
	checkErr(err)

	data, err := os.ReadFile(absPath)
	checkErr(err)

	problems := parseData(data)

	gameResult, err := playGame(problems)
	checkErr(err)

	fmt.Printf("You scored %d out of %d", gameResult.correctAnswers, gameResult.numberOfQuestions)
}
