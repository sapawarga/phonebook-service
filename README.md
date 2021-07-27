[![Go Report Card](https://goreportcard.com/badge/github.com/sapawarga/phonebook-service)](https://goreportcard.com/report/github.com/sapawarga/phonebook-service)
[![Maintainability](https://api.codeclimate.com/v1/badges/d620fba429567c496754/maintainability)](https://codeclimate.com/github/sapawarga/phonebook-service/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/d620fba429567c496754/test_coverage)](https://codeclimate.com/github/sapawarga/phonebook-service/test_coverage)
# phonebook-service
Sapawarga service for "Nomor Penting" feature.

## Quick Setup

1. Clone this repository with this command `git clone git@github.com:sapawarga/phonebook-service.git`
2. Use `.env` for environment variable and copy from `.env.example`
3. Use port that written in `docker-compose.yml`
4. Run `docker-compose up --build`

## Stack Libraries
1. [Gomock](https://github.com/golang/mock)
2. [Gokit](https://github.com/go-kit/kit)
3. [GRPC](https://grpc.io/docs/languages/go/basics/)
4. [Gorilla Mux](https://github.com/gorilla/mux)


## Package Structure
Golang use pattern, one directory is one package. This is repository's structure directory

```sh
cmd 
    - database
    - grpc
endpoint
repository
    - mysql
    - postgres
transport
    - grpc
    - http
usecase
    - phonebook
mocks
    - testcases
    - mock_repository
config
model
helper
```
### 1. CMD
Directory `cmd` acts as `infrastructure` of all entire service. Initialitation of database, thirdparty apps, routing, and module are in this package. This package using name `main` then this package will be the main package that client will access. 

### 2. Endpoint
Directory `endpoint` acts as encoding for request from client and response to client. Every validations request from clients must be here so each request that sent to usecase is clear and valid.

### 3. Transport
This directory has function for how the service accessing and how the service receive request then return response. This package has policy to arrange how request can be sent by any kind of protocols

### 4. Usecase
This directory is for create all business logic for the service. There is no more validation of the request from client. Moreover, this package is include like usecase, usecase test and usecase interface. Usecase test is for testing all business logic. It is using unit test

### 5. Mocks
This package is consisted of mocking and testcase. Mocking is generated from repository interface moreover it has function for simulating action from real repository. Testcase is for collecting any scenarios of every single unit test. It is using array for collecting all scenarios and in unit test just loop for each scenario.

### 6, Config
This package function is for configuration of any constant or variables that use in whole service.

### 7. Model
This package is for create struct that will be used in repository and usecase

### 8. Helper
This package is utility package that can be used without initialization first



