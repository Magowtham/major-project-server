package repository

import (
	"database/sql"

	"github.com/Magowtham/dehydrater-server/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Init() error {
	query := `CREATE TABLE IF NOT EXISTS steps (
				step_number VARCHAR(255) PRIMARY KEY,
				step_time VARCHAR(255) NOT NULL,
				step_temp VARCHAR(255) NOT NULL
			)`

	_, err := r.db.Exec(query)

	return err
}

func (r *PostgresRepository) AddStep(steps []*models.DeviceStep) error {
	for _, step := range steps {
		query := `INSERT INTO steps (step_number,step_time,step_temp) VALUES ($1,$2,$3)`
		if _, err := r.db.Exec(query, step.StepNumber, step.Time, step.Temperature); err != nil {
			return err
		}
	}
	return nil
}

func (r *PostgresRepository) GetSteps() (*models.DeviceStepResponse, error) {
	query1 := `SELECT step_number,step_time,step_temp FROM steps`

	rows, err := r.db.Query(query1)

	if err != nil {
		return nil, err
	}

	var deviceStepResponse models.DeviceStepResponse

	for rows.Next() {
		var deviceStep models.DeviceStep

		if err := rows.Scan(&deviceStep.StepNumber, &deviceStep.Time, &deviceStep.Temperature); err != nil {
			return nil, err
		}

		deviceStepResponse.Steps = append(deviceStepResponse.Steps, &deviceStep)
	}

	return &deviceStepResponse, nil
}

func (r *PostgresRepository) DeleteSteps() error {
	query2 := `DELETE FROM steps`

	if _, err := r.db.Exec(query2); err != nil {
		return err
	}

	return nil
}
