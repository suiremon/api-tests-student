package favorites

import (
	"net/http"
	"testing"

	client "api-tests-template/internal/client/http/favorites"
	advModels "api-tests-template/internal/managers/advertisement/models"
	"api-tests-template/internal/managers/favorites/models"
	"api-tests-template/internal/managers/response"
)

// AddFavorite добавляет объявление и возвращает запись избранного.
func AddFavorite(t *testing.T, token, id string) models.Favorite {
	t.Helper()
	return response.JSON[models.Favorite](t, client.HttpPostFavorite(t, token, id, http.StatusCreated))
}

// GetFavorites возвращает список избранного текущего пользователя.
func GetFavorites(t *testing.T, token, search string) models.FavoritesList {
	t.Helper()
	return response.JSON[models.FavoritesList](t, client.HttpGetFavorites(t, token, search, http.StatusOK))
}

// RemoveFavorite удаляет запись избранного и закрывает пустой ответ.
func RemoveFavorite(t *testing.T, token, id string) {
	t.Helper()
	response.Close(t, client.HttpDeleteFavorite(t, token, id, http.StatusNoContent))
}

// AddFavoriteError возвращает ошибку добавления в избранное.
func AddFavoriteError(t *testing.T, token, id string, status int) advModels.ErrorResponse {
	t.Helper()
	return response.JSON[advModels.ErrorResponse](t, client.HttpPostFavorite(t, token, id, status))
}

// RemoveFavoriteError возвращает ошибку удаления из избранного.
func RemoveFavoriteError(t *testing.T, token, id string, status int) advModels.ErrorResponse {
	t.Helper()
	return response.JSON[advModels.ErrorResponse](t, client.HttpDeleteFavorite(t, token, id, status))
}

// UnauthorizedError выполняет операцию избранного с ожидаемым отказом авторизации.
func UnauthorizedError(t *testing.T, method, token, id string) advModels.ErrorResponse {
	t.Helper()
	switch method {
	case http.MethodGet:
		return response.JSON[advModels.ErrorResponse](t, client.HttpGetFavorites(t, token, "", http.StatusUnauthorized))
	case http.MethodPost:
		return AddFavoriteError(t, token, id, http.StatusUnauthorized)
	case http.MethodDelete:
		return RemoveFavoriteError(t, token, id, http.StatusUnauthorized)
	default:
		t.Fatalf("Неизвестная операция избранного: %s", method)
		return advModels.ErrorResponse{}
	}
}
