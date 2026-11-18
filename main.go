package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Erik-Jerman-Photography/vremenske-postaje/vreme"
	"github.com/urfave/cli/v3"
)

func main() {

	cmd := &cli.Command{
		Name:  "vreme",
		Usage: "beleženje meritev vremenskih postaj",
		Flags: []cli.Flag{
			&cli.FloatFlag{
				Name:  "minTemp",
				Value: -50,
				Usage: "najnižja dovoljena temperatura meritve (°C)",
			},
			&cli.FloatFlag{
				Name:  "maxTemp",
				Value: 50,
				Usage: "najvišja dovoljena temperatura meritve (°C)",
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Value: 3,
				Usage: "najmanjše število meritev za izračun statistike",
			},
		},
		Action: zazeni,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}

func zazeni(ctx context.Context, cmd *cli.Command) error {
	minTemp := cmd.Float("minTemp")
	maxTemp := cmd.Float("maxTemp")
	stMeritev := cmd.Int("stMeritev")

	postaje := make(map[string]vreme.Postaja)

	vreme.DodajPostajo(postaje, "LJ01", "Tivoli", "Ljubljana")
	vreme.DodajPostajo(postaje, "LJ02", "Vič", "Ljubljana")
	vreme.DodajPostajo(postaje, "NM01", "Ragovo", "Novo mesto")
	vreme.DodajPostajo(postaje, "NM02", "Center", "Novo mesto")

	meritve := map[string][]float64{
		"LJ01": {23.5, 22.0, 21.0},
		"LJ02": {21.5, 22.1},
		"NM01": {23.0, 22.5, 20.0},
		"NM02": {5.0, 12.3, 10.8},
	}
	for oznaka, temperature := range meritve {
		for _, t := range temperature {
			vreme.DodajMeritev(postaje, oznaka, t, minTemp, maxTemp)
		}
	}

	fmt.Println("prvi:")
	vreme.IzpisMeritev(postaje)
	fmt.Println()

	vreme.DodajMeritev(postaje, "NM02", -1.0, minTemp, maxTemp)
	fmt.Println("drugi:")
	vreme.IzpisMeritev(postaje)
	fmt.Println()

	fmt.Println("porocilo:")
	vreme.IzpisPorocila(postaje, stMeritev)
	fmt.Println()

	vreme.DodajMeritev(postaje, "LJ02", 52.1, minTemp, maxTemp)
	vreme.DodajMeritev(postaje, "KP01", 32.8, minTemp, maxTemp)
	return nil
}
