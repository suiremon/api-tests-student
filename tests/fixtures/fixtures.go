package fixtures

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/advertisement"
	"api-tests-template/internal/managers/advertisement/models"
	advModels "api-tests-template/internal/managers/advertisements/models"
	"api-tests-template/internal/managers/auth"
	authModels "api-tests-template/internal/managers/auth/models"
	"api-tests-template/internal/managers/favorites"
	"api-tests-template/internal/utils"
)

type Advertisement struct {
	advModels.Advertisement
	token   string
	deleted bool
}

type UnauthorizedCase struct {
	Name    string
	Token   string
	Message string
}

// Login загружает окружение и авторизует основного тестового пользователя.
func Login(t *testing.T) authModels.LoginOkResponse {
	t.Helper()
	utils.LoadEnv()
	t.Log("Входим в аккаунт владельца")
	return auth.Login(t, os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))
}

// LoginPair авторизует владельца объявления и другого пользователя избранного.
func LoginPair(t *testing.T) (authModels.LoginOkResponse, authModels.LoginOkResponse) {
	t.Helper()
	owner := Login(t)
	t.Log("Входим во второй аккаунт")
	reader := auth.Login(t, os.Getenv("TEST_SECOND_LOGIN"), os.Getenv("TEST_SECOND_PASSWORD"))
	require.NotEqual(t, owner.User.Id, reader.User.Id, "Для избранного нужны два разных пользователя")
	return owner, reader
}

// AdvertisementRequest готовит валидный multipart-запрос с уникальным названием и фото.
func AdvertisementRequest(t *testing.T) models.CreateRequest {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	require.True(t, ok, "Не удалось определить путь к тестовой фотографии")
	return models.CreateRequest{
		Fields: map[string]string{
			"title":       "API тест " + utils.RandomString(20),
			"description": "Объявление для автоматической проверки API",
			"price":       "123.45",
			"quantity":    "3",
		},
		Photos: []string{filepath.Join(filepath.Dir(file), "..", "testdata", "photo.png")},
	}
}

// CreateAdvertisement создаёт объявление и регистрирует удаление даже при падении теста.
func CreateAdvertisement(t *testing.T, token string, request models.CreateRequest) *Advertisement {
	t.Helper()
	t.Log("Создаём тестовое объявление", request.Fields["title"])
	ad := &Advertisement{Advertisement: advertisement.CreateAdvertisement(t, token, request), token: token}
	require.NotEmpty(t, ad.ID, "Создание объявления не вернуло ID")
	t.Cleanup(func() {
		if !ad.deleted {
			ad.Delete(t)
		}
	})
	return ad
}

// Delete удаляет тестовое объявление и предотвращает повторное удаление при очистке.
func (ad *Advertisement) Delete(t *testing.T) {
	t.Helper()
	t.Log("Удаляем тестовое объявление", ad.ID)
	advertisement.DeleteAdvertisement(t, ad.token, ad.ID)
	ad.deleted = true
}

// AddFavorite добавляет запись и регистрирует очистку, если её ещё не удалил сценарий.
func AddFavorite(t *testing.T, token, id string) {
	t.Helper()
	item := favorites.AddFavorite(t, token, id)
	require.Equal(t, id, item.Advertisement.ID, "В избранное добавлено другое объявление")
	t.Cleanup(func() {
		for _, entry := range favorites.GetFavorites(t, token, "").Items {
			if entry.Advertisement.ID == id {
				t.Log("Убираем тестовое объявление из избранного", id)
				favorites.RemoveFavorite(t, token, id)
				return
			}
		}
	})
}

// UnauthorizedCases возвращает отсутствующий и заведомо неправильный токены.
func UnauthorizedCases() []UnauthorizedCase {
	return []UnauthorizedCase{
		{Name: "без токена", Message: "Authorization header required"},
		{Name: "неверный токен", Token: "incorrect_token", Message: "Invalid or expired token"},
	}
}
