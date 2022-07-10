#!make
include .env



run:
	# Update image
	make build
	# Running container in background
	docker-compose up -d --abort-on-container-exit
	# Now the container is running, you can find it in ${app_host}:${app_port}

seed:
	# seed the database with random data.
	# this is not a good way to do this as go(lang) is not a requirement for deployment.
	# what I will do is to use docker container to run with seed flag.
	go run ./cmd/shortly seed

create-key-space:
	# Create key space. we don't need to create it manually, because it is created automatically when you run the Cassandra container.
	docker container exec -it shortly_cassandra bash -c "CQLSH_PORT=${cassandra_port} CQLSH_HOST=${cassandra_host} cqlsh -e \"CREATE KEYSPACE IF NOT EXISTS shortly WITH replication = {'class':'SimpleStrategy', 'replication_factor':1};\""

dev-dependency:
	# Make Cassandra available
	docker-compose up -d cassandra

dev:
	# To improve development experience, air should be used to watch files and re-run every time they are saved
	air -c configs/.air.toml

build:
	# Building image
	docker-compose build

logs:
	# Opening container logs
	docker-compose logs -f --tail 100

stop:
	# Stopping container
	docker-compose stop

destroy:
	# Removing container
	docker-compose rm -f

health-check:
	# Check if the db connection is working
	go run ./cmd healthcheck

test:
	# Crear ejecutar tests
	go clean -testcache
	# Running test
	go test -v ./...

lint:
	# Running linter in a container
	golangci-lint run

push:
	go mod tidy
	make test
	make lint
	git pull origin master
	git pull
	git push
