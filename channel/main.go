package main

import (
	"fmt"
	"time"
)

var irc = make(chan string)
var sms = make(chan string)

func pinger(canal chan string) {
	for {
		canal <- "ping"
	}
}

func ponger(canal chan string) {
	for {
		canal <- "pong"
	}
}

func eai(canal chan string) {
	for {
		canal <- "e ai"
	}
}
func blz(canal chan string) {
	for {
		canal <- "blz"
	}
}

func impressora() {
	var msg string
	for {
		select {
		case msg = <-irc:
			fmt.Println(msg)
			time.Sleep(time.Second * 1)
		case msg = <-sms:
			fmt.Println("zzz ...z zzz", msg)
		}
		time.Sleep(time.Second * 1)
	}
}

func main() {
	go pinger(irc)
	go ponger(irc)
	go eai(sms)
	go eai(sms)
	go impressora()

	var entrada string
	fmt.Scan(&entrada)
}
