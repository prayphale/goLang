package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFile := flag.String("CSV", "problems.csv", "a csv file in the format of 'question.answer'")
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

	correctAnswers := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s =\n", i+1, p.question)
		var answers string
		fmt.Scanf("%s\n", &answers)
		if answers == p.answer {
			correctAnswers++
		}
	}

	fmt.Printf("scored %d out of %d.\n", correctAnswers, len(problems))
}

func parseLines(line [][]string) []problem {
	ret := make([]problem, len(line))

	for i, lin := range line {
		ret[i] = problem{
			question: lin[0],
			answer:   lin[1],
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
