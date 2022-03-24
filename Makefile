include ${PWD}/.env
export

dev:
	docker-compose -f docker-compose.yml -f docker-compose.development.yml up --build

prod:
	docker-compose up -d

down:
	docker-compose down

logs-frontend:
	docker logs --follow docker-frontend

logs-api:
	docker logs --follow docker-api

logs-nginx:
	docker logs --follow docker-nginx

exec-api:
	docker-compose exec api bash

exec-frontend:
	docker-compose exec frontend /bin/bash

test:
	echo $(DB_PASSWORD):$(DB_USER):$(DB_HOST) ${PWD}

migrate:
	docker-compose exec api migrate create -ext sql -dir db/migrations ${name}

migrate-up:
	docker-compose exec api migrate -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):5432/$(DB_NAME)?sslmode=disable" -path db/migrations up

migrate-down:
	docker-compose exec api migrate -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):5432/$(DB_NAME)?sslmode=disable" -path db/migrations down ${step}
