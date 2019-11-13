package workflows

import (
	"backend-test/models"
	"backend-test/server/database"
	"backend-test/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//Create inserts a new workflow in database
func Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println(err.Error())
		utils.HTTPResponse(w, http.StatusInternalServerError, "It wasn't possible convert Request Body to JSON", false)
		return
	}

	var workflow models.Workflow
	err = json.Unmarshal(body, &workflow)

	var sql string

	//Convert data: map[string]interface{} to JSON
	data, err := json.Marshal(workflow.Data)
	if err != nil {
		log.Println(err.Error())
		utils.HTTPResponse(w, http.StatusInternalServerError, "An error was ocurred while inserting data. Please, try again.", false)
		return
	}

	sql = utils.BuildString(`INSERT INTO workflows ("uuid", "name", "status", "data") VALUES ((SELECT uuid_generate_v4()), '`, workflow.Name, `', '`, workflow.Status, `', '`, string(data), `');`)

	//Inserting workflow in database
	stmt, err := database.Conn.Exec(sql)
	if err != nil {
		log.Println(err.Error())
		utils.HTTPResponse(w, http.StatusInternalServerError, "An error was ocurred while inserting data. Please, try again.", false)
		return
	}

	//Checking if workflow was inserted successfully
	rows, _ := stmt.RowsAffected()
	if rows == 0 {
		log.Println(err.Error())
		utils.HTTPResponse(w, http.StatusInternalServerError, "It wasn't possible to insert workflow", false)
		return
	}

	var workflowID int64
	database.Conn.QueryRow("SELECT MAX(id) FROM workflows").Scan(&workflowID)

	sql = ""
	for _, step := range workflow.Steps {
		sql += utils.BuildString(`INSERT INTO steps ("name", "workflow") VALUES ('`, step.Name, `', `, workflowID, `);`)
	}

	stmt, err = database.Conn.Exec(sql)
	if err != nil {
		log.Println(err.Error())
		utils.HTTPResponse(w, http.StatusInternalServerError, "An error was ocurred while inserting data. Please, try again.", false)
		return
	}

	rows, _ = stmt.RowsAffected()
	if rows == 0 {
		log.Println(err.Error())
		utils.HTTPResponse(w, http.StatusInternalServerError, "It wasn't possible to insert steps", false)
		return
	}

	//Reload queue
	MountQueue()
	utils.HTTPResponse(w, http.StatusOK, "Workflow inserted successfully!", false)
}
