
package models

type Product struct {
	ID                              int     `json:"id"`
	SearchTerms                     string `json:"search_terms"`
	ProductName                     string `json:"product_name"`
	CTRLast7Days                    float64 `json:"ctr_last_7_days"`
	CityID                          float64 `json:"city_id"`
	QueryTypeHead                   float64 `json:"query_type_head"`
	CTRProduct30Days                float64 `json:"ctr_product_30_days"`
	TotalClicks                     float64 `json:"total_clicks"`
	QueryTypeTail                   float64 `json:"query_type_tail"`
	ProductVariantID                float64 `json:"product_variant_id"`
	ProductCTRCity30Days            float64 `json:"product_ctr_city_30_days"`
	CategoryName                    string 	`json:"category_name"`
	SessionViews                    float64 `json:"session_views"`
	CTRLast30Days                   float64 `json:"ctr_last_30_days"`
	QueryProductsClicksLast30Days   float64 `json:"query_products_clicks_last_30_days"`
	TotalUniqueOrders               float64 `json:"total_unique_orders"`
	QueryProductSimilarity          float64 `json:"query_product_similarity"`
	SubcategoryName                 string `json:"subcategory_name"`
	IsClicked                       float64 `json:"is_clicked"`
	Savings                         float64 `json:"savings"`
	LatestMargin                    float64 `json:"latest_margin"`
	QueryProductPLTClicks30Days     float64 `json:"query_product_plt_clicks_30_days"`
	SavingsWithPass                 float64 `json:"savings_with_pass"`
	AdRevenue                       float64 `json:"ad_revenue"`
	ProductATCs30Days               float64 `json:"product_atcs_30_days"`
}


