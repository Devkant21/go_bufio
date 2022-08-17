package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	hi := "good to see you again"
	fmt.Println(hi)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("On a scale of 1-10 how happy are you?")

	/* storing the reader data in the variable */

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks Cheer up", input)

}
