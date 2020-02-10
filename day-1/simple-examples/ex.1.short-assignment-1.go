package main

import (
	"fmt"
)


func f1(isOK bool, id int, mac string) (bool, int, string) {
	return isOK, id, mac
}

func f2(isOK bool, id int) (bool, int) {
	return isOK, id
}

func main() {
	//fmt.Printf("without backspace: \tthis's a simple text\n\n")
	//fmt.Printf("with backspace: \bthis's a simple text\n\n")

	/* isOK1 := false
	id1 := 10
	mac1 := "00:00" */

	msg := "[000]: all 3 new return-value-variable."
	isOK1, id1, mac1 := f1(true, 100, "10:00")
	fmt.Println(msg)
	fmt.Printf("[000]:  isOK1: %t,  id1: %d,  mac1: %s\n\n", isOK1, id1, mac1)

	// new 3rd return-value-variable
	// 1st and 2nd return-value-variables aren't recreated and nor compiler issues error of redefinition of these variables.
	// whilst new 3rd return-value-variable is created, value of each of the existing variables, 1st and 2nd in this case, has been updated.
	msg = `[001]: new 3rd return-value-variable
1st and 2nd return-value-variables aren't recreated and nor compiler issues error of redefinition of these variables.
whilst new 3rd return-value-variable is created, value of each of the existing variables, 1st and 2nd in this case, has been updated.`
	isOK1, id1, mac2 := f1(false, 200, "10:01")
	fmt.Println(msg)
	fmt.Printf("[001]:  isOK1: %t,  id1: %d,  mac2: %s\n\n", isOK1, id1, mac2)

	// new 2nd return-value-variable
	// 1st and 3rd return-value-variables aren't recreated and nor compiler issues error of redefinition of these variables.
	// whilst new 2nd return-value-variable is created, value of each of the existing variables, 1st and 3rd in this case, has been updated.
	msg = `[010]: new 2nd return-value-variable
1st and 3rd return-value-variables aren't recreated and nor compiler issues error of redefinition of these variables.
whilst new 2nd return-value-variable is created, value of each of the existing variables, 1st and 3rd in this case, has been updated.`
	isOK1, id2, mac2 := f1(true, 300, "10:10")
	fmt.Println(msg)
	fmt.Printf("[010]:  isOK1: %t,  id2: %d,  mac2: %s\n\n", isOK1, id2, mac2)

	// new 2nd and 3rd return-value-variables
	// 1st return-value-variables isn't recreated and nor compiler issues error of redefinition of this variable.
	// whilst new 2nd and 3rd return-value-variables are created, value of existing variable, 1st in this case, has been updated.
	msg = `[011]: new 2nd and 3rd return-value-variables
1st return-value-variables isn't recreated and nor compiler issues error of redefinition of this variable.
whilst new 2nd and 3rd return-value-variables are created, value of existing variable, 1st in this case, has been updated.`
	isOK1, id3, mac3 := f1(false, 400, "10:11")
	fmt.Println(msg)
	fmt.Printf("[011]:  isOK1: %t,  id3: %d,  mac3: %s\n\n", isOK1, id3, mac3)

	// new 1st return-value-variable
	// 2nd and 3rd return-value-variables aren't recreated and nor compiler issues error of redefinition of these variables.
	// whilst new 1st return-value-variable is created, value of each of the existing variables, 2nd and 3rd in this case, has been updated.
	msg = `[100]: new 1st return-value-variable
2nd and 3rd return-value-variables aren't recreated and nor compiler issues error of redefinition of these variables.
whilst new 1st return-value-variable is created, value of each of the existing variables, 2nd and 3rd in this case, has been updated.`
	isOK2, id3, mac3 := f1(true, 500, "11:00")
	fmt.Println(msg)
	fmt.Printf("[100]:  isOK2: %t,  id3: %d,  mac3: %s\n", isOK2, id3, mac3)

	isOK100, id100 := f2(false, 200)
	fmt.Printf("f2(false, 200)(new, new):  isOK100: %t,  id100: %d\n", isOK100, id100)

	isOK100, id200 := f2(true, 300)
	fmt.Printf("f2(true, 300)(old, new):  isOK100: %t,  id200: %d\n", isOK100, id200)

	isOK200, id200 := f2(true, 400)
	fmt.Printf("f2(true, 400)(new, old):  isOK200: %t,  id200: %d\n", isOK200, id200)

	return
}
