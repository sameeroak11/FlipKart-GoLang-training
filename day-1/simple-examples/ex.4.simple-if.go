package main


import (
	"fmt"
)


func main() {
	/* flag := true

	//if flag == false {
	if !flag {
		fmt.Printf("this's it.\n")
	} else {
		fmt.Printf("this'sn't it.\n")
	}

	flag = false
	switch flag {
		case flag:
			fmt.Printf("swith this's it.\n")
			break
		case flag == false:
			fmt.Printf("swith this'sn't it.\n")
			break
		default:
			fmt.Printf("no default for bool.\n")
			break
	}

	switch {
		case flag:
			fmt.Printf("swith this's it.\n")
			break
		case !flag:
			fmt.Printf("swith this'sn't it.\n")
			break
		default:
			fmt.Printf("no default for bool.\n")
			break
	} */

	n := 2
	mod := 10

	if (n % mod) == 1 {
		fmt.Printf("case 1: n: %d,  mod: %d,  val: %d\n", n, mod, n % mod)
	} else if (n % mod) == 2 {
		fmt.Printf("case 2: n: %d,  mod: %d,  val: %d\n", n, mod, n % mod)
	} else if (n % mod) == 3 {
		fmt.Printf("case 3: n: %d,  mod: %d,  val: %d\n", n, mod, n % mod)
	} else if (n % mod) == 4 {
		fmt.Printf("case 4: n: %d,  mod: %d,  val: %d\n", n, mod, n % mod)
	} else {
		fmt.Printf("default: n: %d,  mod: %d,  val: %d\n", n, mod, n % mod)
	}

	/* switch n % mod {
		case 1:
			fmt.Printf("case 1: n: %d,  mod: %d,  val: %d\n", n, mod, n % mod)
		case 2:
			fmt.Printf("case 2: n: %d,  mod: %d,  val: %d\n", n, mod, n % mod)
		case 3:
			fmt.Printf("case 3: n: %d,  mod: %d,  val: %d\n", n, mod, n % mod)
		case 4:
			fmt.Printf("case 4: n: %d,  mod: %d,  val: %d\n", n, mod, n % mod)
		default:
			fmt.Printf("default: n: %d,  mod: %d,  val: %d\n", n, mod, n % mod)
	} */
}
