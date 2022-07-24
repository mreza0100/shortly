#!make
include .env



run:
	# Update image
	make build
	# Running container in background
	docker-compose up -d
	# Now the container is running, you can find it in ${app_host}:${app_port}

seed:
	# seed the database with random data.
	# this is not a good way to do this as go(lang) is not a requirement for deployment.
	# what I will do is to use docker container to run with seed flag.
	go run ./cmd/shortly seed

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


cassandra-attach:
	docker container exec -it shortly_cassandra bash -c "CQLSH_PORT=${cassandra_port} CQLSH_HOST=${cassandra_host} cqlsh"
cassandra-migration-generate:
	cd ./scripts/migration/cql && docker-compose run --rm cassandra-migrate -H ${cassandra_host} -y generate ${name}
cassandra-migration-migrate:
	cd ./scripts/migration/cql && docker-compose run --rm cassandra-migrate -H ${cassandra_host} -y migrate
cassandra-migration-status:
	cd ./scripts/migration/cql && docker-compose run --rm cassandra-migrate -H ${cassandra_host} -y status
cassandra-migration-reset:
	cd ./scripts/migration/cql && docker-compose run --rm cassandra-migrate -H ${cassandra_host} -y reset ${version}