package response

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

// JSON разбирает тело HTTP-ответа в указанную модель и закрывает его.
func JSON[T any](t *testing.T, raw *http.Response) T {
	t.Helper()
	defer Close(t, raw)
	body, err := io.ReadAll(raw.Body)
	require.NoError(t, err, "Не удалось прочитать тело ответа")
	var result T
	require.NoError(t, json.Unmarshal(body, &result), "Не удалось разобрать JSON: %s", body)
	return result
}

// Close освобождает тело HTTP-ответа и проверяет ошибку закрытия.
func Close(t *testing.T, raw *http.Response) {
	t.Helper()
	require.NoError(t, raw.Body.Close(), "Не удалось закрыть тело ответа")
}
