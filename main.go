package main

import (
	"fmt"
	// "bufio"
	// "os"
)

// func userInputValid(value, planet) {
// 	if(yes) {
// 		ask for planet name
// 	}else if(no){
// 		fmt.Printf("I'm sorry I didn't get that.\n")
// 	}

// }

func main() {

	fmt.Printf("Welcome to the Solar System!\n")
	fmt.Printf("There are nine planets to explore.\n")
	fmt.Printf("What is your name?\n")
	var name string
    fmt.Scanln(&name)
	fmt.Printf("Nice to meet you, %q", name,"! My name is Eliza, I am an old friend of Alexa.\n")
	fmt.Printf("Lets go on an adventure!\n")
	fmt.Printf("Shall I randomly choose a planet for you to visit? (Y or N)\n")
	var answer string
	fmt.Scanln(&answer)
	if(answer == "Y"){
		fmt.Printf("Travelling to [insert random planet]...\n")
	}else if(answer == "N"){
		fmt.Printf("Please enter the name of the planet you would like to visit.\n")
		//Do planet math		
	}else{
		fmt.Printf("I'm sorry I didn't get that.\n")
		//recursion
	}
	fmt.Printf("Welcome to the Solar System!\n")
	fmt.Printf("Travelling to [insert planet]...\n")
	fmt.Printf("Welcome to Mars\n")
	fmt.Printf("Welcome to Saturn\n")
	fmt.Printf("Welcome to Neptune\n")
	fmt.Printf("Welcome to Uranus\n")
	fmt.Printf("Welcome to Pluto\n")
	fmt.Printf("Welcome to Venus\n")
	fmt.Printf("Welcome to Jupiter\n")
	fmt.Printf("Welcome to Mercury\n")
	fmt.Printf("Welcome to Earth\n")
	fmt.Printf("Welcome to Null PlAnEt\n")

}