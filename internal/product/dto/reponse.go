package dto

type ProductSimple struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Stock     uint32 `json:"stock"`
	Available string `json:"available"`
}

type ProductSimpleList []ProductSimple

type ProductDetail struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Price     uint32 `json:"price"`
	Stock     uint32 `json:"stock"`
	Available string `json:"available"`
	ShopId    string `json:"shop_id"`
	// img
}
