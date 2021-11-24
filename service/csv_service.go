package service

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"

	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
)

//ReadCSV: Reads a .csv file specific path
func ReadCSV(path string, pkms *[]model.Pokemon) error {
	//Potential Errors of CSV.
	CSVError := errors.New("Couldn't open CSV File")
	ColumnParseError := errors.New("First Column of CSV most have an Integer Value")

	//Opening the csv file using the path.
	file, err := os.Open(path)
	if err != nil {
		return CSVError
	}

	//Read the content of the file using the package csv.
	r := csv.NewReader(file)

	//For loop that read all records from CSV File.
	for {
		//Read the row of the CSV content.
		record, errCsv := r.Read()

		//io.EOF error trigger by the end of the file.
		if errCsv != nil {
			if errCsv == io.EOF {
				break
			}

			return errCsv
		}

		// Parse validation of ID value (Integer).
		ID, errS := strconv.Atoi(record[0])
		if errS != nil {
			return ColumnParseError
		}

		//Append to the structured Slice of Pokemons.
		*pkms = append(*pkms, model.Pokemon{
			ID:   ID,
			Name: record[1],
		})

	}

	return nil

}
