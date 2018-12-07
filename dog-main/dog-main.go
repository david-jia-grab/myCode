package main

import (
	"fmt"

	"github.com/david-jia-grab/myCode/dog"
)

func main() {
	d := dog.Dog{"Rex", dog.Poodle}
	fmt.Println(d)
}
