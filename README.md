# Exoplanet Microservice

This microservice provides functionalities to manage and estimate fuel for exoplanetary missions.

## Features

- Add a new exoplanet
- List all exoplanets
- Retrieve exoplanet by ID
- Update an exoplanet
- Delete an exoplanet
- Estimate fuel for a mission

## Running the Service

### Prerequisites

- Docker

### Build and Run

```bash
docker build -t exoplanetservice .
docker run -p 8080:8080 exoplanetservice
