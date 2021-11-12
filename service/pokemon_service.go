package service

import (
	"WizelineApi/domain/model"
	"errors"
)

func GetPokemonById(pokemons []model.Pokemon, id int) (model.Pokemon, error){
	t_pkm := model.Pokemon{}

	for _,pk := range pokemons{
		if pk.ID == id {
			t_pkm = pk
			break
		}
	}

	if t_pkm.Name == "" {
		return t_pkm, errors.New("Couldnt find pokemon")
	}else{
		return t_pkm, nil
	}

}