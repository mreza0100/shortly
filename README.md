# URL Shortener

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
### 4. run dependencies
```
make run
```