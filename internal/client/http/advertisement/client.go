package advertisement

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
	"api-tests-template/internal/managers/advertisement/models"
)

// HttpPostAdvertisement создаёт объявление из multipart-полей и фотографий.
func HttpPostAdvertisement(t *testing.T, token string, body models.CreateRequest, status int) *http.Response {
	t.Helper()
	request := apiRunner.GetRunner().Auth(token).Create().Post(path.AdvertisementPath)
	for name, value := range body.Fields {
		request.MultipartFormData(name, value)
	}
	if len(body.Photos) > 0 {
		request.MultipartFile("photos", body.Photos...)
	}
	return request.Expect(t).Status(status).End().Response
}

// HttpGetAdvertisement получает карточку объявления по ID.
func HttpGetAdvertisement(t *testing.T, id string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Create().Get(path.AdvertisementPath).
		Query("id", id).Expect(t).Status(http.StatusOK).End().Response
}

// HttpDeleteAdvertisement удаляет объявление его владельца.
func HttpDeleteAdvertisement(t *testing.T, token, id string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Delete(path.AdvertisementPath).
		Query("id", id).Expect(t).Status(http.StatusNoContent).End().Response
}

// HttpGetPhotos получает список фотографий объявления.
func HttpGetPhotos(t *testing.T, token, id string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().
		Get(fmt.Sprintf(path.AdvertisementPhotosPath, id)).Expect(t).Status(http.StatusOK).End().Response
}

// HttpGetPhoto скачивает фотографию с сервера хранения.
func HttpGetPhoto(t *testing.T, photoURL string) *http.Response {
	t.Helper()
	client := &http.Client{Timeout: 30 * time.Second}
	request, err := http.NewRequest(http.MethodGet, photoURL, nil)
	require.NoError(t, err)
	raw, err := client.Do(request)
	require.NoError(t, err, "Фотография недоступна: %s", photoURL)
	if raw.StatusCode != http.StatusOK {
		require.NoError(t, raw.Body.Close())
		t.Fatalf("Фотография %s вернула статус %d вместо 200", photoURL, raw.StatusCode)
	}
	return raw
}
