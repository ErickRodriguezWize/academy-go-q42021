package interactor

import (
	"testing"

	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
	spotierr "github.com/ErickRodriguezWize/academy-go-q42021/errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mock struct of ArtistService
type mockArtistService struct {
	mock.Mock
}

// SearchArtist: mock method of SearchArtist
func (ma mockArtistService) SearchArtist(name string) (model.Artist, error) {
	arg := ma.Called()
	return arg.Get(0).(model.Artist), arg.Error(1)
}

// Mock struct of WriteService
type mockWriteService struct {
	mock.Mock
}

// Write: mock method of Write.
func (mf mockWriteService) Write(artist model.Artist) error {
	arg := mf.Called()
	return arg.Error(0)
}

// TestArtistInteractor_SearchArtist :unit testing
func TestArtistInteractor_SearchArtist(t *testing.T) {
	//testcases
	testCases := []struct {
		name           string
		artist         string
		responseArtist model.Artist
		errorArtist    error
		errorWrite     error
		hasError       bool
	}{
		{
			name:           "Found Queens",
			artist:         "Queen",
			responseArtist: model.Artist{ID: "1Sp1JX3Degylt1q79Cx0iX", Name: "Queens", SpotifyURL: "https://open.spotify.com/artist/1Sp1JX3Degylt1q79Cx0iX"},
			errorArtist:    nil,
			errorWrite:     nil,
			hasError:       false,
		},
		{
			name:           "Found Linkin Park",
			artist:         "linkin+park",
			responseArtist: model.Artist{ID: "6XyY86QOPPrYVGvF9ch6wz", Name: "Linkin Park", SpotifyURL: "https://open.spotify.com/artist/6XyY86QOPPrYVGvF9ch6wz"},
			errorArtist:    nil,
			errorWrite:     nil,
			hasError:       false,
		},
		{
			name:           "Not Found Artist",
			artist:         "quenxious",
			responseArtist: model.Artist{},
			errorArtist:    spotierr.ErrArtistNotFound,
			errorWrite:     nil,
			hasError:       true,
		},
		{
			name:           "Coma in name",
			artist:         "papa,roach",
			responseArtist: model.Artist{},
			errorArtist:    spotierr.ErrArtistNotFound,
			errorWrite:     nil,
			hasError:       true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Init mocks.
			mockArtist := mockArtistService{}
			mockArtist.On("SearchArtist").Return(tc.responseArtist, tc.errorArtist)
			mockWrite := mockWriteService{}
			mockWrite.On("Write").Return(tc.errorWrite)

			// Implemented Mocks.
			service := NewArtistInteractor(mockWrite, mockArtist)

			// Execute method.
			result, err := service.SearchArtist(tc.artist)

			//assert results.
			assert.EqualValues(t, result, tc.responseArtist)
			if tc.hasError {
				assert.EqualError(t, err, tc.errorArtist.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}

}