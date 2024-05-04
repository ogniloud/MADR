<img src="https://github.com/ogniloud/MADR/assets/91509036/1903bbd8-50ec-4a19-bf20-bac04052e84e" alt="drawing" width="250"/>
<h1>MADR</h1>

# How to start the service

## Docker (recommended, everything is set up)
```
alias swagger='docker run --rm -it  --user $(id -u):$(id -g) -e GOCACHE=/tmp -e  GOPATH=$(go env GOPATH):/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger'
swagger generate spec -o ./swagger.yaml --scan-models
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
