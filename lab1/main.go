package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	var ilośćRund int
	flag.IntVar(&ilośćRund, "rundy", 5, "Ilość rund")
	var strategiaZmiany bool
	flag.BoolVar(&strategiaZmiany, "zamiana", false, "Strategia")
	var ilośćPudełek int
	flag.IntVar(&ilośćPudełek, "pudełka", 3, "Ilość pudełek")
	flag.Parse()

	fmt.Printf("%d rund, zamiana: %t, %d pudełek\n", ilośćRund, strategiaZmiany, ilośćPudełek)

	symulacja(ilośćRund, ilośćPudełek, strategiaZmiany)
}

func symulacja(ilośćRund int, ilośćPudełek int, strategiaZmiany bool) {
	var ilośćWygranych int

	for i := 0; i < ilośćRund; i++ {
		wynikRundy := runda(int(ilośćPudełek), strategiaZmiany)
		if wynikRundy {
			ilośćWygranych += 1
		}
	}
	var procentWygranych float32 = float32(ilośćWygranych) / float32(ilośćRund) * 100
	fmt.Printf("Wygrano %d z %d rund, przegrano %d rund. Procent wygranych: %.2f\n", ilośćWygranych, ilośćRund, ilośćRund-ilośćWygranych, procentWygranych)

}

func runda(ilośćPudełek int, zamiana bool) bool {
	pudełkoZWygraną := rand.Intn(ilośćPudełek)
	wybranePudełko := rand.Intn(ilośćPudełek)
	otwartePudełko := rand.Intn(ilośćPudełek)
	for otwartePudełko == pudełkoZWygraną || otwartePudełko == wybranePudełko {
		otwartePudełko = rand.Intn(ilośćPudełek)
	}
	if zamiana {
		nowyWyborPudełka := wybranePudełko
		for nowyWyborPudełka == wybranePudełko || nowyWyborPudełka == otwartePudełko {
			nowyWyborPudełka = rand.Intn(ilośćPudełek)
		}
		wybranePudełko = nowyWyborPudełka
	}
	return wybranePudełko == pudełkoZWygraną
}
