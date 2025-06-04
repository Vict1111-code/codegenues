package main

import "github.com/01-edu/z01"

func main() {
	for r := 'a'; r <= 'z'; r++ {
		if r == 'a' {
			z01.PrintRune(r)
		} else if r == 'e' {
			z01.PrintRune(r)
		} else if r == 'i' {
			z01.PrintRune(r)
		} else if r == 'o' {
			z01.PrintRune(r)
		} else if r == 'u' {
			z01.PrintRune(r)
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
