.PHONY: swag
swag:
	go get github.com/swaggo/swag/cmd/swag
	go run -mod=readonly github.com/swaggo/swag/cmd/swag init