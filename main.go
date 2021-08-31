package main

import "fmt"

func main() {
	// string variable
	// var myName string = "Eddy" // if a variable is not being used it causes a error and won't be compiled
	// var nameWithoutType = "Test Name"
	// var settingUpVariable string

	// fmt.Println(myName, nameWithoutType, settingUpVariable)

	// myName = "Name changed to YddE"
	// nameWithoutType = "Is this still a string?"
	// settingUpVariable = "Using the variable"
	// fmt.Println(myName, nameWithoutType, settingUpVariable)

	// // shorthand

	// nameGo := "Gopher" // you can only use this type of variable declaration inside of variables
	// fmt.Println(nameGo)

	//-----------------------------------
	// intagers and floats
	// var age int = 100
	// var ageTwo = 200
	// ageThree := 300
	// fmt.Println(age, ageTwo, ageThree)

	// you scan specify the int bits
	// uint8       the set of all unsigned  8-bit integers (0 to 255)
	// uint16      the set of all unsigned 16-bit integers (0 to 65535)
	// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
	// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

	// int8        the set of all signed  8-bit integers (-128 to 127)
	// int16       the set of all signed 16-bit integers (-32768 to 32767)
	// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
	// int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
	// var num8Bit int8 = 21
	// var numThee uint // can't have a negative number

	// float
	var scoreOne float32 = 25.39
	fmt.println(scoreOne)

}
