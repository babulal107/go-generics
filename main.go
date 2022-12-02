package main

import "fmt"

// Generics allow types to be parameters
// Go 1.18+

// 3 main features :
// 1. Type parameters ( with constraint)
// 2. Type inference
// 3. Type set

func main() {
	//fmt.Println(" min : ", minInt(1, 2))
	//fmt.Println(" min : ", minFloat(0.4, 0.2))

	//fmt.Println(" min : ", min[int](1, 2))
	fmt.Println(" min : ", min[float64](0.4, 0.2))

	type supFloat float64
	var sf supFloat = 2.5
	fmt.Println(" min : ", minWithTypeInference(sf, 0.2))
	fmt.Println(" min : ", minWithTypeSet(sf, 1))

	//PrintGenericData[int](1, 2, 3, 4)
	//PrintGenericData[float64](1.0, 0.2, 3.4, 4.0)
	//PrintGenericData[string]("a", "b", "c", "d")

}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minFloat(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// 1. Type parameters
//
//func PrintGenericData[T interface{}](a ...T) { // print any type of data
//	fmt.Println(a)
//}

//interface/any type of constraint not checking underline data type that's why it's can accept any data type values.
//func min[T interface{}](a, b T) T {
//	//if a < b {  // here it will through compile time error And need to convert specific type by using reflection then compare value
//	//	return a
//	//}
//	return a
//}

// constraint check underline data type not allow if we pass supFloat which have underlined float64
func min[T float64 | int](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// 2. Type inference
//
// constraint check underline data type also allow to pass supFloat with type inference (~)
func minWithTypeInference[T ~float64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// 3. Type set
//
type typeSet interface {
	~float64 | int
}

// minWithTypeSet... it's checking underline data type & accept float64 & int both value
func minWithTypeSet[T typeSet](a, b T) T {
	if a < b {
		return a
	}
	return b
}
