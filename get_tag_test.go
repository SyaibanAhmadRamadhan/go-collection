package collection_test

import (
	"fmt"
	"github.com/SyaibanAhmadRamadhan/go-collection"
	"testing"
	"time"
)

type ProductItem struct {
	ID                  int64               `db:"id" json:"id"`
	UserID              int64               `db:"user_id" json:"user_id"`
	Variant1            *ProductItemVariant `db:"variant_1" json:"variant_1"`
	Variant2            *ProductItemVariant `db:"variant_2" json:"variant_2"`
	SubCategoryItemName string              `db:"sub_category_item_name" json:"sub_category_item_name"`
	Name                string              `db:"name" json:"name"`
	Description         string              `db:"description" json:"description"`
	Price               float64             `db:"price" json:"price"`
	Stock               int32               `db:"stock" json:"stock"`
	Sku                 *string             `db:"sku" json:"sku"`
	Weight              int32               `db:"weight" json:"weight"`
	PackageLength       float64             `db:"package_length" json:"package_length"`
	PackageWidth        float64             `db:"package_width" json:"package_width"`
	PackageHeight       float64             `db:"package_height" json:"package_height"`
	DimensionalWeight   float64             `db:"dimensional_weight" json:"dimensional_weight"`
	IsActive            bool                `db:"is_active" json:"is_active"`
	ProductCondition    string              `db:"product_condition" json:"product_condition"`
	MinimumPurchase     int32               `db:"minimum_purchase" json:"minimum_purchase"`
	SizeGuideImage      *string             `db:"size_guide_image" json:"size_guide_image"`
	CreatedAt           time.Time           `db:"created_at" json:"created_at"`
	UpdatedAt           time.Time           `db:"updated_at" json:"updated_at"`
	DeletedAt           *time.Time          `db:"deleted_at" json:"deleted_at"`
}

type ProductItemVariant struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func TestGetTags(t *testing.T) {
	product := ProductItem{}
	tags := collection.GetTags(product, "db", "id")
	fmt.Println(tags) // O
}
