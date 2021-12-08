package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ErickRodriguezWize/academy-go-q42021/domain/model"
	spotierr "github.com/ErickRodriguezWize/academy-go-q42021/errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockArtistInteractor struct{
	mock.Mock
}

func(ma mockArtistInteractor) SearchArtist(name string) (model.Artist, error){
	arg:= ma.Called()
	return arg.Get(0).(model.Artist), arg.Error(1)
}

func(ma mockArtistInteractor) StoreArtist(artist model.Artist) error{
	arg:= ma.Called()
	return arg.Error(0)
}

// TestArtistController_SearchArtist: unit test.
func TestArtistController_SearchArtist(t *testing.T){
	//testcases
	testcases := []struct {
		name string 
		method    string
		endpoint  string
		response model.Artist
		error error
		want_code int
	}{
		{
			name:"SearchArtist: Linkin Park",
			method: "GET",
			endpoint: "/artists/linkin+park",
			response: model.Artist{ID:"6XyY86QOPPrYVGvF9ch6wz", Name: "Linkin Park", SpotifyURL:"https://open.spotify.com/artist/6XyY86QOPPrYVGvF9ch6wz"},
			error: nil,
			want_code: 200,
		},
		{
			name:"Search: Queens",
			method: "GET",
			endpoint: "/artists/queens",
			response: model.Artist{ID:"1Sp1JX3Degylt1q79Cx0iX", Name: "Queens", SpotifyURL:"https://open.spotify.com/artist/1Sp1JX3Degylt1q79Cx0iX"},
			error: nil,
			want_code: 200,
		},
		{
			name:"Not found Artist",
			method: "GET",
			endpoint: "/artists/queenssichu",
			response: model.Artist{},
			error: spotierr.ErrArtistNotFound,
			want_code: 400,
		},
		{
			name:"Not found with coma",
			method: "GET",
			endpoint: "/artists/papa,roach",
			response: model.Artist{},
			error: spotierr.ErrArtistNotFound,
			want_code: 400,
		},
	}
	
	for _,tsc := range testcases{
		// Start Test.
		t.Run(tsc.name, func(t *testing.T) {
			// Init Mocks
			mockInteractor:= mockArtistInteractor{}
			mockInteractor.On("SearchArtist").Return(tsc.response,tsc.error)

			// Dummy res http.ResponseWriter and req *http.Request
			req, _ := http.NewRequest(tsc.method, tsc.endpoint, nil)
			res := httptest.NewRecorder()

			// Implemented mocks
			service := NewArtistController(mockInteractor)

			// Executed methods
			service.SearchArtist(res,req)
			
			// Asserts
			assert.EqualValues(t, res.Code, tsc.want_code)
		})
	}

}