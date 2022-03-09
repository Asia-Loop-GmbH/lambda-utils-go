package dbadmin

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductType string

const (
	ProductTypeSimple   ProductType = "simple"
	ProductTypeVariable ProductType = "variable"
)

type ProductAttribute struct {
	WPID      int      `bson:"id" json:"wpId"` // it's ok in patch request because we won't change this
	Name      string   `bson:"name" json:"name"`
	Position  int      `bson:"position" json:"position"`
	Variation bool     `bson:"variation" json:"variation"`
	Visible   bool     `bson:"visible" json:"visible"`
	Options   []string `bson:"options" json:"options"`
}

type ProductVariationAttribute struct {
	WPID   int    `bson:"id" json:"wpId"` // it's ok in patch request because we won't change this
	Name   string `bson:"name" json:"name"`
	Option string `bson:"option" json:"option"`
}

type ProductVariation struct {
	WPID         int                         `bson:"id" json:"wpId"` // it's ok in patch request because we won't change this
	Price        string                      `bson:"price" json:"price"`
	RegularPrice string                      `bson:"regularPrice" json:"regularPrice"`
	SalePrice    string                      `bson:"salePrice" json:"salePrice"`
	Attributes   []ProductVariationAttribute `bson:"attributes" json:"attributes"`
}

type Product struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	WPID         int                `bson:"id" json:"wpId"` // it's ok in patch request because we won't change this
	Name         string             `bson:"name" json:"name"`
	Permalink    string             `bson:"permalink" json:"permalink"`
	Type         ProductType        `bson:"type" json:"type"`
	SKU          string             `bson:"sku" json:"sku"`
	Price        string             `bson:"price" json:"price"`
	RegularPrice string             `bson:"regularPrice" json:"regularPrice"`
	SalePrice    string             `bson:"salePrice" json:"salePrice"`
	TaxClass     string             `bson:"taxClass" json:"taxClass"`
	Categories   []string           `bson:"categories" json:"categories"`
	Images       []string           `bson:"images" json:"images"`
	Attributes   []ProductAttribute `bson:"attributes" json:"attributes"`
	Variations   []ProductVariation `bson:"variations" json:"variations"`
	OutOfStockIn []string           `bson:"outOfStockIn" json:"outOfStockIn"`
	Pfand        string             `bson:"pfand" json:"pfand"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt" json:"updatedAt"`
}
