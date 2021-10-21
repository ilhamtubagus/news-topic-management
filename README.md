# Domain Driven Design Golang
This project demonstrates how we can implement domain design approach in developing golang RESTful API. As described by Eric Evans in his book, Domain-Driven Design is an approach to software development that centers the development on programming a domain model that has a rich understanding of the processes and rules of a domain.

## Case Study
- News and topic management.
- CRUD on **news** and **tags**.
- One news can contains multiple tags e.g. "Safe investment" might contains tags
"investment", "mutual fund", etc.
- One topic has multiple news e.g. "investment" topic might contains "how to start
investment", "mutual fund is safe investment type", etc.
- Enable filter by news status ("draft", "deleted", "publish").
- Enable filter news by its topics.
- At the first request API call from the client, the service should get data from the database
and stored onto redis.

## How to Run
1. Clone this repo.
2. Install go dependecies.
```
go mod download
```
3. Build and run.
```
docker-compose up
```
**Make sure docker and docker-compose installed on your computer.**
