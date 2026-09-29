package global

const (
	RouteAddItemProduct        string = "/products/additem"
	RouteViewItemDetail        string = "/products/viewitemdetail/:item_id"
	RouteViewItemList          string = "/products/viewitemlist"
	RouteUpdateItemAvailabilty string = "/products/updateavailabilty"
	RouteAddItemStock          string = "/products/additemstock"
	RouteProductCheckout       string = "/products/:id/checkout"
	RouteAddShop               string = "/shops/addshop"
	RouteViewShopList          string = "/shops/viewshoplist"
	RouteViewShopItemList      string = "/shops/:shop_id/viewitem"
	RouteSignIn                string = "/users/signin"

	AvailableAvailable     string = "available"
	AvailableNot_available string = "not_available"
	AvailableOut_of_stock  string = "out_of_stock"
)
