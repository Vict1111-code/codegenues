package main

import "github.com/01-edu/z01"

func main() {
	for r := 'A'; r <= 'Z'; r++ {
		if r == 'A' {
			z01.PrintRune(r)
		} else if r == 'E' {
			z01.PrintRune(r)
		} else if r == 'I' {
			z01.PrintRune(r)
		} else if r == 'O' {
			z01.PrintRune(r)
		} else if r == 'U' {
			z01.PrintRune(r)
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
