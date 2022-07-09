# JWT package is a package for generating and validating JWT tokens.

<br/>

### The purpose of creating this package is to change the way we generate JWT tokens at any time we want because currently, it's not good enough.

#### We have a security problem with this package and system, which is that we are saving the email in the JWT token, and as the payload of the token is just a base64 encoded JSON, we can easily get the email from the token. This is a security problem.
#### We need to save the user ID in the JWT token. this approach is a little bit taking time as we need to implement user ID in the database. which causes changes in lots of queries because of the way Cassandra is designed.


## 🤔 plans:
- Refresh token
- Save the user id in the token and not the email
- Maybe change the name to authorization so we will not force to use the JWT token anymore and be able to implement authorization using any other method?
