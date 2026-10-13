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
	for _, k := range ongefmt_waarden[16890:16902] {
		fmt.Println(k)
	}
	kwartierwaarden, err := parseKwartieren(ongefmt_waarden)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Aantal kwartierwaarden:%d\n", len(kwartierwaarden)/2)
	for _, k := range kwartierwaarden[16906:16918] {
		fmt.Println(k)
	}
	afnametotaal := totaal(kwartierwaarden, "Afname")
	injectietotaal := totaal(kwartierwaarden, "Injectie")

	fmt.Printf("Afname: %.3f kWh\n", afnametotaal)
	fmt.Printf("Injectie: %.3f kWh\n", injectietotaal)

}
