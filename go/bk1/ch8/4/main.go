package main

import (
	"errors"
	"fmt"
	"math/big"
	// "os"
)


func Factorial(n int,f func(*big.Int,*big.Int)big.Int) (big.Int, error) {
	acc := big.NewInt(1)
	def0 := big.NewInt(0)

	if n == 0 || n == 1 {
		return *acc, nil
	}
	if n < 0 {
		return *def0, errors.New("values less than zero unacceptable")
	}

	for i:=1; i<=n;i++ {
		*acc = f(acc, big.NewInt(int64(i)))
	}

	return *acc, nil
}

func main() {
	var n int
	fmt.Println("Get factorial of a whole number(Integer) greater or equal to zero.")
	fmt.Println("Note: This will not handle numbers with fractional parts")
	fmt.Println("")

	for {
		fmt.Println("Please enter your number:")
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		if n < 0 {
			fmt.Println("Error reading input:", " Numbers less than zero unaccepted")
			continue
		}
		break
	}

	fact, err := Factorial(n, func(u1, u2 *big.Int) big.Int {
		defM := big.NewInt(1)

		return *(defM.Mul(u1, u2))
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Factorial of ",n, " is ", fact.Text(10))
}