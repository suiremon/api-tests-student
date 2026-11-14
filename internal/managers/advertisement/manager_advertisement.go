package advertisement

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	client "api-tests-template/internal/client/http/advertisement"
	"api-tests-template/internal/managers/advertisement/models"
	advModels "api-tests-template/internal/managers/advertisements/models"
	"api-tests-template/internal/managers/response"
)

// CreateAdvertisement создаёт объявление и возвращает его параметры.
func CreateAdvertisement(t *testing.T, token string, request models.CreateRequest) advModels.Advertisement {
	t.Helper()
	return response.JSON[advModels.Advertisement](t, client.HttpPostAdvertisement(t, token, request, http.StatusCreated))
}

// CreateAdvertisementError возвращает ошибку неуспешного создания объявления.
func CreateAdvertisementError(t *testing.T, token string, request models.CreateRequest, status int) models.ErrorResponse {
	t.Helper()
	raw := client.HttpPostAdvertisement(t, token, request, status)
	if raw.StatusCode == http.StatusCreated {
		ad := response.JSON[advModels.Advertisement](t, raw)
		if ad.ID != "" {
			t.Cleanup(func() {
				DeleteAdvertisement(t, token, ad.ID)
			})
		}
		t.Fatalf("Невалидный запрос создал объявление %s вместо ошибки %d", ad.ID, status)
	}
	return response.JSON[models.ErrorResponse](t, raw)
}

// GetAdvertisement возвращает карточку объявления по ID.
func GetAdvertisement(t *testing.T, id string) advModels.Advertisement {
	t.Helper()
	return response.JSON[advModels.Advertisement](t, client.HttpGetAdvertisement(t, id))
}

// DeleteAdvertisement удаляет объявление и закрывает пустой ответ.
func DeleteAdvertisement(t *testing.T, token, id string) {
	t.Helper()
	response.Close(t, client.HttpDeleteAdvertisement(t, token, id))
}

// GetPhotos возвращает фотографии объявления.
func GetPhotos(t *testing.T, token, id string) []advModels.Photo {
	t.Helper()
	return response.JSON[[]advModels.Photo](t, client.HttpGetPhotos(t, token, id))
}

// DownloadPhoto возвращает содержимое загруженной фотографии.
func DownloadPhoto(t *testing.T, photoURL string) []byte {
	t.Helper()
	raw := client.HttpGetPhoto(t, photoURL)
	defer response.Close(t, raw)
	body, err := io.ReadAll(raw.Body)
	require.NoError(t, err, "Не удалось прочитать фотографию")
	return body
}
