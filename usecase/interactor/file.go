package interactor 

import (
	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
)

type iReadService interface{
	ReadAll() ([]model.Pokemon, error)
}

type iWriteService interface{
	Write(artist model.Artist) error
}

type iFileService interface {
	iReadService
	iWriteService
}