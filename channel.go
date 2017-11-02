package main

import (
	"time"
	"math/rand"
	"fmt"
	"strconv"
)

var s1 = rand.NewSource(time.Now().UnixNano())
var r1 = rand.New(s1)

func generator (gout chan string, id int) {

	for {
		time.Sleep(time.Second * time.Duration(r1.Intn(8)))
		gout <- "Message in from - " + strconv.Itoa(id)
	}
}

func main() {
	d:= make(chan string, 10)
	go generator(d, 1)
	go generator(d, 2)

	for {
		select{
		case msg:= <- d:
			fmt.Println(msg)
		}
	}

}
