package workflows

import (
	"backend-test/models"
	"backend-test/server/database"
	"backend-test/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Change updates workflows' status
func Change(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var uuid string
	if uuid = params["uuid"]; uuid == "" {
		utils.HTTPResponse(w, http.StatusInternalServerError, "Workflow's uuid was not passed", false)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println(err.Error())
		utils.HTTPResponse(w, http.StatusInternalServerError, "It wasn't possible convert Request Body to JSON", false)
		return
	}

	var status models.WorkflowStatus
	err = json.Unmarshal(body, &status)
	if err != nil {
		log.Println(err.Error())
		utils.HTTPResponse(w, http.StatusInternalServerError, "An error was ocurred while updating workflow. Please, try again.", false)
		return
	}

	if status.Status == "inserted" || status.Status == "consumed" {
		sql := utils.BuildString(`UPDATE workflows SET "status" = '`, status.Status, `' WHERE "uuid" = '`, uuid, `'`)

		//Updating workflow's status in database
		stmt, err := database.Conn.Exec(sql)
		if err != nil {
			log.Println(err.Error())
			utils.HTTPResponse(w, http.StatusInternalServerError, "An error was ocurred while updating data. Please, try again.", false)
			return
		}

		//Checking if workflow was updated successfully
		rows, _ := stmt.RowsAffected()
		if rows == 0 {
			log.Println(err.Error())
			utils.HTTPResponse(w, http.StatusInternalServerError, "It wasn't possible to update workflow", false)
			return
		}
	} else {
		log.Println(err.Error())
		utils.HTTPResponse(w, http.StatusInternalServerError, "Invalid value for status. It must be 'inserted' or 'consumed'.", false)
		return
	}

	utils.HTTPResponse(w, http.StatusOK, "Workflow updated successfully!", false)
}
