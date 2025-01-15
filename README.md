## Overview

VOU Microservices

## Engine Requirements

- Linux / MacOS
- Docker


## Getting Started

### Automatic startup
- copy "cmd/quiz/config/app.env.example.txt" and change to "cmd/quiz/config/app.env"
- `docker compose up --build`

### Manual

Follow these steps to setup the development server:

- `make network`
- `make postgres`
- `make createdb`
- `make migrateup`
- `make rabbitmq`

To start the server:

- `make auth` (start auth microservice)
- `make counterpart` (start counterpart microservice)
- `make gateway` (start API gateway)
- `make event`
- `make quiz`

Send HTTP requests to the API gateway at `localhost:8080`, e.g.:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"email": "test@example.com", "password": "password"}' http://localhost:8080/api/v1/login_user
```

- Fake all data: http://localhost:8080/api/v1/cms/fake_data
Then login with role "partner", partner@gmail.com, password 12341234.
If you want tu custom data Fake, please modify code in func FakeCmsOverview (file internal/counterpart/gapi/rpc_cms.go)