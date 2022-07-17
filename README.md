# Shortly

### Shortly is a Link shortener service that allows you to shorten links and share them with your friends 😃
# Features
- Create short links with your account
- Share links with anyone
- Permanent redirect links

# 🤔 Future plans:
- Create custom short links
- Strong authentication system
- Delete links
- Edit links
- Edit account
- Group permissions
- Location and device type data
- Session management
- Branded links - with a custom domain or custom route
- Data export
- Link history and reporting
- Expiration for links
- Access links based on location

## Deploy
### Requirements
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Make](https://stackoverflow.com/questions/3915067/what-are-makefiles-make-install)

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
- [Make](https://stackoverflow.com/questions/3915067/what-are-makefiles-make-install)
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
# API Endpoints
## Signup
```
curl --location --request POST '10.0.0.10:10000/user/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "mreza@gmail.com",
    "password": "1234"
}'
```
## Login
```
curl --location --request POST '10.0.0.10:10000/user/signin' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "mreza@gmail.com",
    "password": "1234"
}'
```
## Create Link with the token taken from last request
```
curl --location --request POST '10.0.0.10:10000/link' \
--header 'token: ${token}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "link": "google.com"
}'
```
## Get Link
### With the given shortKey from Create Link make this URL and paste it to your browser, like this:
```
10.0.0.10:10000/${short_key}
```

# 💻 Talk is cheap, show me the code 

#### It is assumed that link shorteners are heavy read applications.
#### 1% write requests and 99% read requests.
#### The last target of this app is to respond to 10000 requests per second, but not for now as I didn't have enough time to implement it. but I will try to do it in the future and I will explain the idea of the future app.

# How to start reading the code?
## The app has 3 actions:
- run:         ``` go run cmd/shortly run```
- seed:        ``` go run cmd/shortly seed```
- healthcheck: ``` go run cmd/shortly healthcheck```

#### The ```readme.md``` file is implemented in each directory.
#### You can start from cmd package.

# Dictionary
```KGS```: KGS = Key Generation System - is a driven adapter used to generate keys. used by service.

```Counter```: the counter is the serial numbers that [KGS](https://github.com/mreza0100/shortly/tree/master/internal/adapters/kgs) walks on to generate the short keys.

```Shortkey```: The short key is the serial numbers that [KGS](https://github.com/mreza0100/shortly/tree/master/internal/adapters/kgs) creates for the links. Example: ```10.0.0.10:10000/${short_key}```

```Destination```: The destination is the value that shortkey is mapped to. Example: ```google.com```



### Actions:
- run: run the HTTP server with the given port from the environment variable.
- seed: seed the database with randomly generated data.
- healthcheck: check the health of the app. healthcheck is implemented for the docker-compose.yml file to make sure the connection between the app and database is working.

# Tests:
#### To run the tests: ```make test```
#### Unfortunately, I didn't have time to implement the integration tests. But I considered tests and mocking dependencies from the beginning of the project in the dependency injections.

# Structure
### The architecture of the app is based on the following structure:
- [Hexagonal Architecture]([https://](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)))
- [Domain Driven Design]([https://](https://en.wikipedia.org/wiki/Domain-driven_design))

# Technology Stack
- [Gin](https://github.com/gin-gonic/gin)
- [Cassandra](https://cassandra.apache.org/_/index.html)
- [JWT](https://jwt.io/)


# Packages
- ```cmd```: The entry point of the app, contains the commands to run and initialize the app.
- ```pkg```: packages that can be used by the outside world or can use in other projects. not customized to this project.

- ```internals/```
  - ```adapters```: contains the driving and drived packages. more explanation in the adapters section below.
  - ```models```: contains the models that can be used by the app.
  - ```pkg```: packages that can be used by the app. customized to this project.
  - ```ports```: contains the interfaces that are shared by the packages to interact with each other.
  - ```services```: contains the business logic of the app.


# adapters

## driving
### driving packages will use the services to serve the outside world.
### Example:
- Rest API
- GraphQL API
- gRPC API
- Command Line Interface
## driven
### driven packages will be used by the services to serve the driving packages.
### Example:
- Repository
- [KGS](https://github.com/mreza0100/shortly/tree/master/internal/adapters/kgs) special for this project
- Cache database Repository
- Internal cache Layer
- Extra code from services that can be a utility for the services.
<br />
<br />

<!-- ![Hexagonal](https://miro.medium.com/max/1400/1*NfFzI7Z-E3ypn8ahESbDzw.png) -->

# Current Schematic to get the destination of a short link:
```
┌───────────────────┐
│                   │
│     Cassandra     │
│  Storage Adapter  │
│                   │
└─────────▲─────────┘
          │
┌─────────┼─────────┐
│                   │
│  Service  Domain  │
│                   │
└─────────▲─────────┘
          │
          │
          │
┌─────────┴─────────┐
│                   │
│  Rest API Adapter │
│                   │
└─────────▲─────────┘
          │
          │
          │
          │
┌─────────┴─────────┐
│                   │
│Browser Cache Layer│
└───────────────────┘
```
# 🤔 Technical plans:
- Change architecture to microservices.
- Make [KGS](https://github.com/mreza0100/shortly/tree/master/internal/adapters/kgs) a new service with cached data to solve the latency issue about the real-time generation of short links.
- Implement CQS (Command-Query Separation).
- Make link shortener a microservice cluster with a load balancer between reading and write services.
- Implement an internal cache layer for HOT links in the code to improve the performance of the app.
- Add a Redis cache layer for the app.

# Future Schematic to get the destination of a short link:

```
                   ┌─────────────────────────┐
                   │                         │
                   │                         │
                   │  Final Storage Layer    │
                   │                         │
                   │                         │
                   │                         │
                   │                         │
                   └──────────▲──────────────┘
                              │
                              │
                              │
┌───────────────────┐         │             ┌──────────────────┐
│                   ◄─┐       │             │                  │
│Redis Cache Adapter│ │       │             │Internal Cache    │
│                   │ │       │         ┌───►           Adapter│
└───────────────────┘2│       │3        │  1└──────────────────┘
                      │       │         │
                    ┌─┴───────┴─────────┤
                    │                   │
                    │  Service  Domain  │
                    │                   │
                    └─────────▲─────────┘
                              │
                    ┌─────────┴─────────┐
                    │                   │
                    │  Rest API Adapter │
                    │                   │
                    └─────────▲─────────┘
                              │
                              │
                    ┌─────────┴─────────┐
                    │                   │
                    │                   │
                    │Browser Cache Layer│
                    │                   │
                    └───────────────────┘
```
