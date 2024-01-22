# The application let create records using REST API
### Before using you need to install:
1. Golang
2. Swagger
3. Docker-compose
4. Makefile
5. Migrations
### This App contains the following methods:
[post] /auth/sign-up - to create new user.<br />
[post] /auth/sign-in - user authentication.<br />
[get] /api/note - get all notes.<br />
[post] /api/note - create new note.<br />
[get] /api/note/{id} - get note by id.<br />
[put] /api/note/{id} - update note by id.<br />
[delete] /api/note/{id} - delete note by id.<br />
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

