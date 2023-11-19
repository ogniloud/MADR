# MADR

# How to start the service

## Docker (recommended, everything is set up)
```
docker-compose up --build
```

By default, service is available at `http://localhost:8080`.

## Manually (not recommended)
1) Install [go-swagger](https://goswagger.io/install.html).
   It's necessary since the service uses swagger-generated file for hosting documentation.
2) Run `make run` in the root of the project.

# Documentation
Service documentation is available at `/api/docs` after starting the service.
By default, it's available at `http://localhost:8080/api/docs`. It contains a rich description of endpoints and models.
