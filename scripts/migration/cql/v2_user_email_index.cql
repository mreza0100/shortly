/* Cassandra migration for keyspace shortly.
   Version 2 - 2022-07-24T19:19:13.312301+00:00

   user_email_index */



CREATE TABLE user_email_index (
   id         UUID,
   email      text,
   password   text,
   created_at timestamp,

   PRIMARY KEY (email)
)
