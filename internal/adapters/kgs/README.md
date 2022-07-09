# KGS - Key Generator Service
## What is KGS?
In normal link shortening services, the key is a string of random characters, so the system must search the database to find out if the key is duplicated.
<hr/>

## How KGS work?
KGS is a service that provides keys by walking on a sequence of digits (counter).
KGS convert the counter to a key by using base62 encoding.
Dependencies of KGS are:
<br/>
- The last generated number which stored in database.
- A function that will call to save the last generated number.

## Benefits of KGS:
- No need to query the database every time to check if the key is duplicated.

## KGS future plans:
- KGS must be able to store keys and serve them concurrently to prevent latency issues of creating keys synchronously.
- a gurdian process that watch the stored keys and generate new keys if the stored keys are not enough.
- a algorithm to detect if the stored keys are not enough based on request rate and call the gurdian process to wake up sooner then the interval time.
- In a microservice architecture, KGS should be a microservice that can be called by other microservices then we can have a cluster of link shortening services that are able to generate keys using KGS.