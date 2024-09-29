package main

import (
	"fmt"
	"sort"
	"time"
)

func reverseVariable(a, b string) (string, string) {
	return b, a
}

func simpleGoRoutine(name string) {
	fmt.Println("This is a sample Go routine" + name)
}

var a, b, c = 1, 1.1, "Baby"

func main() {
	fmt.Println("Hello, Go!")
	v1, v2 := "Dileep", "Pandey"
	fmt.Println(reverseVariable(v1, v2))
	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println(a, b, c)
	go simpleGoRoutine("Dileep")
	fmt.Println("Main go routine started")
	time.Sleep(2 * time.Second)
	fmt.Println("Main go routine execution completed")
	arr := [6]string{"seema", "baby", "komal", "ab", "bc", "cd"}
	for _, a := range arr {
		fmt.Println(a)
	}

	triangleObj := Triangle{base: 23, height: 12}
	sqrObj := Square{length: 12}
	rectangleObj := Rectangle{breadth: 12, length: 25}

	shapes := [3]Shape{triangleObj, sqrObj, rectangleObj}

	for _, shape := range shapes {
		fmt.Println(shape.area())
	}

	humans := []Human{
		{name: "dileep", age: 22},
		{name: "seema", age: 27},
		{name: "komal", age: 23},
	}
	sort.Sort(ageFactor(humans))
	fmt.Println(humans)
}
