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

project layout and architecture
cmd directory hold the source code for the main app. 
usage : cmd/<app-name>/main.go

internal directory : contains private application code that you want other projects to import

pkg : reusable, public libraries that external projects can import

api : contains API specifications
web : client side code
scripts : automation scripts
build : Dockerfile and K8s configs 
configs : configuration file
test : test files

my-go-project/
├── go.mod
├── go.sum
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── api/
│   │   └── handlers.go
│   └── database/
│       └── db.go
├── web/
│   ├── static/
│   │   └── style.css
│   └── templates/
│       └── index.html
└── README.md

go fmt : Code formatting

go vet : Catching Suspicious Code

go build : compile the package

//cross compilation
GOOS=linux GOARCH=amd64 go build -o my_api_linux ./cmd/api

Here is a breakdown of their differences:

1. Worker Pool
Primary Goal: To control concurrency and limit resource usage.

Structure: A fixed number of "worker" goroutines are created at the start. These workers wait for tasks to be sent to them via a shared jobs channel.

Concurrency: The level of concurrency is fixed and determined by the number of workers. This prevents the system from being overwhelmed if a large number of tasks arrive at once.

Best For: I/O-bound tasks or when you need to prevent resource exhaustion (e.g., from too many database connections, HTTP requests, or memory allocations). It is the ideal pattern for handling a massive, potentially unbounded number of jobs with a limited set of resources.

Analogy: A factory with a fixed number of workers. Orders (jobs) arrive in a queue. The workers pick up the next available job, complete it, and then go back to the queue to get another one.

2. Fan-In/Fan-Out
Primary Goal: To parallelize a set of tasks and then aggregate their results.

Structure:

Fan-Out: A single "producer" goroutine distributes a stream of work items to a separate goroutine for each item. This creates a "fan" of concurrently executing goroutines. The number of goroutines is dynamic and scales with the number of tasks.

Fan-In: A single "collector" goroutine aggregates the results from multiple results channels (one from each of the fan-out goroutines) into a single unified results channel. This consolidates the "fan" back into a single stream.

Concurrency: The level of concurrency is dynamic and corresponds to the number of tasks. For N tasks, you might launch N goroutines.

Best For: CPU-bound tasks where you want to fully utilize all available cores. It is effective when the number of tasks is known and not excessively large, and the work can be easily divided.

Analogy: A main office (producer) sends out individual packages (jobs) to a large number of delivery drivers (goroutines). Once the drivers deliver their packages, they all report their completion status to a single manager (fan-in) who logs all the results.