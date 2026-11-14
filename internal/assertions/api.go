package assertions

import (
	"testing"

	"github.com/stretchr/testify/require"

	advModels "api-tests-template/internal/managers/advertisement/models"
	"api-tests-template/internal/managers/advertisements/models"
	favModels "api-tests-template/internal/managers/favorites/models"
)

// APIError сравнивает код и сообщение ошибки API с ожидаемыми.
func APIError(t *testing.T, actual advModels.ErrorResponse, code, message string) {
	t.Helper()
	require.Equal(t, code, actual.Error, "Неверный код ошибки API")
	require.Equal(t, message, actual.Message, "Неверное сообщение ошибки API")
}

// ContainsAdvertisementByID проверяет наличие конкретного объявления в поиске.
func ContainsAdvertisementByID(t *testing.T, items []models.Advertisement, id string) {
	t.Helper()
	ids := make([]string, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	require.Contains(t, ids, id, "В поиске нет созданного объявления")
}

// FavoritePresence проверяет наличие или отсутствие объявления в избранном.
func FavoritePresence(t *testing.T, list favModels.FavoritesList, id string, expected bool) {
	t.Helper()
	found := false
	for _, item := range list.Items {
		if item.Advertisement.ID == id {
			found = true
			break
		}
	}
	require.Equal(t, expected, found, "Неожиданное наличие объявления %s в избранном", id)
}
