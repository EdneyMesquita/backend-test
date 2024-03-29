package models

import (
	"encoding/json"
	"log"
)

type (
	//Workflow is the model for converting results to JSON
	Workflow struct {
		Name   string                 `json:"name"`
		Status string                 `json:"status"`
		Data   map[string]interface{} `json:"data"`
		Steps  []Step                 `json:"steps"`
	}

	//WorkflowStatus is the model for mapping data from request to struct
	WorkflowStatus struct {
		Status string `json:"status"`
	}

	//FullWorkflow is the model for converting results to JSON
	FullWorkflow struct {
		UUID   string `json:"uuid"`
		Name   string `json:"name"`
		Status string `json:"status"`
		Data   string `json:"data"`
		Steps  []Step `json:"steps"`
	}

	//Workflows is the model for converting results to JSON
	Workflows struct {
		Workflows []FullWorkflow `json:"workflows"`
	}
)

//ToJSON convets the user struct to JSON
func (workflows *Workflows) ToJSON() string {
	json, err := json.Marshal(workflows)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return string(json)
}
