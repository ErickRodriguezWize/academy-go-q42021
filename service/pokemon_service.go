package service

import (
	"errors"

	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
)

//Servicio utilizado para buscar por ID dentro de un slice del modelod e Pokemon
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