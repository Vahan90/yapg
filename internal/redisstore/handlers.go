package redisstore

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func PushHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	lines := strings.Split(string(body), "\n")
	var toStore []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		toStore = append(toStore, line)
	}

	id := uuid.New().String()
	key := "metric:" + id
	if err := Rdb.RPush(Ctx, key, toStore).Err(); err != nil {
		http.Error(w, "failed to store metrics", http.StatusInternalServerError)
		return
	}

	// Rdb.Expire(Ctx, key, 1*time.Second)
	w.WriteHeader(http.StatusAccepted)
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	keys, err := Rdb.Keys(Ctx, "metric:*").Result()
	if err != nil {
		http.Error(w, "failed to query metrics", http.StatusInternalServerError)
		return
	}

	for _, key := range keys {
		metrics, err := Rdb.LRange(Ctx, key, 0, -1).Result()
		if err != nil {
			continue
		}

		for _, line := range metrics {
			fmt.Fprintln(w, line)
		}

		Rdb.Del(Ctx, key)
	}
}
