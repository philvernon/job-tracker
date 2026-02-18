package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/philvernon/job-tracker/internal/db"
	"github.com/philvernon/job-tracker/internal/models"
	"github.com/philvernon/job-tracker/internal/repository"
)

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5174")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	db, err := db.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	jobRepo := repository.NewJobRepository(db)

	mux := http.NewServeMux()

	mux.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		jobUrl := r.URL.Query().Get("jobUrl")
		if jobUrl == "" {
			http.Error(w, "Job URL is missing", http.StatusBadRequest)
			return
		}

		job := models.Job{
			URL:         jobUrl,
			Description: "This is the description for test job",
			Title:       "TestJob",
		}

		err = jobRepo.Create(&job)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to add job", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(job)
	})

	mux.HandleFunc("/getAll", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		allJobs, err := jobRepo.GetAllJobs()
		if err != nil {
			http.Error(w, "Failed to get jobs", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(allJobs)
	})

	handler := withCORS(mux)

	log.Println("Server running on :8080")
	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal(err)
	}
}
