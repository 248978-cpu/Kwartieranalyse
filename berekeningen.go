package main

import "time"

func totaal(kwartierwaarden []Kwartierwaarde, soort string) float64 {
	var tot float64
	for _, k := range kwartierwaarden {
		if k.Soort == soort {
			tot += k.Waarde
		}
	}
	return tot
}

func maxKW(kwartierwaarden []Kwartierwaarde, soort string) float64 {
	var max float64 = 0
	for _, k := range kwartierwaarden {
		if k.Soort != soort {
			continue
		}
		if k.Waarde > max {
			max = k.Waarde
		}
	}
	return max
}

type MaandPiek struct {
	Tijdstip time.Time
	Max      float64
}

func PiekenPerMaand(kwartierwaarden []Kwartierwaarde) [12]MaandPiek {
	var pieken [12]MaandPiek
	for _, k := range kwartierwaarden {
		m := k.Tijdstip.Month() - 1
		if k.Waarde*4 > pieken[m].Max {
			pieken[m] = MaandPiek{Tijdstip: k.Tijdstip, Max: k.Waarde * 4}
		}
	}
	return pieken
}

func GemiddeldeMaandPiek(pieken [12]MaandPiek) float64 {
	var totaal float64
	var n int
	for _, p := range pieken {
		if !p.Tijdstip.IsZero() { // maand komt voor in de data
			totaal += p.Max
			n++
		}
	}
	if n == 0 {
		return 0
	}
	return totaal / float64(n)
}
