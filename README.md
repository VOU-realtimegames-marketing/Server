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
