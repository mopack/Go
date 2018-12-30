// LetsChat
package main

import (
	"fmt"
	"os"
)

type UserInfo struct {
	ID string
	ch chan string
}

func main() {
	// Data Vars Global
	const UserLimit = 1000
	var Users [UserLimit]UserInfo
	var UserNumber = 0

	// Running Vars
	var getS, getR string // inputString, inputRemain
	var u UserInfo        // for index
	var i int             // for index
	var MyID string       //

	// -- Starter
	fmt.Println("Welcome to LetsChat!")

InputID:
	// Limit Check
	if UserNumber >= UserLimit {
		fmt.Println("Sorry, LetsChat is full currently. Please reload in the while.")
		os.Exit(0)
	}

	fmt.Print("Please input your id (length 1~10):")
	fmt.Scanf("%s", &getS)
	fmt.Scanf("%s", &getR)

	// Length Check
	if len(getS) > 10 || len(getS) == 0 {
		fmt.Println("Your id is out of range, please input again.")
		goto InputID
	}

	// Same Name Check
	for _, u = range Users {
		if u.ID == getS {
			fmt.Println("Your id have been used, please input again.")
			goto InputID
		}
	}

	MyID = getS

	// Reset Check
	fmt.Println("Your id will be:", MyID, ". Is that ok? (Y/N):")
	fmt.Scanf("%s", &getS)
	fmt.Scanf("%s", &getR)

	if getS != "Y" && getS != "y" { //Y, y
		fmt.Println("Please reset your id.")
		goto InputID
	}
	
	fmt.Println("You're now login into LetsChat")
	fmt.Println("If you wanna logout just text \"EXIT\"")

	// Board Cast Login to Others
	for i = 0; i < UserNumber; i++ {
		Users[i].ch <- ("Welcome: " + )
	}
	
	// Setup Account
}
