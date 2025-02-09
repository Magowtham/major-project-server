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
	query1 := `CREATE TABLE IF NOT EXISTS steps (
				step_number VARCHAR(255) PRIMARY KEY,
				step_time VARCHAR(255) NOT NULL,
				step_temp VARCHAR(255) NOT NULL
			)`

	query2 := `CREATE TABLE IF NOT EXISTS analytics (
				step VARCHAR(255) PRIMARY KEY,
				temp VARCHAR(255) NOT NULL
			)`

	_, err := r.db.Exec(query1)

	if err != nil {
		return err
	}

	_, err = r.db.Exec(query2)

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

func (r *PostgresRepository) UpdateAnalytics(step, temp string) error {
	query := `
			INSERT INTO analytics (step, temp)
				VALUES ($1, $2)
				ON CONFLICT (step)
			DO UPDATE SET temp = EXCLUDED.temp;
`

	_, err := r.db.Exec(query, step, temp)
	return err
}

func (r *PostgresRepository) GetAnalytics() ([]*models.Analytics, error) {
	query := `SELECT step,temp FROM analytics`

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	var analytics []*models.Analytics

	for rows.Next() {
		var analytic models.Analytics

		if err := rows.Scan(&analytic.Step, &analytic.Temp); err != nil {
			return nil, err
		}

		analytics = append(analytics, &analytic)
	}

	return analytics, nil
}

func (r *PostgresRepository) DeleteAnalytics() error {
	query := `DELETE FROM analytics`

	_, err := r.db.Exec(query)

	return err
}
