package service

import (
	"encoding/csv"
	"io"
	"log"
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

// WriteArtistIntoCSV: Append data(model.Artist) to the file '/test/bateria_csv/artist.csv'.
// Return error.
func WriteArtistIntoCSV(path string, artist model.Artist) error {
	//Opening/Creating csv file.
	// O_WRONLY: Open File in Write only mode.
	// O_CREATE: Create new file if none exists.
	// O_APPEND: Append data to the file when writing.
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println("err", err.Error())
		return csverr.FileError
	}
	defer file.Close()

	//Create a Writer, to append in the csv file.
	w := csv.NewWriter(file)

	// Create a slice of string (ask by Write Method).
	data := []string{artist.ID, artist.Name, artist.SpotifyURL}
	if err := w.Write(data); err != nil {
		return csverr.BadWrite
	}

	defer w.Flush()

	return nil
}
