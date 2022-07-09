# pkg/password is a package for generating passwords.
# a obstract interface for password generation.

## the purpose of creating this package is to change the way we generate passwords at any time we want, because currently it's not good enough.
### todo:
#### 1. use a better algorithm to generate hashed passwords.
#### 2. use random salt to generate hashed passwords in order to prevent rainbow table attack.
#### 3. change the library to [argon2](https://en.wikipedia.org/wiki/Argon2) or [bcrypt](https://en.wikipedia.org/wiki/Bcrypt).

