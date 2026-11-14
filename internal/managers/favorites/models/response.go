package models

import advModels "api-tests-template/internal/managers/advertisements/models"

type Favorite struct {
	ID            string                  `json:"favorite_id"`
	AddedAt       string                  `json:"added_at"`
	Advertisement advModels.Advertisement `json:"advertisement"`
}

type FavoritesList struct {
	Items      []Favorite `json:"items"`
	Total      int        `json:"total"`
	Page       int        `json:"page"`
	PageSize   int        `json:"page_size"`
	TotalPages int        `json:"total_pages"`
}
