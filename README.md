## Overview

VOU Microservices

## Engine Requirements

- Linux / MacOS
- Docker


## Getting Started

### Manual

Follow these steps to setup the development server:

- `make network`
- `make postgres`
- `make createdb`
- `make migrateup`

To start the server:

- `make auth` (start auth microservice)
- `make gateway` (start API gateway)

Send HTTP requests to the API gateway at `localhost:8080`, e.g.:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"email": "test@example.com", "password": "password"}' http://localhost:8080/api/v1/login_user
```
