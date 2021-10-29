package main

import (
	"fmt"
	"log"
	"onelab_hw3/atoi"
	"onelab_hw3/fibonacci"
	"onelab_hw3/formatter"
	"onelab_hw3/itoa"
	"onelab_hw3/reverse"
	"onelab_hw3/rune_by_index"
)

func main() {

	// testing itoa.Itoa(int) string

	test1Cases := []int{123, -123, 2147483647, 2147483648, -2147483648, -2147483649}

	for _, test1Case := range test1Cases {
		str := itoa.Itoa(test1Case)
		fmt.Println(str)
	}

	// testing atoi.Atoi(string) (int, error)

	test2Cases := []string{"123", "0", "-0", "-1", "-123", "-555", "ababab", "1231asdasd-", "1-456"}

	for _, test2Case := range test2Cases {
		n, err := atoi.Atoi(test2Case)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(n)
	}

	// testing reverse.Reverse(string) string

	test3Cases := []string{"asd", "", "Palindrome", "kek"}

	for _, test3Case := range test3Cases {
		s := reverse.Reverse(test3Case)
		fmt.Println(s)
	}

	formatter.Format("fake_main.go")

	generator := fibonacci.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print(generator(), " ")
	}
	fmt.Println()

	s := "bla"
	i := 3

	r, err := rune_by_index.RuneByIndex(&s, &i);
	if err != nil {
		log.Printf("%s", err.Error())
	}

	fmt.Println(r)
}