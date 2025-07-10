package is a way to organise the code.
if function_name, variable_name, type_name visible to outside package if first letter is uppercase

Module is a collection of Go packages

go.mod file: This file defines the module's path, the Go version it requires, and its dependencies 

go mod init [modulepath] : initialize a new module

go mod tidy : adds missing module dependencies and removes unused one

go mod download: Downloads the modules required by your current module.

go get [packagepath]: Adds a new dependency or updates an existing one to its latest version. 

go.sum: This file contains cryptographic checksums of the contents of your module's dependencies

standard libraries

io: Basic interfaces for I/O primitives.

os: Operating system functionalities (files, environment variables, processes).

path/filepath: Utilities for manipulating file paths.

net/http: For building web servers and clients.

json: Encoding and decoding JSON data.

encoding/gob, encoding/xml: Other encoding formats.

bufio: Buffered I/O.

sort: Sorting slices and collections.

regexp: Regular expressions.

sync: Advanced synchronization primitives (Mutexes, WaitGroups - useful when channels aren't the best fit for shared memory).

context: For managing deadlines, cancellations, and request-scoped values across API boundaries and goroutines.

testing: The built-in testing framework.