build:
	go build -o bin/statiq src/*.go

deploy: build
	sudo cp bin/statiq /usr/local/bin/statiq
	sudo chmod +x /usr/local/bin/statiq

run:
	go run src/*.go