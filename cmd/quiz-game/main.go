package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/mrinaald/my-gophercises/assets"
)


type Problem struct {
	question string
	answer string
}

func main() {
	var fileName string
	flag.StringVar(&fileName, "filepath", assets.QuizGameDefaultProblemFile, "The CSV file containing quiz problems.")

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

	var userAnswer string
	correctAnswers := 0
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)
		fmt.Scanf("%s\n", &userAnswer)

		if userAnswer == problem.answer {
			correctAnswers++
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
