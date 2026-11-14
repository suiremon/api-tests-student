package favorites

import (
	"net/http"
	"testing"

	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
)

// HttpPostFavorite добавляет объявление в избранное с проверкой статуса.
func HttpPostFavorite(t *testing.T, token, id string, status int) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Post(path.FavoritesPath).
		JSON(map[string]string{"advertisement_id": id}).
		Expect(t).Status(status).End().Response
}

// HttpGetFavorites получает до 50 записей избранного с фильтром по названию.
func HttpGetFavorites(t *testing.T, token, search string, status int) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Get(path.FavoritesPath).
		Query("page_size", "50").Query("search", search).
		Expect(t).Status(status).End().Response
}

// HttpDeleteFavorite удаляет объявление из избранного с проверкой статуса.
func HttpDeleteFavorite(t *testing.T, token, id string, status int) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Delete(path.FavoritesPath + "/" + id).
		Expect(t).Status(status).End().Response
}
