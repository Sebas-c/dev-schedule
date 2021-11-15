package main

import (
	"os"

	table "github.com/jedib0t/go-pretty/v6/table"
)

func main() {
	primaries := [6]string{"Nicky", "Daniel", "Glenn", "Emanuel", "Martin", "Sebastien"}

	secondaries := primaries[1:]
	secondaries = append(secondaries, primaries[0])

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Week #", "Primary", "Secondary"})

	weekCounter := 46
	for i := 0; i < 10; i++ {
		for j := 0; j < len(primaries); j++ {
			t.AppendRows([]table.Row{
				{weekCounter, primaries[j], secondaries[j]},
			})
			weekCounter++
			if weekCounter > 52 {
				weekCounter = 1
			}
			t.AppendSeparator()
		}
		// t.AppendSeparator()

		// Swapping secondaries around
		tempSecondary := secondaries[1:]
		tempSecondary = append(tempSecondary, secondaries[0])

		if tempSecondary[0] == primaries[0] {
			secondaries = tempSecondary
			tempSecondary = secondaries[1:]
			tempSecondary = append(tempSecondary, secondaries[0])
		}

		secondaries = tempSecondary
	}
	t.Render()

}
