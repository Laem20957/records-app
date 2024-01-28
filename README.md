# The application let create records using REST API
### Before using you need to install:
1. Golang v1.21.5
2. Swagger
3. Docker-compose
4. Makefile
5. Migrations
### This App contains the following methods:
[post] /auth/sign-up - to create new user.<br />
[post] /auth/sign-in - user authentication.<br />
[get] /api/record - get all records.<br />
[post] /api/record - create new record.<br />
[get] /api/record/{id} - get record by id.<br />
[put] /api/record/{id} - update record by id.<br />
[delete] /api/record/{id} - delete record by id.<br />
### Installing:
```
git clone https://github.com/Laem20957/records-app.git
```
```
make migrate
```
```
make build
```
```
make run
```
### After launching the application will available on this link - localhost:8080/swagger/index.html

