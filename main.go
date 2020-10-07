package main

import (
	// "os"
	// "bufio"
	"fmt"
	"fastmath/gameplay"
	"fastmath/scoreboard"
)
	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

	// f, err := os.OpenFile("scores.dat",
	// os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// check(err)
	// defer f.Close()
	// f.Sync()
	// w := bufio.NewWriter(f)
	// n, err := w.WriteString("user1 2.6\n")
	// check(err)
	// fmt.Println("Score Written ", n)
	// w.Flush()
	scoreboard.LeaderBoard()
	var username string
	fmt.Println("Enter Your Username")
	fmt.Scanln(&username)
	fmt.Println("Hello, ", username, " Let's Play!")
	fmt.Println()
	gameplay.Play(username, "L1")
}