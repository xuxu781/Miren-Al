package lty_controllers

import (
	"encoding/json"
	"net/http"

	"lty_backend/lty_config"
)

type CategoryItem struct {
	Name string   `json:"name"`
	Sub  []string `json:"sub"`
}

// SyncInspirationCategories overrides all categories with the provided list
func SyncInspirationCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var categories []CategoryItem
	if err := json.NewDecoder(r.Body).Decode(&categories); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "参数错误"})
		return
	}

	tx, err := lty_config.DB.Begin()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "开启事务失败"})
		return
	}
	defer tx.Rollback()

	// 1. Clear existing
	_, err = tx.Exec("DELETE FROM lty_inspiration_categories")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "清理旧数据失败"})
		return
	}

	// 2. Insert new
	stmt, err := tx.Prepare("INSERT INTO lty_inspiration_categories (name, sub_categories, sort_order) VALUES (?, ?, ?)")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "准备插入语句失败"})
		return
	}
	defer stmt.Close()

	for i, cat := range categories {
		subBytes, _ := json.Marshal(cat.Sub)
		if cat.Sub == nil {
			subBytes = []byte("[]")
		}
		_, err = stmt.Exec(cat.Name, string(subBytes), i)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "插入数据失败: " + err.Error()})
			return
		}
	}

	if err := tx.Commit(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "提交事务失败"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "同步成功"})
}

// GetAllInspirationCategories gets the categories array
func GetAllInspirationCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := lty_config.DB.Query("SELECT name, sub_categories FROM lty_inspiration_categories ORDER BY sort_order ASC")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "查询失败: " + err.Error()})
		return
	}
	defer rows.Close()

	var result []CategoryItem
	for rows.Next() {
		var name string
		var subStr string
		if err := rows.Scan(&name, &subStr); err != nil {
			continue
		}
		var sub []string
		if subStr != "" {
			json.Unmarshal([]byte(subStr), &sub)
		}
		if sub == nil {
			sub = []string{}
		}
		result = append(result, CategoryItem{
			Name: name,
			Sub:  sub,
		})
	}

	if result == nil {
		result = []CategoryItem{}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": result,
	})
}
