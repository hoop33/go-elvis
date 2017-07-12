package main

import elvis "github.com/hoop33/go-elvis"

func main() {
	a := ""
	a = elvis.Elvis(a, "b").(string)
	println(a) // prints b

	a = "a"
	a = elvis.Elvis(a, "b").(string)
	println(a) // prints a

	p := "Hello"
	var p1, p2, p3 *string
	p1 = nil
	p2 = &p
	p3 = elvis.Elvis(p1, p2).(*string)
	println(*p3) // prints Hello

	var c string
	c = elvis.Elvis("a", "b").(string)
	println(c) // prints a

	c = elvis.Elvis(nil, "b").(string)
	println(c) // prints b

	c = elvis.Elvis("", "b").(string)
	println(c) // prints b

	var c2 int
	c2 = elvis.Elvis(1, 2).(int)
	println(c2) // prints 1

	c2 = elvis.Elvis(0, 2).(int)
	println(c2) // prints 2

	x := 7
	c = elvis.Ternary(x == 7, "a", "b").(string)
	println(c) // prints a

	c = elvis.Ternary(x == 1, "a", "b").(string)
	println(c) // prints b

	c2 = elvis.Ternary(x == 7, 0, 1).(int)
	println(c2) // prints 0

	c2 = elvis.Ternary(x == 1, 0, 1).(int)
	println(c2) // prints 1
}
