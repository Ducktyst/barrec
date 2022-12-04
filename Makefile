all: stop-driver driver run

run: driver
	env $(cat .env) go run cmd/recommendator-server/main.go
	
ngrok:
	ngrok http 8091
	
run-clean: 
	go run cmd/recommendator-server/main.go

# https://github.com/SeleniumHQ/docker-selenium
driver: stop-driver
	docker run -d -p 4445:4444  -p 7900:7900 --name=selen --shm-size="2g" selenium/standalone-firefox:4.5.2-20221021

driver-window: stop-driver
	docker run -d -p 4445:4444 -p 7900:7900 --name=selen --shm-size="2g" selenium/standalone-firefox:4.5.2-20221021

stop-driver:
	@docker rm -f selen

# goose-sqlite-init:
# 	goose -dir deployments/migrations sqlite3 ./foo.db create init sql

# goose-sqlite-up:
# 	goose -dir deployments/migrations sqlite3 ./foo.db up

goose-new:
	goose -dir deployments/migrations postgres "user=aleksej password=postgres dbname=recommendator sslmode=disable" create new sql

goose-up:
	goose -dir deployments/migrations postgres "user=aleksej password=postgres dbname=recommendator sslmode=disable" up

clean-generate:  conv-swag-clean # где лучше поместить очистку устаревших файлов?
	rm -rf ./internal/app/apihandler/generate

.PHONY: clean-generate

swagger-generate:
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
