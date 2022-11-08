// package main

// import "fmt"

// func main() {

// 	// var testMap map[string]int
// 	// testMap = map[string]int{
// 	// 	"one":   1,
// 	// 	"two":   2,
// 	// 	"three": 3,
// 	// }

// 	// k := "two"
// 	// v, ok := testMap[k]
// 	// if ok {
// 	// 	fmt.Printf("The element of key %q: %d\n", k, v)
// 	// } else {
// 	// 	fmt.Println("Not found!")
// 	// }

// 	m := make(map[string]string)
// 	m["name"] = "Sherry"
// 	m["address"] = "address"
// 	fmt.Printf("m: %v\n", m)

// 	value, exists := m["name"]

// 	for key, value := range m {
// 		fmt.Printf("key: %v \t key: %v\n", key, value)
// 	}

// 	if exists {
// 		fmt.Printf("name-value: %v\n", value)
// 	} else {
// 		fmt.Printf("name-value: %v\n", value)
// 	}
// 	delete(m, "address")
// 	fmt.Printf("m: %v\n", m)
// }
