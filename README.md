## Change all imports as per your local directory
## Please change all db configurations as per your local envoirment 

## Project Architecture (Clean Code)


        ├── db                                                   <-- Database Connections
        │   └── db.go 
        ├── main.go                                              <-- Entrypoint
        ├── mocks                                                <-- Mocks for testing
        │   ├── repositories_mock.go
        │   └── service_mock.go
        ├── Readme.md
        └── services                                             <-- Services listed here
            ├── cmd
            │   └── service
            │       └── service.go                                <-- DB connection and all dependancies injection done
            |                                                         here 
            └── pkg
                ├── endpoints                                     <-- Routes
                │   └── router.go
                ├── handlers                                      <-- Handler function
                │   ├── handlers.go
                │   └── handlers_test.go
                ├── models                                        <-- Models
                │   └── info.go
                ├── repositories                                  <-- database operations listed here
                │   ├── repositories.go
                │   └── repositories_test.go
                └── service                                       <-- business logic layer
                    ├── cities.go
                    └── service.go






## To run the project: 
         go build && ./intTask

## Find database in db folder 
         Create Database city
         And export cities.sql in database

## Endpoints:
        GET: localhost:9001/cities
        GET:  localhost:9001/cities/{cityName}