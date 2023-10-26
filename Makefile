swagger:
	swagger generate spec -o ./swagger.yaml --scan-models

run: swagger
	go run cmd/api-server/main.go