# README

This example demonstrate how use the OpenAPI generator [dredger](https://github.com/fraunhoferfokus/dredger). Create the OpenAPI.yml file, which contains all API endpoints for the todos service.

Then generate the code for the _todos_ service including the frontend code base using the _dredger_ API generator:

    dredger generate ./OpenAPI.yaml -o . -f -n todos

## Changed files

Only the following files has been adapted for the business logic and the front end:

-   entities/todosObj.go
-   usecases/todos.go
-   rest/addTodo.go
-   rest/deleteTodo.go
-   rest/doneTodo.go
-   rest/index.go
-   rest/root.go
-   rest/todos.go
-   rest/getReady.go

-   web/pages/index.templ
-   web/pages/todos.templ
-   web/pages/locales/\*

-   .gitignore

## Issues

-   TODO: Add tests
