package scoreboard

import (
	"os"
	"bufio"
	"fmt"
	_"strings"
)
func check(e error) {
    if e != nil {
        panic(e)
    }
}


// PutScoreInScoreBoard Put scores in score board
func PutScoreInScoreBoard(username string, score float64) {
	f, err := os.OpenFile("scores.dat",
	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer f.Close()
	f.Sync()
	w := bufio.NewWriter(f)
	scoreInString := fmt.Sprintf("%f", score)
	writeit := username + " : " + scoreInString
	w.WriteString(writeit + "\n")
	//check(err)
	w.Flush()
}

// LeaderBoard Prints the top 3 players
func LeaderBoard() {
	f, err := os.OpenFile("scores.dat", os.O_RDONLY, 0644)
	check(err)
	defer func(){
		f.Close()
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		fmt.Println(line)
	}
}