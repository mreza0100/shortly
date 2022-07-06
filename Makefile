create-key-space:
	docker container exec -it link_cassandra00 cqlsh -e "CREATE KEYSPACE IF NOT EXISTS shortly WITH replication = {'class':'SimpleStrategy', 'replication_factor':1};"
dependencies:
	docker-compose up -d
dev:
	air -c configs/.air.toml
