package dto

// /products/addItem
// Method : post
// {
// Name:”string”
// Price:0219
// Stock:312321
// Available:”string”
// Shop_id:”string”
// img:file
// }

type AddItemRequest struct {
	Name      string `form:"name"`
	Price     uint32 `form:"price"`
	Stock     uint32 `form:"stock"`
	Available string `form:"available"`
	ShopId    string `form:"shop_id"`
	// Img       *multipart.FileHeader
}

type AddItemStockRequest struct {
	ItemId string `form:"item_id"`
	Stock  uint32 `form:"stock"`
}

type CheckoutProductRequest struct {
	Quantity uint32 `form:"quantity"`
	Id       string `param:"id"`
}

type UpdateAvailabiltyRequest struct {
	Id          string `form:"id"`
	Availabilty string `form:"availabilty"`
}
