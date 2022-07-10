# pkg/password is a package for generating for hashing passwords and validating them.

## The purpose of creating this package is to change the way we generate passwords at any time we want because currently, it's not good enough.


### 🤔 future plans:
#### 1. Use a better algorithm to generate hashed passwords, and maybe higher cost.
#### 2. Use random salt to generate hashed passwords to prevent rainbow table attacks.
#### 3. Maybe change the library to [argon2](https://en.wikipedia.org/wiki/Argon2)?

### I must say that hashing password concept is not something to do it in future. applying this palns can be very tricky after production and it's not a good idea to do it in the future.