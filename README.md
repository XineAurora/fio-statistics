# fio-statistics
 
migration tool:  
'go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest'  
run migration script:  
'migrate -database {DB_URL} -path ./internal/database/migration/ up'  
 
