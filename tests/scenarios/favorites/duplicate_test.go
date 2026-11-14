package favorites

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/managers/favorites"
	"api-tests-template/tests/fixtures"
)

func TestFavoriteDuplicate(t *testing.T) {
	owner, reader := fixtures.LoginPair(t)
	ad := fixtures.CreateAdvertisement(t, owner.Token, fixtures.AdvertisementRequest(t))
	fixtures.AddFavorite(t, reader.Token, ad.ID)
	t.Log("Повторно добавляем то же объявление")
	response := favorites.AddFavoriteError(t, reader.Token, ad.ID, http.StatusConflict)
	assertions.APIError(t, response, "already_exists", "Advertisement is already in favorites")
	list := favorites.GetFavorites(t, reader.Token, ad.Title)
	assertions.FavoritePresence(t, list, ad.ID, true)
	require.Equal(t, 1, list.Total, "Повторное добавление создало дубликат")
	t.Log("Получили 409, дубликат не появился")
}
