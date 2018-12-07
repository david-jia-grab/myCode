package dog

import "fmt"

type Dog struct {
	Name     string
	TheBreed Breed
}

// we now have a value receiver
func (d Dog) String() string {
	return fmt.Sprintf("My name is %s, I'm a %s! Woof!", d.Name, d.TheBreed)
}
