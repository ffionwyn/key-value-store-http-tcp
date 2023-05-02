package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	store "tcp/store"
)

//Routes the request depending on the method.
func parseRequest(input string) error {
	valueLower := strings.ToLower(input)
	parts := strings.Split(valueLower, "|")
	action := parts[0]

	switch action {
	case "add":
		err := validateParts(parts)
		if err != nil {
			return err
		}
		name := parts[1]
		location := parts[2]
		age := parts[3]
		if err != nil {
			return err
		}
		if !store.CheckPerson(name) {
			store.AddToStorage(name, location, age)
		} else {
			return fmt.Errorf("person already in the storage")
		}
	case "get":
		name := parts[1]
		loc, _, _ := store.GetPersonLocation(name)
		printTerminalMessages(fmt.Sprintf("Location found: %s", loc))
	case "delete":
		name := parts[1]
		err := store.DeletePerson(name)
		if err != nil {
			return err
		}
		log.Println("person deleted")
	case "update":
		err := validateParts(parts)
		if err != nil {
			return err
		}
		name := parts[1]
		location := parts[2]
		age := parts[3]
		if err != nil {
			return err
		}
		if store.CheckPerson(name) {
			store.AddToStorage(name, location, age)
		} else {
			return fmt.Errorf("person does not exist")
		}
	default:
		fmt.Println("Action not supported")
	}
	return nil
}

//Function for validating user input ensuring all 3 elements are filled in.
func validateParts(parts []string) error {
	if len(parts) != 4 {
		return fmt.Errorf("expected full name, location and age got %d", len(parts))
	}
	return nil
}

//Function for printing messages in the terminal.
func printTerminalMessages(message string) {
	fmt.Printf("%v: %s\n", time.Now(), message)
}

//  go run tcp-server/main.go 8000
func StartTCPServer(port string) {
	l, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}
		parseRequest(string(netData))

		fmt.Print("-> ", string(netData))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))
	}
}