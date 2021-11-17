package service

import (
	"encoding/csv"
	"io"
	"os"
	"errors"
	"strconv"

	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
)

var (
	CSVError = errors.New("Couldn't open CSV File")
	ColumnParseError = errors.New("First Column of CSV most have an Integer Value")
)


//ReadCSV reads a .csv file specific path 
func ReadCSV(path string, pkms *[]model.Pokemon) (error){
	file, err := os.Open(path)
	
	if err != nil {
		return  CSVError //errors.New("Couldn't Open CSV File")
	}

	r:= csv.NewReader(file)

	//For loop that read all records from CSV File.
	for {
		//Lee el primer (row) del csv. 
		record, errCsv := r.Read()

		//io.EOF error trigger by the end of the file.  
		if errCsv != nil { 
			if errCsv== io.EOF{
				break
			}
			
			return errCsv
		}

		// Parse validation of ID value (Integer). 
		ID, errS := strconv.Atoi(record[0])
		if errS!=nil{
			return ColumnParseError
		}

		//Append to the structured Slice of Pokemons. 
		*pkms = append(*pkms, model.Pokemon{
			ID: ID,
			Name: record[1],
		})

	}

	return nil

}