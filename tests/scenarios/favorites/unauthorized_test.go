package favorites

import (
	"net/http"
	"testing"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/managers/favorites"
	"api-tests-template/tests/fixtures"
)

func TestFavoritesUnauthorized(t *testing.T) {
	owner := fixtures.Login(t)
	ad := fixtures.CreateAdvertisement(t, owner.Token, fixtures.AdvertisementRequest(t))
	for _, tc := range fixtures.UnauthorizedCases() {
		t.Run(tc.Name, func(t *testing.T) {
			for _, method := range []string{http.MethodGet, http.MethodPost, http.MethodDelete} {
				t.Run(method, func(t *testing.T) {
					t.Log("Запрашиваем избранное:", tc.Name, method)
					response := favorites.UnauthorizedError(t, method, tc.Token, ad.ID)
					assertions.APIError(t, response, "unauthorized", tc.Message)
					t.Log("Получили 401 с нужной ошибкой")
				})
			}
		})
	}
}
