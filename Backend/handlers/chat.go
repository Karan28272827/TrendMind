package handlers

import (
	"encoding/json"
	"net/http"
	"server/models"
	"database/sql"
)

func StoreUserQueryHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
			return
		}

		var userQuery models.UserQueries

		if err := json.NewDecoder(r.Body).Decode(&userQuery); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		var productResponses []string

		rows, err := db.Query(`
			SELECT product_name 
			FROM products 
			WHERE product_name ILIKE $1 
			LIMIT 10
		`, "%"+userQuery.Query+"%")
		if err != nil {
			http.Error(w, "Error querying products: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var productName string
			if err := rows.Scan(&productName); err != nil {
				http.Error(w, "Error scanning row: "+err.Error(), http.StatusInternalServerError)
				return
			}
			productResponses = append(productResponses, productName)
		}

		if len(productResponses) == 0 {
			http.Error(w, "No matching products found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"user_id":     userQuery.UserID,
			"chat_id":     userQuery.ChatID,
			"query":       userQuery.Query,
			"response":    productResponses,
			"time_stamps": userQuery.TimeStamps,
		}
		json.NewEncoder(w).Encode(response)
	}
}

func GetProductsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`
			SELECT id, search_term, product_name, ctr_last_7_days, city_id, query_type_head, 
				   ctr_product_30_days, total_clicks, query_type_tail,
				   product_ctr_city_30_days, category_name, session_views, ctr_last_30_days,
				   query_products_clicks_last_30_days, total_unique_orders, query_product_similarity,
				   subcategory_name, is_clicked, savings, latest_margin, query_product_plt_clicks_30_days,
				   savings_with_pass, ad_revenue, product_atcs_30_days
			FROM products
		`)
		if err != nil {
			http.Error(w, "Error fetching products: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var products []models.Product
		for rows.Next() {
			var p models.Product
			err := rows.Scan(
				&p.ID, &p.SearchTerms, &p.ProductName, &p.CTRLast7Days, &p.CityID, &p.QueryTypeHead,
				&p.CTRProduct30Days, &p.TotalClicks, &p.QueryTypeTail,
				&p.ProductCTRCity30Days, &p.CategoryName, &p.SessionViews, &p.CTRLast30Days,
				&p.QueryProductsClicksLast30Days, &p.TotalUniqueOrders, &p.QueryProductSimilarity,
				&p.SubcategoryName, &p.IsClicked, &p.Savings, &p.LatestMargin,
				&p.QueryProductPLTClicks30Days, &p.SavingsWithPass, &p.AdRevenue, &p.ProductATCs30Days,
			)
			if err != nil {
				http.Error(w, "Error scanning product: "+err.Error(), http.StatusInternalServerError)
				return
			}
			products = append(products, p)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}
}
