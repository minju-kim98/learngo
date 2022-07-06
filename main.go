package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	//const name string = "Minju"
	var name string = "Minju"
	totalLength, upperName := lenAndUpper("Minju")
	fmt.Println(totalLength, upperName)
	fmt.Println(lenAndUpper2(name))
	fmt.Println(supperAdd(1, 2, 3, 4, 5, 6))
	fmt.Println(canIDrink(18))
	testPointer()
	arrayTest()
	mapTest()
}

//Functions
func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

//naked return and defer
func lenAndUpper2(name string) (length int, upperCase string) {
	defer fmt.Println("I'm done!")
	length = len(name)
	upperCase = strings.ToUpper(name)
	return
}

//for, range, ...args
func supperAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func supperAdd2(numbers ...int) int {
	total := 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return total
}

//If with a Twist: variable expression
func canIDrink(age int) bool {
	if age < 21 {
		return false
	}
	return true
}
func canIDrink2(age int) bool {
	if koreanAge := age + 1; koreanAge < 18 { //variable expression
		return false
	}
	return true
}

//switch
func canIDrink3(age int) bool {
	switch {
	case age < 21:
		return false
	case age == 21:
		return true
	case age > 50:
		return false
	}
	return true
}
func canIDrink4(age int) bool {
	switch koreanAge := age + 1; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

//pointers
func testPointer() {
	a := 2
	b := &a
	fmt.Println(a, b, &a, *b)
}

//arrays and slices
func arrayTest() {
	namesArray := [5]string{"Minju", "Helena", "Tony"}
	namesSlice := []string{"Minju", "Helena", "Tony"}
	namesSlice = append(namesSlice, "Jia")
	fmt.Println(namesArray, namesSlice)
}

//maps
func mapTest() {
	minju := map[string]string{"name": "minju", "age": "25"}
	fmt.Println(minju)
	//for-loop in map
	for key, value := range minju {
		fmt.Println(key, value)
	}

	//using struct
	favFood := []string{"jerk", "blueberry"}
	//tony := person{"Tony", 21, favFood}
	//OR:
	tony := person{name: "Tony", age: 21, favFood: favFood}
	fmt.Println(tony)
}

type person struct {
	name    string
	age     int
	favFood []string
}
