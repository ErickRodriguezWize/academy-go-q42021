package service

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
	csverr "github.com/ErickRodriguezWize/academy-go-q42021/errors"
)

type CsvService struct {
	config model.Config
}

// NewCsvService: Constructor for CsvService struct. 
func NewCsvService(config model.Config) CsvService {
	return CsvService{config}
}

// ReadCSV: Reads a .csv file specific path and create an array  with the content of csv file.
func (cs CsvService) ReadCSV(pkms *[]model.Pokemon) error {
	// Opening the csv file using the path.
	path := cs.config.PokemonCsvPath
	file, err := os.Open(path)
	if err != nil {
		return csverr.ErrFileError
	}

	// Read the content of the file using the package csv.
	r := csv.NewReader(file)

	// For loop that read all records from CSV File.
	for {
		//Read the row of the CSV content.
		record, errCsv := r.Read()

		// io.EOF error trigger by the end of the file.
		if errCsv != nil {
			if errors.Is(errCsv, io.EOF) {
				break
			}

			return csverr.ErrEndOfFile
		}

		// Parse validation of ID value (Integer).
		id, err := strconv.Atoi(record[0])

		if err != nil {
			return csverr.ErrColumnParseError
		}

		// Append to the structured Slice of Pokemons.
		*pkms = append(*pkms, model.Pokemon{
			ID:   id,
			Name: record[1],
		})

	}

	return nil

}

// WriteArtistIntoCSV: Append data(model.Artist) to the file '/test/bateria_csv/artist.csv'.
// Return error.
func (cs CsvService) StoreArtist(artist model.Artist) error {
	path := cs.config.ArtistCsvPath

	// Opening/Creating csv file.
	// O_WRONLY: Open File in Write only mode.
	// O_CREATE: Create new file if none exists.
	// O_APPEND: Append data to the file when writing.
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println("err", err.Error())

		return csverr.ErrFileError
	}
	defer file.Close()

	// Create a Writer, to append in the csv file.
	w := csv.NewWriter(file)

	// Create a slice of string (ask by Write Method).
	data := []string{artist.ID, artist.Name, artist.SpotifyURL}
	if err := w.Write(data); err != nil {
		return csverr.ErrBadWrite
	}

	defer w.Flush()

	return nil
}