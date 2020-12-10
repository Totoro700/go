package main

import (
    "fmt -s"
)

func main(){
    fmt.Println("Hello!")
    fmt.Println("What is 1 * 1?")
    var input;
    fmt.Scanline(&input)
    if (input == "1"){
		fmt.Println("Correct!")
	}else {
		fmt.Println("Incorrect!")
	}
}
