package dto

type AddShopRequest struct {
	Name     string `form:"name"`
	Location string `form:"location"`
	// Img *multipart.FileHeader
}

type GetShopItemRequest struct {
	ShopId string `param:"shop_id"`
}
