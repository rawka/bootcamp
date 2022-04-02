package main

import (
	"fmt"
	"os"
)

const (
	usage = "Usage: [username] [password]"
	errUsr = "Access denied for %q\n"
	errPwd = "Invalid password for %q\n"
	grant = "Access granted to %q\n"
	
	user, user2 = "user", "user2"
	pass, pass2 = "pass", "pass2"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]

	if u != user && u != user2 {
		fmt.Printf(errUsr, u) 
	} else if u == user && p == pass {
		fmt.Printf(grant, u)
	} else if u == user2 && p == pass2 {
		fmt.Printf(grant, u)
	} else {
		fmt.Printf(errPwd, u) 
	}
}
