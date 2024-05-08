package main

import (
	"fmt"
	"lab2/funkcje"
	"math"
	"sort"
	"strings"
)

func main() {
	var imie string
	var nazwisko string
	fmt.Println("Podaj imię i nazwisko, bez polskich znaków:")
	fmt.Scanln(&imie, &nazwisko)
	nick := []byte(strings.ToLower(imie[:3]) + strings.ToLower(nazwisko[:3]))
	fmt.Println(fmt.Sprint(nick))
	fmt.Printf("Twój nick to %s\n", nick)

	silnaLiczba := znajdzSilnaLiczbe(nick)
	fmt.Printf("Silna liczba dla nicku %s to %d\n", nick, silnaLiczba)

	słabaLiczba := znajdzSlabaliczbe(silnaLiczba)
	fmt.Printf("Najbliższa silnej liczbie słaba liczba to %d (Fib(%d) wywołana %d razy)\n", słabaLiczba, słabaLiczba, funkcje.LiczbaWywołańFib[słabaLiczba])

	fmt.Printf("Całkowita liczba wywołań fib: %d\n", funkcje.CałkowitaLiczbaWywołańFib)
	{
		keys := make([]int, 0, len(funkcje.LiczbaWywołańFib))
		for k := range funkcje.LiczbaWywołańFib {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		for _, k := range keys {
			v := funkcje.LiczbaWywołańFib[k]
			fmt.Printf("fibonacci(%d) wywołano %d razy\n", k, v)
		}
	}
	// fmt.Printf("\n------------ Performance test ------------\n\n")
	// funkcje.TestFib()
	// funkcje.MeanFib()
}

func znajdzSlabaliczbe(silnaLiczba int) int {
	funkcje.Fibonacci(30)
	słabaLiczba := 0
	for k, v := range funkcje.LiczbaWywołańFib {
		if math.Abs(float64(silnaLiczba-v)) < math.Abs(float64(silnaLiczba-funkcje.LiczbaWywołańFib[słabaLiczba])) {
			słabaLiczba = k
		}
		// fmt.Printf("fibonacci(%d) wywołano %d razy\n", k, v)
	}
	return słabaLiczba
}

func znajdzSilnaLiczbe(nick []byte) int {
	i := 0
Outer:
	for {
		potencjalna_silna := funkcje.Silnia(i)
		for _, v := range nick { // sprawdzamy czy wszystkie litery nicku są w silni
			if !strings.Contains(potencjalna_silna.String(), fmt.Sprint(v)) {
				i += 1
				continue Outer
			}
		}
		return i
	}
}
