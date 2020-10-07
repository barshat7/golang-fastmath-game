package gameplay

import (
	"fmt"
	"flag"
	"os"
	"bufio"
	"regexp"
	"strconv"
	"strings"
	"fastmath/scorer"
	"math/rand"
	"time"
	"fastmath/scoreboard"
)


var questions [] question
const totalQuestions = 10

// Gameplay holds the current session of the user
type Gameplay struct {

	user string
	level string
	deliveredQuestionCount int
	deliveredQuestions [] int
	currentScore float64
}

func (g Gameplay) displayScore() {
	fmt.Println("Your Score Is ", g.currentScore)
}

type question struct {
	level string
	id int
	questionText string
	answer string
}

func (q question) getQuestion() string {
	return q.questionText
}

func (q question) verifyCorrect(userAnswer string) bool {
	return q.answer == userAnswer
}


func check(err error) {
	if err != nil {
		panic(err)
	}
}

func parseQuestion(questionLine string) question {
	delimeter := regexp.MustCompile(`:`)
	s := delimeter.Split(questionLine, -1)
	level := string(s[0])
	id, _ := strconv.Atoi(string(s[1])) 
	questionText := string(s[2])
	answer := strings.TrimSpace(string(s[3]))
	question := question {level: level, id: id, questionText: questionText, answer: answer}
	return question
}

func loadQuestions() [] question {
	var allQuestions [] question
	datafile := "./data/questions.dat"
	fptr := flag.String("fpath", datafile, "file to be read from")
	flag.Parse()
	f, err := os.Open(*fptr)
	check(err)
	defer func(){
		f.Close()
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		allQuestions = append(allQuestions, parseQuestion(s.Text()))
	}
	check(s.Err())
	return allQuestions
}

func loadAllQuestionsInMemory(level string) {
	questions = loadQuestions()
}

func generateRandomIDInRange(totalQuestions int) func() int {
	theRange := totalQuestions
	rand.Seed(time.Now().UnixNano())
	return func () int {
		return rand.Intn(theRange)
	}
}

// Play The Game
func Play(user string, level string) {
	gamePlay := Gameplay{user: user, level: level, deliveredQuestionCount: 0, currentScore: 0.0}
	loadAllQuestionsInMemory(level)
	randomness := generateRandomIDInRange(len(questions))
	for {
		quid := randomness()
		if (contains(gamePlay.deliveredQuestions, quid)) {
			continue
		}
		q := questions[quid]
		start := getCurrentTimeInMillis()
		fmt.Println(q.getQuestion())
		var answer string
		fmt.Scanln(&answer)
		end := getCurrentTimeInMillis()
		gamePlay.currentScore = scorer.Score(q.verifyCorrect(answer), gamePlay.currentScore, end - start)
		gamePlay.deliveredQuestions = append(gamePlay.deliveredQuestions, quid)
		gamePlay.deliveredQuestionCount ++
		if (gamePlay.deliveredQuestionCount == totalQuestions) {
			break
		}
	}
	gamePlay.displayScore()
	scoreboard.PutScoreInScoreBoard(gamePlay.user, gamePlay.currentScore)
}

func getCurrentTimeInMillis() int64 {
    return time.Now().UnixNano() / int64(time.Millisecond)
}

func contains(quids [] int, quid int) bool {
	for _, q := range quids {
		if (q == quid) {
			return true
		}
	}
	return false
}