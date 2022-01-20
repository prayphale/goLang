package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFile := flag.String("CSV", "problems.csv", "a csv file in the format of 'question.answer'")
	timeLimit := flag.Int("Limit", 30, "the time ")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		fmt.Printf("Failed to open the CSV file: %s", *csvFile)
		os.Exit(1)
	}
	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the CSV file")
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correctanswers := 0

quizloop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s =\n", i+1, p.question)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break quizloop
		case answer := <-answerCh:
			if answer == p.answer {
				correctanswers++
			}
		}
	}

	fmt.Printf("\nScored %d out of %d.\n", correctanswers, len(problems))
}

func parseLines(line [][]string) []problem {
	ret := make([]problem, len(line))

	for i, lin := range line {
		ret[i] = problem{
			question: lin[0],
			answer:   strings.TrimSpace(lin[1]),
		}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
