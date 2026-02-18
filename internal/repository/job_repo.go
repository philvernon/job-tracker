package repository

import (
	"database/sql"
	"github.com/philvernon/job-tracker/internal/models"
)

type JobRepository struct {
	DB *sql.DB
}

func NewJobRepository(db *sql.DB) *JobRepository {
	return &JobRepository{DB: db}
}

func (r *JobRepository) Create(job *models.Job) error {
	query := `
	INSERT INTO jobs (title, url, description)
	VALUES (?, ?, ?)
	`

	result, err := r.DB.Exec(query, job.Title, job.URL, job.ID, job.Description)

	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (r *JobRepository) GetAllJobs() ([]models.Job, error) {
	query := `
	SELECT id, url, title, description, date_added 
	FROM jobs
	ORDER BY date_added DESC
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var count int
	err = r.DB.QueryRow("SELECT COUNT(*) FROM jobs").Scan(&count)
	if err != nil {
		return nil, err
	}

	jobs := make([]models.Job, 0, count)

	for rows.Next() {
		if err = rows.Err(); err != nil {
			return nil, err
		}

		job := models.Job{}

		err := rows.Scan(&job.ID, &job.Title, &job.URL, &job.Description, &job.DateAdded)
		if err != nil {
			return nil, err
		}

		jobs = append(jobs, job)
	}

	return jobs, nil
}
