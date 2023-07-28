package shared

import "fmt"

func Bucle() {

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	sum := 1
	for sum <= 10 {
		fmt.Println(sum)
		sum = sum + 1
	}

	l := 1
	for {
		fmt.Println(l)

		l = l + 1
		if l > 10 {
			break
		}
	}
}
