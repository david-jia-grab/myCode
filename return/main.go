package main

import (
	"errors"
	"fmt"
)

func main() {
	i, e := ReturnID(10)
	fmt.Println(i)
	fmt.Println(e)

	i2, e2 := ReturnID2()
	fmt.Println(i2)
	fmt.Println(e2)

	i3, e3 := ReturnID3()
	fmt.Println(i3)
	fmt.Println(e3)

	i4 := AddNumbers(1, 2)
	fmt.Println(i4)
}

// AddNumbers demos result is not needed.
func AddNumbers(a int, b int) (result int) {
	return a + b
}

// ReturnID3 shows defer runs inline
func ReturnID3() (id int, err error) {
	defer func(id int) {
		if id == 10 {
			err = fmt.Errorf("invalid id")
		}
	}(id)

	id = 10

	return
}

// ReturnID is a function
func ReturnID(i int) (id int, err error) {
	id = i
	err = errors.New("oh bad things happened")
	return
}

// ReturnID2 is another func.
func ReturnID2() (id int, err error) {
	defer func() {
		if id == 10 {
			err = fmt.Errorf("invalid id")
		}
	}()

	id = 10

	return
}
