package advertisementCreation

import (
	"net/http"
	"strings"
	"testing"

	"api-tests-template/internal/assertions"
	advertisement "api-tests-template/internal/managers/advertisement"
	"api-tests-template/tests/fixtures"
)

func TestCreateAdvertisementInvalidParameters(t *testing.T) {
	owner := fixtures.Login(t)
	cases := []struct {
		name, field, value, message string
		omit                        bool
		photoCount                  int
	}{
		{"нет title", "title", "", "title is required", true, 1},
		{"пустой title", "title", "", "title is required", false, 1},
		{"длинный title", "title", strings.Repeat("a", 51), "title must not exceed 50 characters provided 51", false, 1},
		{"нет description", "description", "", "description is required", true, 1},
		{"пустой description", "description", "", "description is required", false, 1},
		{"длинный description", "description", strings.Repeat("a", 501), "description must be between 1 and 500 characters", false, 1},
		{"отрицательная цена", "price", "-1", "price must be between 0 and 1000000 inclusive", false, 1},
		{"слишком большая цена", "price", "1000001", "price must be between 0 and 1000000 inclusive", false, 1},
		{"цена не число", "price", "oops", "cannot convert string value oops to float", false, 1},
		{"нулевое количество", "quantity", "0", "quantity must be between 1 and 100 inclusive", false, 1},
		{"отрицательное количество", "quantity", "-1", "quantity must be between 1 and 100 inclusive", false, 1},
		{"слишком большое количество", "quantity", "101", "quantity must be between 1 and 100 inclusive", false, 1},
		{"количество не число", "quantity", "oops", "cannot convert string value to int", false, 1},
		{"нет фотографий", "", "", "at least 1 photo is required", false, 0},
		{"четыре фотографии", "", "", "maximum 3 photos allowed", false, 4},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log("Проверяем невалидный запрос:", tc.name)
			request := fixtures.AdvertisementRequest(t)
			if tc.field != "" {
				request.Fields[tc.field] = tc.value
				if tc.omit {
					delete(request.Fields, tc.field)
				}
			}
			photo := request.Photos[0]
			request.Photos = nil
			for range tc.photoCount {
				request.Photos = append(request.Photos, photo)
			}
			t.Log("Отправляем запрос создания")
			response := advertisement.CreateAdvertisementError(t, owner.Token, request, http.StatusBadRequest)
			assertions.APIError(t, response, "invalid_request", tc.message)
			t.Log("Получили 400 с нужной ошибкой")
		})
	}
}
