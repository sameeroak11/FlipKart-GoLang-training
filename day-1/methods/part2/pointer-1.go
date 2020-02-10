// methods:
// methods on userdefined type of built-in data-type


package main
import "fmt"

type Number int

// method inc has a receiver of type pointer to Number
func (pn *Number) inc(msg string) {
	*pn++
	fmt.Printf("in inc():  %s,  *pn: %d\n", msg, *pn)
}

// method print has a receiver of type Number
func (n Number) print(msg string) {
	fmt.Printf("in print():  %s:  n: %d\n", msg, n)
}

func main() {
	i := Number(10) // i is of type Number and is equal to 10
	fmt.Println("i is equal to", i)

	fmt.Println("incremented using value receiver:")
	i.inc("passed receiver is a base value, however, received as a pointer.") // same as (&i).inc().
	                                                                         // method expects a pointer, but this still works since method can accept base value.
	fmt.Printf("1st increment. method with value as receiver. val of i: %d\n", i)

	fmt.Println("incremented using pointer receiver:")
	(&i).inc("passed receiver is a pointer.") // and this, of course, works as expected.
	fmt.Printf("2nd increment. method with pointer as receiver. val of i: %d\n", i)

	p := &i //p is a pointer to i

	fmt.Println("on stdout with pointer receiver:")
	p.print("passed receiver is a pointer, however, received as a base value.") // same as (*p).print().
	                                                                            // method expects a value, but this still works since method can accept a
	                                                                            // pointer to the base value.
	fmt.Println("on stdout with value receiver:")
	i.print("receiver is base value.") // and this, of course, works as expected.
}
