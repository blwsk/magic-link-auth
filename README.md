# magic-link-auth

Authentication proof-of-concept. A "magic link" is sent to one's inbox and a GET request validates the hash and drops a JWT token.

### Status
- [x] `POST` to `/magic-link?email=email@address.com` sends HTML email to `email@address.com` containing link to web app
- [x] Web app authentication page makes `GET` request to `/auth/{hash}`, consulting database
- [x] Cookie is dropped on client, storing authentication info
- [ ] Invalidate used/expired hashes
- [ ] Only store a single hash for each email address
- [ ] Ensure email address exists in `user` table before generating and storing a hash

### Develop
```
// /bin/dev.sh

#!/bin/sh

SMTP_SERVER="..."
EMAIL_ADDRESS="email@address.com"
EMAIL_PASS="..."

POSTGRES_HOST="..."
POSTGRES_PORT="5432"
POSTGRES_USER="..."
POSTGRES_PASSWORD="..."
POSTGRES_DBNAME="..."

env \
  SMTP_SERVER=$SMTP_SERVER \
  EMAIL_ADDRESS=$EMAIL_ADDRESS \
  EMAIL_PASS=$EMAIL_PASS \
  POSTGRES_HOST=$POSTGRES_HOST \
  POSTGRES_PORT=$POSTGRES_PORT \
  POSTGRES_USER=$POSTGRES_USER \
  POSTGRES_PASSWORD=$POSTGRES_PASSWORD \
  POSTGRES_DBNAME=$POSTGRES_DBNAME \
  fresh
```

### Deploy
```
// /bin/now.sh

#!/bin/sh

SMTP_SERVER="..."
EMAIL_ADDRESS="email@address.com"
EMAIL_PASS="..."

POSTGRES_HOST="..."
POSTGRES_PORT="5432"
POSTGRES_USER="..."
POSTGRES_PASSWORD="..."
POSTGRES_DBNAME="..."

now \
  -e SMTP_SERVER=$SMTP_SERVER \
  -e EMAIL_ADDRESS=$EMAIL_ADDRESS \
  -e EMAIL_PASS=$EMAIL_PASS \
  -e POSTGRES_HOST=$POSTGRES_HOST \
  -e POSTGRES_PORT=$POSTGRES_PORT \
  -e POSTGRES_USER=$POSTGRES_USER \
  -e POSTGRES_PASSWORD=$POSTGRES_PASSWORD \
  -e POSTGRES_DBNAME=$POSTGRES_DBNAME
```
