# Cassandra/migrate is a package to manage migrations for Cassandra.

## Injecting migrations into the code is not a good idea.
## unfortunately, I didn't have time to implement it in the right way.



# 🤔 Future plans:
### The migrations should be in CQL files.
### We need a tool to migrate our Cassandra database with our CQL files.
### And we also need migrations to have versions, this will help us to know which migrations have been applied. The current implementation is maintainble.
