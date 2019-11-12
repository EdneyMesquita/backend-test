package workflows

import (
	"backend-test/models"
	"backend-test/server/database"
	"backend-test/utils"
	b64 "encoding/base64"
	"log"
	"net/http"
	"strings"
)

//List returns a whole list of workflows
func List(w http.ResponseWriter, r *http.Request) {
	sql := utils.BuildString(`SELECT w."uuid", w."name", w."data", w."status", STRING_AGG('{"name": "' || s."name" || '"}', ', ') as step from workflows w
	LEFT JOIN steps s ON(s.workflow = w.id)
	GROUP BY w.id`)

	rows, err := database.Conn.Query(sql)
	if err != nil {
		log.Println(err.Error())
		utils.HTTPResponse(w, http.StatusInternalServerError, "Cannot connect to database. Try again", false)
		return
	}

	var (
		workflows []models.FullWorkflow
		data      string
		steps     string
	)
	for rows.Next() {
		var workflow models.FullWorkflow
		rows.Scan(
			&workflow.UUID,
			&workflow.Name,
			&data,
			&workflow.Status,
			&steps,
		)
		workflow.Data, _ = b64.StdEncoding.DecodeString(data)
		if steps != "" {
			workflow.Steps = strings.Split(steps, ",")
		}
		workflows = append(workflows, workflow)
	}

	workflowsList := &models.Workflows{
		Workflows: workflows,
	}

	utils.HTTPResponse(w, http.StatusOK, workflowsList.ToJSON(), true)
}
