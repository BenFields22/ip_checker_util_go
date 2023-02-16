package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func execute() {

	// here we perform the pwd command.
	// we can store the output of this in our out variable
	// and catch any errors in err
	out, err := exec.Command("curl", "https://ipecho.net/plain").Output()

	// if there is an error with our execution
	// handle it here
	if err != nil {
		fmt.Printf("%s", err)
	}
	// as the out variable defined above is of type []byte we need to convert
	// this to a string or else we will see garbage printed out in our console
	// this is how we convert it to a string
	fmt.Print("\nPublic IP: ")
	output := string(out[:])
	fmt.Println(output)

	// let's try the pwd command herer
	out, err = exec.Command("hostname", "-i").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Print("Private IP: ")
	output = string(out[:])
	words := strings.Fields(output)
	fmt.Println(words[2] + "\n")
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		execute()
	}
}
