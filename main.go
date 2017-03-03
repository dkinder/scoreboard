package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Players struct {
	Player1 string
	Score1  string
	File1   string
	Player2 string
	Score2  string
	File2   string
	List    []string
}

var player string

func scoreboard(w http.ResponseWriter, r *http.Request) {

	list, errs := os.Open("./players.txt")
	check(errs)

	var lines []string
	scanner := bufio.NewScanner(list)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var files [4]string
	files[0] = "/var/tmp/player1.txt"
	files[1] = "/var/tmp/player2.txt"
	files[2] = "/var/tmp/score1.txt"
	files[3] = "/var/tmp/score2.txt"

	dat1, err := ioutil.ReadFile(files[0])
	check(err)
	dat2, err := ioutil.ReadFile(files[1])
	check(err)
	dat3, err := ioutil.ReadFile(files[2])
	check(err)
	dat4, err := ioutil.ReadFile(files[3])
	check(err)
	if r.Method == "GET" {
		player := Players{
			Player1: string(dat1),
			Player2: string(dat2),
			Score1:  string(dat3),
			Score2:  string(dat4),
			List:    lines,
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
			List:    lines,
		}

		r.ParseForm()
		t, err := template.ParseFiles("scoreboard.gtpl")
		check(err)
		t.Execute(w, player)
		check(err)
		//Write values to file
		f1, err := os.OpenFile(files[0], os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0655)
		check(err)
		f2, err := os.OpenFile(files[1], os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0655)
		check(err)
		f3, err := os.OpenFile(files[2], os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0655)
		check(err)
		f4, err := os.OpenFile(files[3], os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0655)
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

	port := ":9090"
	http.HandleFunc("/scoreboard", scoreboard)
	fmt.Printf("Listening at http://localhost%s/scoreboard", port)
	err := http.ListenAndServe(port, nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
