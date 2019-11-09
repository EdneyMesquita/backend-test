package models

type (
	//Workflow is the model for converting results to JSON
	Workflow struct {
		UUID   int64    `json:"uuid"`
		Status string   `json:"status"`
		Data   string   `json:"data"`
		Steps  []string `json:"steps"`
	}
)
