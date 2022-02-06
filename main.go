package main

import "fmt"

func main() {
	// do you wanna host
	var (
		is_host string
		c       = make(chan bool)
	)
	print("Do you wanna host (y/n) : ")
	fmt.Scanln(&is_host)
	if is_host == "y" {
		setup := setup_room()
		fmt.Println(setup)
		go host(setup, c)
		wait_join(c)
		close_room(c)
	} else {
		// join a room
		room := join()
		// play
		play(room)
	}
}
