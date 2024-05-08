package main

import (
	"flag"
	"fmt"
	"math/rand"
	"slices"
)

func main() {
	var ilośćRund int
	flag.IntVar(&ilośćRund, "rundy", 5, "Ilość rund")
	var strategiaZmiany bool
	flag.BoolVar(&strategiaZmiany, "zamiana", false, "Strategia zmiany wyboru")
	var ilośćPudełek int
	flag.IntVar(&ilośćPudełek, "pudełka", 3, "Ilość pudełek")
	var ilośćOtwartych int
	flag.IntVar(&ilośćOtwartych, "otwarte", 1, "Ilość pudełek do otwarcia przez przez prowadzacego")
	flag.Parse()

	fmt.Printf("%d rund, zamiana: %t, %d pudełek\n", ilośćRund, strategiaZmiany, ilośćPudełek)

	symulacja(ilośćRund, ilośćPudełek, strategiaZmiany, ilośćOtwartych)
}

func symulacja(ilośćRund int, ilośćPudełek int, strategiaZmiany bool, ilośćOtwartych int) {
	var ilośćWygranych int

	for i := 0; i < ilośćRund; i++ {
		wynikRundy := runda(int(ilośćPudełek), strategiaZmiany, ilośćOtwartych)
		if wynikRundy {
			ilośćWygranych += 1
		}
	}
	var procentWygranych float32 = float32(ilośćWygranych) / float32(ilośćRund) * 100
	fmt.Printf("Wygrano %d z %d rund, przegrano %d rund. Procent wygranych: %.2f\n", ilośćWygranych, ilośćRund, ilośćRund-ilośćWygranych, procentWygranych)

}

func runda(ilośćPudełek int, zamiana bool, ilośćOtwartych int) bool {
	pudełkoZWygraną := rand.Intn(ilośćPudełek)
	wybranePudełko := rand.Intn(ilośćPudełek)
	var otwartePudełka = make([]int, ilośćOtwartych)

	for i := 0; i < ilośćOtwartych; i++ {
		wylosowanaLiczba := rand.Intn(ilośćPudełek)
		for slices.Contains(otwartePudełka, wylosowanaLiczba) ||
			wylosowanaLiczba == wybranePudełko ||
			wylosowanaLiczba == pudełkoZWygraną {
			wylosowanaLiczba = rand.Intn(ilośćPudełek)
		}
		otwartePudełka[i] = wylosowanaLiczba
	}
	if zamiana {
		nowyWyborPudełka := wybranePudełko
		for nowyWyborPudełka == wybranePudełko || slices.Contains(otwartePudełka, nowyWyborPudełka) {
			nowyWyborPudełka = rand.Intn(ilośćPudełek)
		}
		wybranePudełko = nowyWyborPudełka
	}
	return wybranePudełko == pudełkoZWygraną
}
