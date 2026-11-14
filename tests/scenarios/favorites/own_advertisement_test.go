package favorites

import (
	"net/http"
	"testing"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/managers/favorites"
	"api-tests-template/tests/fixtures"
)

func TestFavoriteOwnAdvertisement(t *testing.T) {
	owner := fixtures.Login(t)
	ad := fixtures.CreateAdvertisement(t, owner.Token, fixtures.AdvertisementRequest(t))
	t.Log("Пробуем добавить своё объявление в избранное")
	response := favorites.AddFavoriteError(t, owner.Token, ad.ID, http.StatusForbidden)
	assertions.APIError(t, response, "forbidden", "Cannot add own advertisement to favorites")
	assertions.FavoritePresence(t, favorites.GetFavorites(t, owner.Token, ad.Title), ad.ID, false)
	t.Log("Получили 403, своё объявление не добавилось")
}
