package entity

type CreateEntityRequest struct {
	Name                     string   `json:"name"`
	Aliases                  []string `json:"aliases"`
	IngestionIntervalMinutes int      `json:"ingestion_interval_minutes"`
}

type EntityResponse struct {
	ID                       string   `json:"id"`
	Name                     string   `json:"name"`
	Aliases                  []string `json:"aliases"`
	Status                   string   `json:"status"`
	IngestionIntervalMinutes int      `json:"ingestion_interval_minutes"`
}
