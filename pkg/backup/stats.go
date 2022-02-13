package backup

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/open-blockchain-explorer/tnbassist/model"
)

// AppendToCSV append stats to a CSV file
func AppendToCSV(filename string, outdir string, stats *model.LegacyStats) error {
	f, err := os.OpenFile(fmt.Sprintf("%s/%s.csv", outdir, filename), os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	fstats, err := f.Stat()
	if err != nil {
		return err
	}

	csvWriter := csv.NewWriter(f)
	if fstats.Size() == 0 {
		// Write headers
		err := csvWriter.Write(stats.CSVHeaders())
		if err != nil {
			return err
		}
	}
	err = csvWriter.Write(stats.CSVRow())
	if err != nil {
		return err
	}
	csvWriter.Flush()

	return nil
}
