package main

func totaal(kwartierwaarden []Kwartierwaarde, soort string) float64 {
	var tot float64
	for _, k := range kwartierwaarden {
		if k.Soort == soort {
			tot += k.Waarde
		}
	}
	return tot
}

func maxPerMaand(kwartierwaarden []Kwartierwaarde) map[string]float64 {
	return nil
}
