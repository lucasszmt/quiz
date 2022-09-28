package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/kr/pretty"
	"os"
	"strings"
	"time"
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
	answerChan := make(chan string)
	timer := time.NewTimer(time.Second * 10)
problemLoop:
	for i, p := range problems {
		pretty.Printf("Question #%d: %s\n", i+1, p.q)
		go func() {
			var answer string
			if _, err := fmt.Scanf("%s\n", &answer); err != nil {
				exit(err.Error())
			}
			answerChan <- answer
		}()
		select {
		case <-timer.C:
			pretty.Printf("Time over!\n Correct Answers %d Of %d", correct, len(problems))
			break problemLoop
		case a := <-answerChan:
			if a == p.a {
				correct++
			}
		}
	}
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
