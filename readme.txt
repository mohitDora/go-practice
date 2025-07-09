package is a way to organise the code.
if function_name, variable_name, type_name visible to outside package if first letter is uppercase

Module is a collection of Go packages

go.mod file: This file defines the module's path, the Go version it requires, and its dependencies 

go mod init [modulepath] : initialize a new module

go mod tidy : adds missing module dependencies and removes unused one

go mod download: Downloads the modules required by your current module.

go get [packagepath]: Adds a new dependency or updates an existing one to its latest version. 

go.sum: This file contains cryptographic checksums of the contents of your module's dependencies