package printb

import (
	"fmt"

	printc "github.com/AndreyLevchenko/modc"
)

func PrintInfo() {
	fmt.Println("-----------------------------")
	fmt.Printf("ModB version: %d \n", 1)
	printc.PrintInfo()
}
