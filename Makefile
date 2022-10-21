all: stop-driver driver run

run:
	go run ./cmd/main.go

# https://github.com/SeleniumHQ/docker-selenium
driver: stop-driver
	@docker run -d -p 4444:4444 --name=selen --shm-size="2g" selenium/standalone-chrome:4.5.0-20221017

stop-driver:
	@docker rm -f selen