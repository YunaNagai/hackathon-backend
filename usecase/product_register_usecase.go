package usecase

import (
	"database/sql"
	"encoding/json"
	"hackathon-backend/dao"
	"hackathon-backend/model"
	"hackathon-backend/utils"
	"net/http"
)

func RegisterProducts(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// リクエスト用の struct（DB 用の Product とは別）
	var req model.ProductRegisterRequest

	// JSON デコード
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid json"})
		return
	}

	// バリデーション
	if req.SellerID == "" || req.Title == "" || req.Price <= 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid fields"})
		return
	}

	// DB に保存する Product を組み立てる
	product := model.Product{
		ID:          utils.NewULID(),
		SellerID:    req.SellerID,
		Title:       req.Title,
		Price:       req.Price,
		Description: req.Description,
		Status:      "selling",
		ImageURL:    req.ImageURL,
		CreatedAt:   utils.NowString(),
	}

	// DB INSERT
	if err := dao.InsertProduct(db, product); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "db error"})
		return
	}

	// 成功レスポンス
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": product.ID})
}
