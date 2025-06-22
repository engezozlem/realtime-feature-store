package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/engezozlem/realtime-feature-store/store"
)

type FeatureResponse struct {
	EntityID string            `json:"entity_id"`
	Features map[string]string `json:"features"`
}

func FeatureHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 || parts[2] == "" {
		http.Error(w, "Missing entity ID", http.StatusBadRequest)
		return
	}

	entityID := parts[2]
	features, err := store.Rdb.HGetAll(store.Ctx, entityID).Result()
	if err != nil || len(features) == 0 {
		features = map[string]string{
			"user_age": "29",
			"country":  "US",
		}
	}

	resp := FeatureResponse{
		EntityID: entityID,
		Features: features,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
