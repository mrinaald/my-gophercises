package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mrinaald/my-gophercises/assets"
)


type Problem struct {
	question string
	answer string
}

func main() {
	var fileName string
	flag.StringVar(&fileName, "filepath", assets.QuizGameDefaultProblemFile, "The CSV file containing quiz problems.")

	timeLimit := flag.Int("timelimit", 30, "The time limit for the quiz in seconds")
	flag.Parse()

	// just to check whether file exists: https://stackoverflow.com/a/12518877

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
        os.Exit(1)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)

	lines, err := csvReader.ReadAll()
    if err != nil {
        os.Exit(1)
    }

	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correctAnswers := 0
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)
		answerCh := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case userAnswer := <-answerCh :
			if userAnswer == problem.answer {
				correctAnswers++
			}
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d.\n", correctAnswers, len(lines))
			return
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correctAnswers, len(lines))
}

func parseLines(lines [][]string) []Problem {
	problems := make([]Problem, len(lines))
	for i, line := range lines {
		problems[i] = Problem{
			question: line[0],
			answer: strings.TrimSpace(line[1]),
		}
	}

	return problems
}
