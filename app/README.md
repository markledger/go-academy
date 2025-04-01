### Running the Application
From the root folder run: `go run main.go `

### Running the command_line package

This must first be built

`go build ./internal/command_line/command_line.go`

which will place the executable in the project root


`./command_line 9 7 45 89 `


### Tidying modules

`go mod tidy`

go mod tidy ensures that the go.mod file matches the source code in the module. It adds any missing module requirements necessary to build the current module’s packages and dependencies, and it removes requirements on modules that don’t provide any relevant packages. It also adds any missing entries to go.sum and removes unnecessary entries.%    
