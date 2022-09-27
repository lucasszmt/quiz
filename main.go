package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/kr/pretty"
	"os"
	"strings"
)

var (
	filePath = flag.String("csv path", "problems.csv", "a csv filePath containing quiz questions and solutions")
)

func main() {
	flag.Parse()
	f, err := os.Open(*filePath)
	if err != nil {
		exit(fmt.Sprintf("Error opening the file %s", err))
	}
	r := csv.NewReader(f)
	res, _ := r.ReadAll()
	problems := ParseProblem(res)
	var correct int
	for i, p := range problems {
		pretty.Printf("Question #%d: %s\n", i+1, p.q)
		var answer string
		if _, err := fmt.Scanf("%s\n", &answer); err != nil {
			exit(err.Error())
		}
		if answer == p.a {
			correct++
		}
	}
	pretty.Printf("Correct answers: %d", correct)
}

type Problem struct {
	q string
	a string
}

func ParseProblem(records [][]string) []Problem {
	ps := make([]Problem, len(records))
	for i, r := range records {
		ps[i] = Problem{r[0], strings.TrimSpace(r[1])}
	}
	return ps
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
