package advertisementCreation

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/assertions"
	advertisement "api-tests-template/internal/managers/advertisement"
	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/tests/fixtures"
)

func TestCreateAdvertisement(t *testing.T) {
	owner := fixtures.Login(t)
	t.Log("Создаём объявление со всеми полями и тремя фото")
	request := fixtures.AdvertisementRequest(t)
	request.Fields["description"] = strings.Repeat("я", 500)
	request.Fields["price"] = "1000000"
	request.Fields["quantity"] = "100"
	request.Photos = append(request.Photos, request.Photos[0], request.Photos[0])
	ad := fixtures.CreateAdvertisement(t, owner.Token, request)
	require.NotEmpty(t, ad.ID, "Создание не вернуло ID объявления")

	t.Log("Получаем объявление по ID и проверяем поля")
	actual := advertisement.GetAdvertisement(t, ad.ID)
	require.Equal(t, request.Fields["title"], actual.Title, "Не совпадает название")
	require.Equal(t, request.Fields["description"], actual.Description, "Не совпадает описание")
	require.Equal(t, float64(1000000), actual.Price, "Не совпадает цена")
	require.Equal(t, 100, actual.Quantity, "Не совпадает количество")
	require.Equal(t, owner.User.Id, actual.UserID, "Не совпадает владелец")
	require.Equal(t, ad.ID, actual.ID, "Не совпадает ID")
	require.Equal(t, ad.Photos, actual.Photos, "Фотографии отличаются от ответа создания")

	t.Log("Ищем объявление по полному названию")
	list := advertisements.GetAdvertisementsList(t, owner.Token, actual.Title)
	assertions.ContainsAdvertisementByTitle(t, list.Items, actual.Title)
	assertions.ContainsAdvertisementByID(t, list.Items, ad.ID)

	t.Log("Получаем список фото и скачиваем каждое")
	photos := advertisement.GetPhotos(t, owner.Token, ad.ID)
	require.Len(t, photos, len(request.Photos), "Сохранены не все фотографии")
	require.Equal(t, actual.Photos, photos, "Список фотографий отличается от карточки")
	for i, photo := range photos {
		require.NotEmpty(t, photo.ID, "Пустой ID фотографии")
		require.NotEmpty(t, photo.URL, "Пустой URL фотографии")
		require.Equal(t, i, photo.SortOrder, "Неверный порядок фотографий")
		expected, err := os.ReadFile(request.Photos[i])
		require.NoError(t, err)
		require.Equal(t, expected, advertisement.DownloadPhoto(t, photo.URL), "Загруженная фотография отличается от исходной")
	}
	t.Log("Поля объявления и фотографии совпали, поиск нашёл объявление")
}
