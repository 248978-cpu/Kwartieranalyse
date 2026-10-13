package main

import (
	"fmt"
	"os"
)

func main() {
	ongefmt_waarden, err := import_csv(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	kwartierwaarden, err := parseKwartieren(ongefmt_waarden)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, k := range kwartierwaarden[:4] {
		fmt.Println(k)
	}
	afnametotaal := totaal(kwartierwaarden, "Afname")
	injectietotaal := totaal(kwartierwaarden, "Injectie")

	fmt.Printf("Afname: %.3f kWh\n", afnametotaal)
	fmt.Printf("Injectie: %.3f kWh\n", injectietotaal)

}
