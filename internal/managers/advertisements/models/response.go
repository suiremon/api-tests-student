package models

type AdvertisementsList struct {
	Items      []Advertisement `json:"items"`
	Total      int             `json:"total"`
	HasNext    bool            `json:"has_next"`
	NextCursor string          `json:"next_cursor"`
}

type Advertisement struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
