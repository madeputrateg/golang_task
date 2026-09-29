package dto

import "github.com/Teepee/internal/product/dto"

type ViewShopItemResponse struct {
	Name     string                `json:"name"`
	Location string                `json:"location"`
	ShopId   string                `json:"id"`
	ImgPath  string                `json:"img_path"`
	Items    dto.ProductSimpleList `json:"items"`
}

type ShopDetail struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	ShopId   string `json:"id"`
	ImgPath  string `json:"img_path"`
}

type ShopDetailList []ShopDetail
