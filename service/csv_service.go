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

//Funcion para leer un CSV. 
func ReadCSV(path string, pkms *[]model.Pokemon) (error){
	
	file, err := os.Open(path)
	
	if err != nil {
		return  errors.New("Couldn't Open  File")
	}

	r:= csv.NewReader(file)

	//Loop para leer todos los records del CSV
	for {
		//Lee el primer (row) del csv. 
		record, errCsv := r.Read()

		//io.EOF es el error generado cuando se llega al final del archivo
		if errCsv == io.EOF{ //cuando llegue al final del archivo, se realiza el break del loop for.
			break
		}

		if errCsv != nil {
			return errCsv 
		}

		//Validacion para saber que el CSV contenga ID entero en la primera columna 
		id, errS := strconv.Atoi(record[0])
		if errS!=nil{
			return errors.New("First column of CSV most have an Integer Value")
		}

		//Append al slice de structura Pokemons, utilizado en los endpoints de PokemonController
		*pkms = append(*pkms, model.Pokemon{
			ID: id,
			Name: record[1],
		})

	}

	return nil
}
