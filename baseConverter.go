package main

import(
	"fmt"
	"os"
	"strconv"
)

func main(){
	var number int
	var err error

	fmt.Println("Base Converter!\n")

	if len(os.Args) < 2 {
		fmt.Print("Please enter a number to convert:")
		fmt.Scan(&number)
	} else {
		number, err = strconv.Atoi(os.Args[1])
	
		if err != nil {
			fmt.Println("You must enter an integer!")
			panic(err)
		}
	}
	
	fmt.Printf("Decimal value: %d\n", number)
	
	fmt.Println("--------------")

	fmt.Printf("Binary: %b\n", number)
	fmt.Printf("Octal: %O\n", number)
	fmt.Printf("Hex: 0x%X\n", number)
}
