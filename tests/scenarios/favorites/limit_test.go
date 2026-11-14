package favorites

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/managers/favorites"
	"api-tests-template/tests/fixtures"
)

const maxFavorites = 50

func TestFavoriteLimit(t *testing.T) {
	owner, reader := fixtures.LoginPair(t)
	initial := favorites.GetFavorites(t, reader.Token, "")
	require.LessOrEqual(t, initial.Total, maxFavorites, "Изначальный размер избранного превышает лимит")
	t.Log("Заполняем избранное до лимита. Уже есть записей:", initial.Total)
	for i := initial.Total; i < maxFavorites; i++ {
		ad := fixtures.CreateAdvertisement(t, owner.Token, fixtures.AdvertisementRequest(t))
		fixtures.AddFavorite(t, reader.Token, ad.ID)
	}
	require.Equal(t, maxFavorites, favorites.GetFavorites(t, reader.Token, "").Total)
	ad := fixtures.CreateAdvertisement(t, owner.Token, fixtures.AdvertisementRequest(t))
	t.Log("Пробуем добавить 51-е объявление")
	response := favorites.AddFavoriteError(t, reader.Token, ad.ID, http.StatusBadRequest)
	assertions.APIError(t, response, "limit_reached", "Favorites limit of 50 reached")
	assertions.FavoritePresence(t, favorites.GetFavorites(t, reader.Token, ad.Title), ad.ID, false)
	t.Log("Получили ошибку лимита, объявление не добавилось")
}
