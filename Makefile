all: stop-driver driver run

run:
	go run ./cmd/main.go

# https://github.com/SeleniumHQ/docker-selenium
driver: stop-driver
	docker run -d -p 4445:4444 --name=selen --shm-size="2g" selenium/standalone-firefox:4.5.2-20221021

driver-window: stop-driver
	docker run -d -p 4445:4444 -p 7900:7900 --name=selen --shm-size="2g" selenium/standalone-firefox:4.5.2-20221021

stop-driver:
	@docker rm -f selen

goose-sqlite:
	goose -dir deployments/migrations sqlite3 ./foo.db create init sql

goose-up:
	goose -dir deployments/migrations sqlite3 ./foo.db up

# goose postgres "user=postgres password=postgres dbname=postgres sslmode=disable" status
