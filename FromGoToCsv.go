package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	rows := [][]string{
		{"Country", "city", "state"},
		{"usa", "SanJose", "California"},
		{"usa", "dallas", "texas"},
	}

	csvfile, err := os.Create("data.csv")

	if err != nil {
		log.Fatalf("Failed to create file")
	}

	cswriter := csv.NewWriter(csvfile)

	for _, row := range rows {
		_ = cswriter.Write(row)
	}
	cswriter.Flush()
	csvfile.Close()

}
