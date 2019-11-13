package workflows

import (
	"backend-test/models"
	"backend-test/utils"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

//Consume consumes a workflow from queue and generate a CSV file with workflow.Data
func Consume(w http.ResponseWriter, r *http.Request) {
	if Queue.Len() > 0 {
		e := Queue.Front()

		var jsonb map[string]interface{}
		workflow := e.Value.(models.FullWorkflow)
		json.Unmarshal([]byte(workflow.Data), &jsonb)

		file, err := os.Create("./workflow.csv")
		if err != nil {
			log.Println(err.Error())
			utils.HTTPResponse(w, http.StatusInternalServerError, "It wasn't possible to create CSV file", false)
			return
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		for key, value := range jsonb {
			r := make([]string, 0, 1+len(jsonb))
			r = append(r, key)
			r = append(r, fmt.Sprintf("%v", value))
			err := writer.Write(r)
			if err != nil {
				log.Println(err.Error())
				utils.HTTPResponse(w, http.StatusInternalServerError, "It wasn't possible to create CSV file", false)
				return
			}
		}

		Queue.Remove(e)
		utils.HTTPResponse(w, http.StatusOK, "CSV file generated successfully!", false)
	} else {
		utils.HTTPResponse(w, http.StatusOK, "There are any workflow in queue", false)
	}
}
