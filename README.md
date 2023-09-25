# fio-statistics

## Run appication
run application:
``go run ./cmd/fio-statistics/main.go``

## Migrate database
migration tool:  
``go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@lates``  
run migration script:  
``migrate -database {DB_URL} -path ./internal/database/migration/ up``  
 
 ## API
 Threre is 1 endpoint on ``/api/fio/``:
 - ``GET`` Returns list of fios and accepts filter parameters via request body and pagination parameters via query values. Query paremeters are ``page`` and ``onPage``. Filter example:  
 ```JSON
    {
        "name":"Joseph", 
	    "surname": "Joestar",
	    "patronymic": "JoJo",
	    "withoutPatronymic": true,
	    "lowerAge": 20,
	    "upperAge": 30,
	    "gender": "male",
	    "nationality": "US"
    }
 ```  
   All parameters are optional, but body must be provided.
 - ``POST`` Creates new fio provided via request body. Example:
  ```JSON
    {
        "name":"Joseph", 
	    "surname": "Joestar",
	    "patronymic": "JoJo", 
	    "age": 20,
	    "gender": "male",
	    "nationality": "US"
    }
 ``` 
 Patronymic is optional.
 - ``UPDATE`` Updates existing fio. Only id is mandatory, if empty field provided it value does not change. Example:  
  ```JSON
    {
        "id": 1,
        "name":"Joseph", 
	    "surname": "Joestar",
	    "patronymic": "JoJo", 
	    "age": 20,
	    "gender": "male",
	    "nationality": "US"
    }
 ``` 
 - ``DELETE`` Deletes existing fio by id. Request on ``/api/fio/1`` will delete fio with id 1.