dev:
	docker-compose -f docker-compose.yml -f docker-compose.development.yml up -d

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