package main

import "fmt"

func main() {
	//key is TYPE string and value is a slice of TYPE string
	m := make(map[string][]string)
	m = map[string][]string{
		"bond_james":      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	//fmt.Println(m)
	// for key, value := range m {
	// 	fmt.Println("this is a record for ", key)
	// 	for i, v2 := range value {
	// 		fmt.Println("\t \t", i, v2)
	// 	}
	// }

	m["goldman"] = []string{"oppenheimer", "roosevelt", "jackson"}

	// for k, v := range m {
	// 	fmt.Println("The key record is : ", k)
	// 	for i, v2 := range v {
	// 		fmt.Println("\t\t", i, v2)
	// 	}
	// }

	if v, ok := m["no_dr"]; ok {
		fmt.Println("Deleting the following record values ", v)
		delete(m, "no_dr")
	}

	for k, v := range m {
		fmt.Println("The key record is : ", k)
		for i, v2 := range v {
			fmt.Println("\t\t", i, v2)
		}
	}

}
