# Backend Test

### [POST] /workflow
Inserts a new workflow in database

Body
```json
{
	"name": "Testing",
	"status": "inserted",
	"data": {
		"simple": "data"
	},
	"steps": [
		{
			"name": "Testing Step 1"
		},
		{
			"name": "Testing Step 2"
		}
	]
}
```

### [GET] /workflow
Lists all workflows in database

### [PATCH] /workflow/{uuid}
Updates the workflows's status

Body
```json
{
	"status": "consumed"
}
```

### [GET] /workflow/consume
Consumes a workflow from queue and generate a CSV file
