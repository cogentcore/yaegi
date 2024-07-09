package main

func main() {
	foos := []func(){}

	for i := range 3 {
		a, b := i, i
		foos = append(foos, func() { println(i, a, b) })
	}
	foos[0]()
	foos[1]()
	foos[2]()
}

// Output:
// 0 0 0
// 1 1 1
// 2 2 2
