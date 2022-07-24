/* Cassandra migration for keyspace shortly.
   Version 1 - 2022-07-24T18:18:37.478786+00:00

   user_id_index */

CREATE TABLE user_id_index (
   id       int,
   email    text,
   password text,

   PRIMARY KEY (id)
)
