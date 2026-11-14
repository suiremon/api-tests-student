package favorites

import (
	"net/http"
	"testing"

	"api-tests-template/internal/assertions"
	advModels "api-tests-template/internal/managers/advertisement/models"
	"api-tests-template/internal/managers/favorites"
	"api-tests-template/tests/fixtures"
)

const missingAdID = "00000000-0000-4000-8000-000000000000"

func TestFavoritesInvalidAdvertisement(t *testing.T) {
	user := fixtures.Login(t)
	cases := []struct {
		name, id, addMessage, removeMessage string
		status                              int
	}{
		{"несуществующий ID", missingAdID, "Advertisement not found", "Favorite entry not found", http.StatusNotFound},
		{"некорректный ID", "invalid-id", "invalid UUID length: 10", "Invalid advertisement ID", http.StatusBadRequest},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			for _, operation := range []string{"добавление", "удаление"} {
				t.Run(operation, func(t *testing.T) {
					t.Log("Проверяем:", tc.name, "операция:", operation)
					var response advModels.ErrorResponse
					message := tc.addMessage
					if operation == "добавление" {
						response = favorites.AddFavoriteError(t, user.Token, tc.id, tc.status)
					} else {
						response = favorites.RemoveFavoriteError(t, user.Token, tc.id, tc.status)
						message = tc.removeMessage
					}
					code := "invalid_request"
					if tc.status == http.StatusNotFound {
						code = "not_found"
					}
					assertions.APIError(t, response, code, message)
					t.Log("Статус и сообщение ошибки совпали")
				})
			}
		})
	}
	t.Run("пустой ID при добавлении", func(t *testing.T) {
		t.Log("Добавляем в избранное с пустым advertisement_id")
		response := favorites.AddFavoriteError(t, user.Token, "", http.StatusBadRequest)
		assertions.APIError(t, response, "invalid_request", "invalid UUID length: 0")
		t.Log("Получили ошибку пустого ID")
	})
}
