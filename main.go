package main

import (
	"fmt"
	"runtime"
)

func main() {
	var welcome string = "hello world"
	const personalWelcome string = "hello ejiro"
	// welcome := "hello world"
	println(welcome)
	println(personalWelcome)

	var myAge int = 28 + 2
	const pi float32 = 3.14
	// myAge := 27
	println(myAge)
	println(pi)

	var size float32 = 2.0
	println(size * pi)

	var customerHeight int = 150
	var customerAge = 17

	// if customerHeight >= 150 && customerAge >= 18{
	// 	println("can access ride")
	// } else if customerHeight >= 120 {
	// 	println("customer can access children's rides")
	// } else {
	// 	println("customer is too small")
	// }

	switch {
	case customerHeight >= 150 || customerAge >= 18:
		println("can access ride!")
	case customerHeight >= 120:
		println("can access children's rides")
	default:
		println("cannot access rides")
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		println(os)
	case "linux":
		println("linux machine")
	default:
		println("something else")
	}

	//Arrays and slices

	planets := [8]string{"mercury", "venus", "earth", "mars", "jupiter", "saturn", "uranus", "neptune"}
	fmt.Println(planets)

	var planetsArray [8]string
	planetsArray[0] = "mercury"

	fmt.Println(planetsArray)

	planetSlice := []string{"mercury", "venus", "earth", "mars", "jupiter", "saturn", "uranus", "neptune"}
	fmt.Println(planetSlice)

	var planetSliceVerbose []string
	planetSliceVerbose = append(planetSliceVerbose, "mercury")

	fmt.Println(planetSliceVerbose)

	//iterating over a slice
	ages := []int{42, 28, 30, 27, 18}

	for i := 0; i < len(ages); {
		fmt.Println(ages[i])
		i++
	}
}
