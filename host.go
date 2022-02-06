package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

func host(setup room, c chan bool) {
	serv, err := net.Listen("tcp", ":"+setup.port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Server started on port", setup.port)
	fmt.Println("Waiting for players to join")
	// wait for clients
	clients := get_clients(setup, serv)
	// send to clients that the game will begin
	for _, client := range clients {
		fmt.Fprintf(client, "Game will begin soon\n")
	}
	c <- true // send signal to main to continue
	// server side game loop
	serv_play(clients)
	fmt.Println("Game finished")
	c <- true // send signal the game is finished
	// close the room
	<-c
	disconnect_clients(clients)
	serv.Close()
	fmt.Println("Room closed")
}

func serv_play(clients []net.Conn) {
	for i := 0; i < 10; i++ {
		var result []string
		// get the result from each clients in goroutines
		fmt.Println("Round", i+1)
		for _, client := range clients {
			go func(client net.Conn) {
				scan := bufio.NewReader(client)
				n, _ := scan.ReadString('\n')
				fmt.Print("Client", client.RemoteAddr(), ":", n)
				result = append(result, n)
			}(client)
		}
		// wait for all goroutines to finish
		for len(result) < len(clients) {
			continue
		}
		// process the result
		score := process_result(result)
		fmt.Println("Winner is", score)
		// send the result to each clients
		for _, client := range clients {
			fmt.Fprintf(client, "%d\n", score)
		}
	}
}

// process the result
// the single smallest number is the winner
func process_result(result []string) int {
	answer := []int{}
	for _, r := range result {
		// remove the newline character
		r = r[:len(r)-1]
		r, err := strconv.Atoi(r)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		answer = append(answer, r)
	}
	// if there are multiple numbers, remove them
	winner := []int{}
	buff := []int{}
	reccurence := map[int]int{}
	// count the number of times each number appears
	for _, a := range answer {
		if isInt(a, buff) { // if the number is already in the buffer
			reccurence[a]++
		} else { // if the number is not in the buffer
			buff = append(buff, a)
			reccurence[a] = 1
		}
	}
	fmt.Println("reccurence", reccurence)
	// add to winner the number that appears only once
	for k, v := range reccurence {
		if v == 1 {
			winner = append(winner, k)
		}
	}
	fmt.Println("Winner content :", winner)
	if len(winner) < 1 {
		fmt.Println("No winner")
		return -1
	}
	return min(winner)
}

func isInt(i int, a []int) bool {
	for _, v := range a {
		if v == i {
			return true
		}
	}
	return false
}

func min(a []int) int {
	min := a[0]
	if len(a) == 1 {
		return a[0]
	}
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

func get_clients(setup room, serv net.Listener) []net.Conn {
	var clients []net.Conn
	for i := 0; i < setup.n_players; i++ {
		buff_client, err := serv.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		clients = append(clients, buff_client)
		fmt.Fprintf(buff_client, "Welcome to \"%s\", we're waiting for other players %d/%d\n", setup.name, i+1, setup.n_players)
		fmt.Println(len(clients), "Player connected")
	}
	return clients
}

func disconnect_clients(clients []net.Conn) {
	for _, client := range clients {
		client.Close()
	}
}

// wait for players to join
func wait_join(c chan bool) {
	<-c
}

// wait for all player to play the round
// code

// close the room
func close_room(c chan bool) {
	<-c // wait for the game to finish
	fmt.Println("Closing room")
	// send signal to close the room
	c <- true
}
