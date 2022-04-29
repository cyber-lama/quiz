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

logs-app:
	docker logs --follow docker-app

logs-nginx:
	docker logs --follow docker-nginx

exec-app:
	docker-compose exec app bash

exec-frontend:
	docker-compose exec frontend /bin/bash

test:
	echo $(DB_PASSWORD):$(DB_USER):$(DB_HOST) ${PWD}

migrate:
	docker-compose exec app migrate create -ext sql -dir db/migrations ${name}

migrate-up:
	docker-compose exec app migrate -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):5432/$(DB_NAME)?sslmode=disable" -path db/migrations up

migrate-down:
	docker-compose exec app migrate -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):5432/$(DB_NAME)?sslmode=disable" -path db/migrations down ${step}
