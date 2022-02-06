package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// connect to a room
func join() net.Conn {
	var (
		ip   string
		port string
	)
	fmt.Print("Enter the ip: ")
	fmt.Scanln(&ip)
	fmt.Print("Enter the port: ")
	fmt.Scanln(&port)
	serv, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		os.Exit(1)
	}
	scan := bufio.NewScanner(serv)
	// get connected message from the server
	scan.Scan()
	fmt.Println(scan.Text())
	// get wait message from the server
	scan.Scan()
	fmt.Println(scan.Text())
	return serv
}

func play(room net.Conn) {
	fmt.Println("Game started")
	var score int = 0
	scan := bufio.NewScanner(room)
	for i := 0; i < 10; i++ {
		// play the game
		number := get_number()
		fmt.Fprintf(room, "%d\n", number)
		// get the result
		scan.Scan()
		result := scan.Text()
		// print the result
		fmt.Println("The winner sent the number", result, "and you sent", number)
		if result != fmt.Sprint(number) {
			fmt.Println("You lose")
		} else {
			fmt.Println("You win")
			score += number
		}
	}
	fmt.Println("Your score is", score)
}

func get_number() int {
	for {
		fmt.Print("Enter a number between 1-6 : ")
		var number int
		fmt.Scanf("%d", &number)
		if number > 0 && number < 7 {
			return number
		}
	}
}
