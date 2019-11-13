package workflows

import (
	"backend-test/models"
	"backend-test/server/database"
	"backend-test/utils"
	"container/list"
	"log"
	"strings"
)

var Queue *list.List

func MountQueue() {
	Queue = list.New()

	sql := utils.BuildString(`SELECT w."uuid", w."name", w."data", w."status", STRING_AGG(s."name", ', ') as step from workflows w
	INNER JOIN steps s ON(s.workflow = w.id)
	GROUP BY w.id`)

	rows, err := database.Conn.Query(sql)
	if err != nil {
		log.Println(err.Error())
		return
	}

	var (
		workflows []models.FullWorkflow
		steps     string
	)
	for rows.Next() {
		var workflow models.FullWorkflow
		rows.Scan(
			&workflow.UUID,
			&workflow.Name,
			&workflow.Data,
			&workflow.Status,
			&steps,
		)
		if steps != "" {
			stepsSlice := strings.Split(steps, ",")
			var stepsMapped []models.Step
			for _, step := range stepsSlice {
				stepsMapped = append(stepsMapped, models.Step{
					Name: step,
				})
			}
			workflow.Steps = stepsMapped
		}
		workflows = append(workflows, workflow)
	}

	for _, workflow := range workflows {
		Queue.PushBack(workflow)
	}
}
