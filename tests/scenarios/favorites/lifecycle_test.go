package favorites

import (
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/managers/favorites"
	"api-tests-template/tests/fixtures"
)

func TestFavoritesLifecycle(t *testing.T) {
	owner, reader := fixtures.LoginPair(t)
	ad := fixtures.CreateAdvertisement(t, owner.Token, fixtures.AdvertisementRequest(t))

	t.Log("Добавляем чужое объявление в избранное")
	fixtures.AddFavorite(t, reader.Token, ad.ID)
	list := favorites.GetFavorites(t, reader.Token, ad.Title)
	assertions.FavoritePresence(t, list, ad.ID, true)
	require.Equal(t, 1, list.Total, "Объявление должно встретиться в избранном один раз")

	t.Log("Убираем объявление из избранного")
	favorites.RemoveFavorite(t, reader.Token, ad.ID)
	assertions.FavoritePresence(t, favorites.GetFavorites(t, reader.Token, ad.Title), ad.ID, false)

	t.Log("Снова добавляем в избранное, затем удаляем само объявление")
	fixtures.AddFavorite(t, reader.Token, ad.ID)
	assertions.FavoritePresence(t, favorites.GetFavorites(t, reader.Token, ad.Title), ad.ID, true)
	ad.Delete(t)
	list = favorites.GetFavorites(t, reader.Token, ad.Title)
	assertions.FavoritePresence(t, list, ad.ID, false)
	require.Zero(t, list.Total, "Удалённое объявление осталось в избранном")
	t.Log("Удалённое объявление исчезло из избранного")
}
