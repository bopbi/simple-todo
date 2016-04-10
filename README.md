# simple-todo
A Simple REST Server app, using golang and sqlite, build for testing client app, the SQLITE file already contain some example data

## precondition
go-sqlite3

    go get github.com/mattn/go-sqlite3

pat (for routing)

    go get github.com/bmizerany/pat

## REST API
* GET ALL

		GET todos
* GET BY ID, example id is 12

		GET todos/12
* INSERT, it will return the json for new todo

		POST todos
* UPDATE BY ID, example id is 12 and it will return the json for new todo

		PUT todos/12
* DELETE BY ID, example id is 12

		DELETE todos/12
