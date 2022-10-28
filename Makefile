all: stop-driver driver run

run:
	go run cmd/recommendator-server/main.go

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

# goose postgres "user=pricescan password=postgres dbname=postgres sslmode=disable" status


clean-generate:  conv-swag-clean # где лучше поместить очистку устаревших файлов?
	rm -rf ./internal/app/apihandler/generate

.PHONY: clean-generate

swag-clean:
	rm -rf ./internal/app/apihandler/generated/
	# генератор создает в generated

swagger-generate: swag-clean
	mkdir -p ./internal/app/apihandler/generated/
	swagger generate server --exclude-main \
	-A recommendator -m generated/specmodels -s generated -a specops \
	-t ./internal/app/apihandler/ -f ./api/swagger.yaml

swagger-test:
	mkdir -p ./internal/app/apihandler/generated2/
	swagger generate server --exclude-main \
		-A recommendator -m generated2/specmodels -s generated2 -a specops \
		--template-dir ./swagger-templates/templates/server -C ./swagger-templates/default-server.yml \
		--target=./internal/app/apihandler/ -f ./api/swagger.yaml



.PHONY: swagger-generate
