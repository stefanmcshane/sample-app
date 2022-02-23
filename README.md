# Go-Template

## First Time Use

* `make init name=<SERVICE-NAME>`
* `make install-tools`


## Structure
* `service` contains all code for the Go service
* `common` contains common files used acrossed services, including a common makefile

## Service folder

Overall, this folder roughly follows typical Go patterns. There are no official rules on directory structure, 
so these should be strongly considered as guidelines and no more

* `cmd` contains the entrypoint code for the application .i.e server starting point, cli tool entry, or main.go in general
* `config` contains the package for parsing any config loaded into the Go program
* `internal` contains the business logic that explicitly [cannot be imported by other services](https://dave.cheney.net/2019/10/06/use-internal-packages-to-reduce-your-public-api-surface). This should be kept small, and not contain many packages nor public functions or interfaces
* `migrations` can be used for any SQL migrations needed for the service
* `secrets` dir will be loaded into locally run containers as a volume, but not loaded into the image itself
* `tools` contains any shared tooling that should not be imported by other modules, or cannot be abstracted to a module of its own .i.e  extending loggers functionality for this one service such as adding an extra field
* `api` can be used for any externally facing API, or RPC endpoint related code

## Running tests

This repo supports a high-level version of the [testing pyramid](https://martinfowler.com/bliki/TestPyramid.html), for Unit and Integration/Service level tests.

* `make test` will run all tests that do not have the build tag "integration"
* `make test-integration` will run all tests with the "integration" build tag


## Notes
* Add any build tools that are required for dev to `service/tools/tools.go`