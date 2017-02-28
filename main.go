package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var PLAYER1_FILE = "/var/tmp/player1.txt"
var PLAYER2_FILE = "/var/tmp/player2.txt"
var SCORE1_FILE = "/var/tmp/score1.txt"
var SCORE2_FILE = "/var/tmp/score2.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Players struct {
	Player1 string
	Score1  string
	Player2 string
	Score2  string
}

var player string

func scoreboard(w http.ResponseWriter, r *http.Request) {
	dat1, err := ioutil.ReadFile(PLAYER1_FILE)
	check(err)
	dat2, err := ioutil.ReadFile(PLAYER2_FILE)
	check(err)
	dat3, err := ioutil.ReadFile(SCORE1_FILE)
	check(err)
	dat4, err := ioutil.ReadFile(SCORE2_FILE)
	check(err)
	if r.Method == "GET" {
		player := Players{
			Player1: string(dat1),
			Player2: string(dat2),
			Score1:  string(dat3),
			Score2:  string(dat4),
		}
		t, err := template.ParseFiles("scoreboard.gtpl")
		check(err)
		t.Execute(w, player)
	} else {

		player := Players{
			Player1: r.FormValue("player1"),
			Player2: r.FormValue("player2"),
			Score1:  r.FormValue("score1"),
			Score2:  r.FormValue("score2"),
		}

		r.ParseForm()
		t, err := template.ParseFiles("scoreboard.gtpl")
		t.Execute(w, player)
		check(err)
		//Write values to file
		f1, err := os.OpenFile(PLAYER1_FILE, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0655)
		check(err)
		f2, err := os.OpenFile(PLAYER2_FILE, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0655)
		check(err)
		f3, err := os.OpenFile(SCORE1_FILE, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0655)
		check(err)
		f4, err := os.OpenFile(SCORE2_FILE, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0655)
		check(err)
		f1.WriteString(string(player.Player1))
		f2.WriteString(string(player.Player2))
		f3.WriteString(string(player.Score1))
		f4.WriteString(string(player.Score2))
		fmt.Printf("%s: %s\t%s: %s\n", string(player.Player1), string(player.Score1), string(player.Player2), string(player.Score2))
		check(err)
		f1.Sync()
		f1.Close()

	}
}

func main() {

	http.HandleFunc("/scoreboard", scoreboard)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
