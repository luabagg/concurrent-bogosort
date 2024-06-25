package concurrent_bogosort_test

import (
	"fmt"

	coroutine_bogosort "github.com/luabagg/concurrent-bogosort"
)

func ExampleSort() {
	var slice = []int{12, 5, 22}

	fmt.Printf("Current slice: %v\n", slice)

	res, err := coroutine_bogosort.Sort(slice)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("Ordered slice: %v\n", res)
	}

	// Output:
	// Current slice: [12 5 22]
	// Ordered slice: [5 12 22]
}
