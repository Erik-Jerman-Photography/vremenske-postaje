// Package vreme omogoca zapisovanje/belezenje temperaturnih meritev vremenskih postaj in izracun statistike za le-te
// Postaje so v slovarju. Ključ je oznaka postaje ("LJ01")
// Postaje ustrvarjamo s funkcijo [DodajPostajo], mertitve pa s [DodajMeritev]
package vreme

import "fmt"

// Vremenska postaja z imenom, krajem in seznamom temperatur v C
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// DodajPostajo v slovar doda novo mpostajo z oznako, imenom in krajem. Nima meritev. Če že obstaja, izpiše opozorilo.
func DodajPostajo(postaje map[string]Postaja, oznaka, ime, kraj string) {
	if _, ok := postaje[oznaka]; ok {
		fmt.Println("Postaja že obstaja.")
		return
	}
	postaje[oznaka] = Postaja{ime: ime, kraj: kraj}
}

// DodajMeritev postaji z to oznako doda meritev temperature.
// Meritev je veljavna, če je temperatura med minTemp in maxTemp (default -50, 50). Če ni veljavna ali postaja ne obstaja, izpiše sporočilo in meritve ne doda.
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64, minTemp float64, maxTemp float64) {
	if temperatura < minTemp || temperatura > maxTemp {
		fmt.Println("Temperatura ni veljavna.")
		return
	}

	p, ok := postaje[oznaka]
	if !ok {
		fmt.Println("Postaja ne obstaja.")
		return
	}

	p.meritve = append(postaje[oznaka].meritve, temperatura)
	postaje[oznaka] = p

}

// statistika vrne najnizjo, najvisjo in povprečno temperaturo postaje. Če ne obstaja ali ima manj kot stMeritev meritev, je false
func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {

	p, ok := postaje[oznaka]
	if !ok {
		return 0, 0, 0, false
	}

	if len(p.meritve) < 3 {
		return 0, 0, 0, false
	}

	min := p.meritve[0]
	max := p.meritve[0]
	vs := 0.0

	for _, m := range p.meritve {
		if m < min {
			min = m
		}
		if m > max {
			max = m
		}
		vs += m
	}

	povprecje := vs / float64(len(p.meritve))

	return min, max, povprecje, true
}

// IzpisMeritev izpise vse postaje in njihove meritve
func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")
	for oznaka, _ := range postaje {
		fmt.Printf("%s - %s (%s): %.1f\n", oznaka, postaje[oznaka].ime, postaje[oznaka].kraj, postaje[oznaka].meritve)
	}
}

// IzpisPorocila za vsako postajo izpise min, max in avg temperaturo in opis vremena.
// Izracuna se le za postaje, ki imajo vsaj stMeritev meritev.
func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	for oznaka, _ := range postaje {
		var min, max, avg, ok = statistika(postaje, oznaka, stMeritev)
		opis := ""
		switch {
		case avg >= 25.0:
			opis = "Vroče"
		case avg < 25.0 && 15.0 <= avg:
			opis = "Toplo"
		case avg < 15.0 && 5.0 <= avg:
			opis = "Hladno"
		case avg < 5.0:
			opis = "Mrzlo"
		}
		if !ok {
			fmt.Printf("%s: premalo meritev\n", postaje[oznaka].ime)
		} else {
			fmt.Printf("%s: min %.1f °C, max %.1f °C, povprečje %.1f °C -> %s\n", postaje[oznaka].ime, min, max, avg, opis)
		}
	}
}
