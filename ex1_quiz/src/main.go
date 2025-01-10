package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const DEFAULT_TIMEOUT = 30

type config struct {
	problemsFilePath string
	timeoutInSeconds int
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

func playGame(problems []problem, gameResult *gameResult, doneChannel chan<- bool, errorChannel chan<- error) {
	for i, problem := range problems {
		var userAnswer string
		_, err := fmt.Printf("Problem #%d: %s = ", i+1, problem.question)
		if err != nil {
			errorChannel <- err
		}
		_, err = fmt.Scanln(&userAnswer)
		if err != nil {
			errorChannel <- err
		}
		if userAnswer == problem.answer {
			gameResult.correctAnswers++
		} else {
			gameResult.incorrectAnswers++
		}
	}

	doneChannel <- true
}

func main() {
	var config config

	flag.StringVar(&config.problemsFilePath, "file_path", "data/problems.csv", "Path to the file with questions (optional)")
	flag.IntVar(&config.timeoutInSeconds, "timeout", DEFAULT_TIMEOUT, fmt.Sprintf("Timeout. Defaults to %d seconds", DEFAULT_TIMEOUT))
	flag.Parse()

	absPath, err := filepath.Abs(config.problemsFilePath)
	checkErr(err)

	data, err := os.ReadFile(absPath)
	checkErr(err)

	problems := parseData(data)

	doneChannel := make(chan bool)
	errorChannel := make(chan error)
	gameResult := gameResult{
		correctAnswers:    0,
		incorrectAnswers:  0,
		numberOfQuestions: uint(len(problems)),
	}

	go playGame(problems, &gameResult, doneChannel, errorChannel)
	timer := time.NewTimer(time.Duration(config.timeoutInSeconds) * time.Second)

	select {
	case isDone := <-doneChannel:
		if isDone {
			fmt.Printf("You scored %d out of %d", gameResult.correctAnswers, gameResult.numberOfQuestions)
		}
	case err := <-errorChannel:
		panic(err)
	case <-timer.C:
		fmt.Println("\nOut of time!")
		fmt.Printf("You scored %d out of %d", gameResult.correctAnswers, gameResult.numberOfQuestions)
	}

	close(doneChannel)
	close(errorChannel)
}
