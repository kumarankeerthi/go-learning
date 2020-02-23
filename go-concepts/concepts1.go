package main

import (
	// all imported packages must be used atleast once, e
	//lse go will throw compile time error
	"fmt"
	"time"
)

//global variable

// cannnot have special char in variable name but can have "_" like num_items
//var gobal-local-variable string = "Gobal Variable"
var goballocalvariable string = "Gobal Variable"

func main() {

	// single line comments
	/*
		multti line
		comments
	*/

	// three ways to declare a variable

	// var k int=4;
	//var k = 4
	i := 4
	s := "string-var"
	fmt.Printf("value is %d  and type is  %T\n", i, i)
	fmt.Printf("value is %s and type is %T\n", s, s)

	fmt.Println("The time is : ", time.Now())
	fmt.Printf("The time  is %s\n", time.Now())

	// can evaluate expression within print()
	fmt.Println(4 + 4)
	fmt.Println(1.09 - 5)
	fmt.Println(true)
	fmt.Println(5 == 6)
	fmt.Println("sum of 1 and 1 is :", 1+1)

	//*******  Variables ********

	// local variable with same variable name takes precedence
	// here since it not yet initialized within main, it will print the global value
	fmt.Println(goballocalvariable)
	// once declared within main, it will print the local value
	var goballocalvariable string = "local Variable"
	fmt.Println(goballocalvariable)

	// this pattern is same when same variable name is referenced in two functions
	// the value within a given function takes preceedence over the same variable value from a calling func

	// we can alos declare multiple variale of same type at once

	//var f1, f2 float64 = 9.9, 5.5
	f1, f2 := 9.9, 5.5
	fmt.Println(f1, f2)
	fmt.Printf("value of f1 is %f and type is %T\n", f1, f1)

	// ******** Constants ******

	const pi float64 = 3.1415926

	fmt.Println(pi)
	// cannot change const value - it will throw erro
	// pi = 9.11

	// Note: in go, like imports , even vairable that is declared has to used, else we get compile error

	// declaring a varible as one type and then reassinging it with different type value

	var var1 int = 10
	fmt.Println(var1)
	// below will error  but same will work is it is "  var1= 9.0  since it is same as 9 (int)"
	//var1 = 9.9    --- constant 9.9 truncated to integer
	//var1 = "test" -- cannot use "test" (type string) as type int in assignment
	fmt.Println(var1)

	// ******** Loops ******

	// only key work for loops in go is "for"  nothing else

	//  approach#1
	loopvar1 := 100

	for loopvar1 <= 110 {
		fmt.Println(loopvar1)
		loopvar1++
	}

	//appraoch#2

	for loopvar2 := 200; loopvar2 <= 210; loopvar2++ {
		fmt.Println(loopvar2)

	}

	for loopvar3 := 0; loopvar3 <= 10; loopvar3++ {
		if loopvar3%2 == 0 {
			fmt.Println(loopvar3)
		}
	}

	// infitite loop will be like below
	/*
		for {
			-- do something for ever
		}
	*/

	// break can be used to come out of a loop.
	for loopvar3 := 0; loopvar3 <= 20; loopvar3++ {
		if loopvar3 > 15 {
			break
		}

		if loopvar3%2 == 0 {
			fmt.Println(loopvar3)
		}
	}

	//  *******  IF and ELSE   ************

	ifvar1 := 5
	if ifvar1%2 == 0 {
		fmt.Println(ifvar1, " is even")
	} else {
		fmt.Println(ifvar1, " is odd")
	}

	ifvar2 := 500
	if ifvar2 < 50 {
		fmt.Println(ifvar2, " is less than 50")
	} else if ifvar2 > 50 {
		fmt.Println(ifvar2, " is greater than 50")
	}

	// *****  SWITCH  *****

	switchvar := 20
	switch switchvar {
	case 1:
		fmt.Println(" value is 1")
	case 2, 3, 4, 5:
		fmt.Println(" value is 2")
	// note:  we can have case to evaluate expression as well.
	case 10 + 10:
		fmt.Println(" value is 20")
	default:
		fmt.Println(" value is default")

	}

	// another variation can be like this
	switchvar1 := 10
	switch {
	case switchvar1 == 10:
		fmt.Println(" value is 10")
	}

	// *********** ARRAYS **********

	var int_arr [5]int // if not initialized default value will be 0
	var bool_arr [10]bool
	var str_arr [20]string // blanks as default value
	fmt.Println(int_arr)
	fmt.Println(bool_arr)
	fmt.Println(str_arr)

	// chaning value of a array at an index - go is zero based -- starts with "0"
	int_arr[0] = 100
	int_arr[1] = 200
	fmt.Println(int_arr)

	fmt.Println(int_arr[1])

	int_arr1 := [5]int{1, 2, 3, 5, 6}
	fmt.Println(int_arr1)

	//NOTE : if we dont assign the array with complete set of value, the remiaing will be defaulted
	int_arr1 = [5]int{1, 2, 3, 5}
	fmt.Println(int_arr1)

	// getting the lenght of af an arry is by using len() func
	fmt.Println("lenght of array int_arr is", len(int_arr1))

	// two dimentional array

	var aa [5][5]int
	fmt.Println(aa)

	count := 1
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			aa[i][j] = count
			count++
		}
	}
	fmt.Println(aa)

	// *********** SLICES ***********
	slice := make([]string, 3)
	fmt.Println("silce value:", slice)
	slice[0] = "one"
	slice[1] = "two"
	slice[2] = "three"
	fmt.Println("silce value after assingment: ", slice, "and its lenght is", len(slice))
	fmt.Println("2nd value on this slice is", slice[1])

	//append func on slice

	slice = append(slice, "four")
	fmt.Println("value of slice after append is", slice, " and the new lenght is ", len(slice))

	newSlice := make([]string, len(slice))
	copy(newSlice, slice)
	fmt.Println("value of newSlice after copying from slice is", newSlice)
	slice = append(slice, "five", "six")
	fmt.Println("new value of slice is", slice)
	fmt.Println("value of newSlice after altering slice is", newSlice)

	// various ways to slice a slice variable
	fmt.Println("value at slice[0]", slice[0])
	fmt.Println("value at slice[1]", slice[1])
	fmt.Println("value at slice[2]", slice[2])
	fmt.Println("value at slice[3]", slice[3])
	fmt.Println("value at slice[4]", slice[4])
	fmt.Println("value at slice[5]", slice[5])

	temp := slice[2:4] // this gives the value at 2nd and 3rd index of this slice
	fmt.Println("value of temp after slicing 2:4 is", temp)

	temp = slice[:4] // this gives values till the 4th index, excluing the 4th index value
	fmt.Println("value of new temp using [:4] is", temp)

	temp = slice[3:]
	fmt.Println("value of slice using slice[3:]", temp)

	//We can declare and initialize a variable for slice in a single line as well

	slice1 := []string{"a", "b", "c", "d"}
	fmt.Println("value of slice1 is", slice1)

	//slice can be 2d slices simialr to 2d arrays
	//finding , varaible cannot start with numbers :)
	//2D_Slice := make([]int,5)
	twoD_Slice := make([][]int, 5)
	for i := 0; i < 5; i++ {
		twoD_lenght := i + 3
		twoD_Slice[i] = make([]int, twoD_lenght)
		for j := 0; j < twoD_lenght; j++ {
			twoD_Slice[i][j] = i + j
		}
	}

	fmt.Println("value of 2D slice is ", twoD_Slice)

	// ********** MAPS *******

	gomap := make(map[string]int)
	gomap["a"] = 1
	gomap["b"] = 2
	fmt.Println("value of map is", gomap)

	val1, exists := gomap["a"]
	fmt.Println("key and value of gomap[\"a\"] is", val1, exists)
	val1, exists = gomap["z"]
	fmt.Println("key and value of gomap[\"a\"] is", val1, exists)

	delete(gomap, "b")
	fmt.Println("value of gomap[\"b\" after deletion is", gomap)

	// we can also initialze map directly like this
	gomap1 := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(gomap1)

	// ******** RANGE *******

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := 0

	for x, num := range nums {
		sum += num
		fmt.Println(x) // Note, the range of smething returns two values, index and the values itself
	}
	fmt.Println(sum)

	for key, value := range gomap1 {
		fmt.Println("key and value of gomap1 are", key, value)
	}

	// if we just accept one value form range then it will only return the key/index of map/arr

	for x := range nums {
		fmt.Println(x)
	}
	for key := range gomap1 {
		fmt.Println("key of gomap1 are", key)
	}

	// we can iterate ovet string as well - range will retunr the rune of each chars in the string
	for i, c := range "keerthi" {
		fmt.Println("value of i and c is ", i, c)
	}
}
