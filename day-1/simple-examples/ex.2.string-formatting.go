// some examples of common string formatting in fmt.Printf()

package main

import "fmt"
import "os"

type point struct {
	x, y int
}

func main() {

	// go offers several printing "verbs", what commonly known as conversion specifications.
	// following are some examples that'll come handy.
	p := point{1, 2}
	fmt.Printf("%v\n", p)

	// "%+v": for struct's field names.
	fmt.Printf("p1: %+v\n", p)

	// "%#v": prints Go's syntax representation of the value, i.e, category of the type (variable/map/array/function/slice/interface{}/pointer, et al.)
	// of course, value.
	fmt.Printf("p2: %#v\n", p)

	// "%T": type of a value.
	fmt.Printf("type(p): %T\n", p)

	// "%t: boolen type.
	fmt.Printf("formatting booleans in straight-forward: %t\n", true)

	// formatting integers:
	// "%d": standard base-10 formatting.
	fmt.Printf("%d\n", 123)

	// "%b": binary representation.
	fmt.Printf("%b\n", 14)

	// "%c": character that corresponds to the given integer.
	fmt.Printf("char(33): %c\n", 33)

	// "%x": hex encoding.
	fmt.Printf("%x\n", 456)

	// formatting options for floats.
	// "%f": basic decimal formatting point number.
	fmt.Printf("%f\n", 78.9)

	// "%e" and "%E": e notation.
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	// "%s": strings.
	fmt.Printf("string-1: %s\n", "\"string\"")

	// escaping using \:
	str := "string"
	fmt.Printf("string-2: \"%s\"\n", str)

	// "%q": double-quote strings as in Go source.
	fmt.Printf("string-3: %q\n", "\"string\"")

	// "%x": base-16, 2 hex characters per byte.
	fmt.Printf("%x\n", "hex this")

	// "%p": representation of a pointer.
	fmt.Printf("pointer-1: %p\n", &p)

	// width and precision of numbers: result is right-justified and
	// padded with spaces.
	fmt.Printf("numbers-1: |%6d|%6d|\n", 12, 345)
	fmt.Printf("numbers-2: %6d|%6d\n", 12, 345)
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// '-': left-justified o/p.
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// number of characters to be printed in strings. o/p right justified by default.
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	// '-': left-justified o/p.
	fmt.Printf("left-justified-string: |%-6s|%-6s|\n", "foo", "b")

	// fmt.Sprintf() formats and returns a string to the left-hand-side string variable.
	s := fmt.Sprintf("a %s", "string")
	fmt.Println("s:",s)

	// fmt.Fprintf() formats and prints to io.Writers other than os.Stdout.
	fmt.Fprintf(os.Stderr, "on stderr: an %s\n", "error")
}
