package main

import "fmt"

// START DEF OMIT
type Greeter interface {
	Greet() string
}

// Please don't format your code like this, it's just to make it fit!
type FrenchGreeter struct { nom string }
func (n *FrenchGreeter) SetNom(nom string) { n.nom = nom }
func (n *FrenchGreeter) Greet() string {
	return fmt.Sprintf("Bonjour, %s!", n.nom)
}

type EnglishGreeter struct { name string }
func (n *EnglishGreeter) SetName(name string) { n.name = name }
func (n *EnglishGreeter) Greet() string {
	return fmt.Sprintf("Hello, %s!", n.name)
}
 
// Empty struct can be used if you want to implement an interface with no storage
type GenericGreeter struct{}
func (g *GenericGreeter) Greet() string { return "Hi!" }
// END DEF OMIT

// START MAIN OMIT
func PrintWelcome(greeter Greeter, name string) {
	switch g := greeter.(type) {
	case *FrenchGreeter:
		g.SetNom(name)
	case *EnglishGreeter:
		g.SetName(name)
	default:
		fmt.Println("Warning: Don't know how to set name...")
	}
	fmt.Println(greeter.Greet())
}

func main() {
	PrintWelcome(&FrenchGreeter{}, "Alice")
	PrintWelcome(&EnglishGreeter{}, "Bob")
	PrintWelcome(&GenericGreeter{}, "Mr Nobody")
}

// END MAIN OMIT
