# yofio_backend
YoFio - Backend Golang


## Configuration

1 Import the sql that is located on schema/schema.sql
  ```sql
  mysql -u username -p database_name < schema/schema.sql
  ```
  
2 Edit the config file `.env`
  ``` 
DB_TYPE=mysql
USER=username
PASSWORD=password
DB_NAME=database_name
```
3 Build and run the project

```
GO111MODULE=on go build
./yofio_backend
```

Tests

For testing you should have the backend run to rull unit tests
```
cd tests
go test
```

Endpoints


POST → ​ /credit-assignment

```
curl -X POST -d '{"investment": 1500}' http://localhost:9090/api/statistics
```

POST → ​ /statistics ​

```
curl -X POST -d '{}' http://localhost:9090/api/statistics
```

