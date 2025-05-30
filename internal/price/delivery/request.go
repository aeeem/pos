package delivery

type SaveOrUpdate struct {
	Price  int64 `json:"price"`
	Active bool  `json:"active"`
	ItemID int64 `json:"item_id"`
}
