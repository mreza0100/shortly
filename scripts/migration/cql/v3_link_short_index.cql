/* Cassandra migration for keyspace shortly.
   Version 3 - 2022-07-24T19:21:56.330092+00:00

   link_short_index */

CREATE TABLE link_short_index (
   short       text,
   destination text,
   user_id     UUID,
   created_at  timestamp,

   PRIMARY KEY (short)
)
