package service

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
	csverr "github.com/ErickRodriguezWize/academy-go-q42021/errors"
)

//ReadCSV: Reads a .csv file specific path and create an array  with the content of csv file.
func ReadCSV(path string, pkms *[]model.Pokemon) error {
	//Opening the csv file using the path.
	file, err := os.Open(path)
	if err != nil {
		return csverr.FileError
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
			return csverr.ColumnParseError
		}

		//Append to the structured Slice of Pokemons.
		*pkms = append(*pkms, model.Pokemon{
			ID:   ID,
			Name: record[1],
		})

	}

	return nil

}
