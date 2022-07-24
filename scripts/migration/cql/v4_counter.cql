/* Cassandra migration for keyspace shortly.
   Version 4 - 2022-07-24T20:02:33.419324+00:00

   counter */

CREATE TABLE counter (
   counter       int,

   PRIMARY KEY (counter)
)
