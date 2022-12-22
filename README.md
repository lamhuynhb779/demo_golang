# Learn Golang

## Basic commands:

1. ```go mod init <path>``` : _create go.mod file that tracks/describes the module's properties, including its dependencies on other modules and versions of Go._

2. ```go run <path>``` :  _run function "main" in package "main"_

3. ```go help``` :  _is one of many go commands you'll use to get things done with Go_

4. ```go mod tidy [-v] | go get [-u] | go mod download```

5. ```go mod edit -replace <module-name>=<local-module-path>``` :  _go mod edit -replace example.com/greetings=../greetings_

6. ```go build``` :  _to compile the code into an executable_

7. ```go install``` : _to compile and install package (move file executable into Go PATH bin)_

8. ```go work init``` : _go work init ./hello_

9. ```go work use [-r] [dir]``` : _go work use ./example_

10. ```go work edit``` : _edits the go.work file similarly to go mod edit_

11. ```go work sync``` : _syncs dependencies from the workspace’s build list into each of the workspace modules._
