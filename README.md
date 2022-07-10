# URL Shortener

## Deploy
### Requirements
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Makefile](https://stackoverflow.com/questions/3915067/what-are-makefiles-make-install)

### 1. git clone
```
git clone https://github.com/mreza0100/shortly
```
### 2. cd into the directory
```
cd shortly
```
### 3. copy and config(no need for simple deploy) ".env.example" to ".env"
```
cp ./.env.example ./.env
```
### 4. run dependencies and the app
```
make run
```
## Start Development
### Requirements

- [Go Programming Language](https://go.dev/doc/install)
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Makefile](https://stackoverflow.com/questions/3915067/what-are-makefiles-make-install)
- [Air(A tool to live reload Go apps)](https://github.com/cosmtrek/air)


### 1. git clone
```
git clone https://github.com/mreza0100/shortly
```
### 2. cd into the directory
```
cd shortly
```
### 3. copy and config(unless you need) ".env.example" to ".env"
```
cp ./.env.example ./.env
```
### 4. run dependencies
```
make dev-dependency
```
### 5. start air
```
make dev
```
# Talk is cheap, show me the code.

#### It is assumed that link shortners are heavy read applications.
#### 1% write requests and 99% read requests.
#### The last target of this app is to response 10000 requests per second, but not for now as I did't had enough time to implement it. but I will try to do it in the future and I will explain the idea of the future app.

# How to start reading the code?
## The app have 3 actions:
- run:         ``` go run ./cmd/shortly run```
- seed:        ``` go run ./cmd/shortly seed```
- healthcheck: ``` go run ./cmd/shortly healthcheck```

#### The ```readme.md``` file is implemented in each directory.
#### You can start from cmd/shortly directory.

### Actions:
- run: run the http server with the given port from the environment variable.
- seed: seed the database with random generated data.
- healthcheck: check the health of the app. healthcheck is implemented for docker-compose.yml file to make sure the connection between app and database is working.

# Tests:
#### To run the tests: ```make test```
#### Unfortunately, I did't had time to implement the integration tests. But I considered tests and mocking dependencies from the beginning of the project in the dependency injections.

# Structure
### The architecture of the app is based on the following structure:
- Hexagonal Architecture
- Domain Driven Design
