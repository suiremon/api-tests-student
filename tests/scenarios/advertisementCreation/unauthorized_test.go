package advertisementCreation

import (
	"net/http"
	"testing"

	"api-tests-template/internal/assertions"
	advertisement "api-tests-template/internal/managers/advertisement"
	"api-tests-template/tests/fixtures"
)

func TestCreateAdvertisementUnauthorized(t *testing.T) {
	for _, tc := range fixtures.UnauthorizedCases() {
		t.Run(tc.Name, func(t *testing.T) {
			t.Log("Пробуем создать объявление:", tc.Name)
			request := fixtures.AdvertisementRequest(t)
			response := advertisement.CreateAdvertisementError(t, tc.Token, request, http.StatusUnauthorized)
			assertions.APIError(t, response, "unauthorized", tc.Message)
			t.Log("Получили 401 с нужной ошибкой")
		})
	}
}
