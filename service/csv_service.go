package service

import (
	"encoding/csv"
	"io"
	"os"
	"WizelineApi/domain/model"
	//"fmt"
	"errors"
	"strconv"
)

func ReadCSV(path string, pkms *[]model.Pokemon) (error){
	
	file, err := os.Open(path)
	

	if err != nil {
		return  errors.New("Couldn't Open  File")
	}

	r:= csv.NewReader(file)

	for {
		record, errCsv := r.Read()

		//io.EOF es el error generado cuando se llega al final del archivo
		if errCsv == io.EOF{ //cuando llegue al final del archivo, se realiza el break del loop for.
			break
		}

		if errCsv != nil {
			return errCsv 
		}

		id, errS := strconv.Atoi(record[0])
		if errS!=nil{
			return errors.New("First column of CSV most have an Integer Value")
		}

		*pkms = append(*pkms, model.Pokemon{
			ID: id,
			Name: record[1],
		})

	}

	return nil
}
