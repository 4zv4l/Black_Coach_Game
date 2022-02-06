package main

import (
	"bufio"
	"os"
	"strconv"
)

type room struct {
	name      string
	n_players int
	port      string
	// can add more settings here
}

func setup_room() room {
	var (
		setup = room{}
		scan  = bufio.NewScanner(os.Stdin)
		n     = 0
		err   error
	)
	print("Room's name : ")
	scan.Scan()
	setup.name = scan.Text()
	for n < 3 {
		print("How many players (min 3) : ")
		scan.Scan()
		n, err = strconv.Atoi(scan.Text())
		if err != nil {
			println("Please enter a number...")
		}
	}
	setup.n_players = n
	print("Port : ")
	scan.Scan()
	setup.port = scan.Text()
	return setup
}
