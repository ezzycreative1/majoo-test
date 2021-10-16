run-api:
	go run main.go
.PHONY: run api 

run-all:
	make -j 1 run-api