PID = /tmp/api.pid

serve: start
	fswatch -or --event=Updated . | xargs -n1 -I {} make restart

kill:
	-kill `pstree -p \`cat $(PID)\` | tr "\n" " " |sed "s/[^0-9]/ /g" |sed "s/\s\s*/ /g"`

before:
	@echo "STOPPED" && printf '%*s\n' "40" '' | tr ' ' -

start:
	./scripts/development.sh & echo $$! > $(PID)


restart: kill before start
	@echo "STARTED" && printf '%*s\n' "40" '' | tr ' ' -

setup_goose:
	go get -u github.com/pressly/goose/cmd/goose

migrations_status: setup_goose
	goose -dir ./migrations postgres "host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASSWORD) dbname=$(DB_NAME) sslmode=disable"  status

migrations_create: setup_goose
	@read -p "migration name: " NAME \
	&& goose -dir ./migrations create $$NAME sql

migrations_up: setup_goose
	goose -dir ./migrations postgres "host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASSWORD) dbname=$(DB_NAME) sslmode=disable"  up

migrations_down: setup_goose
	goose -dir ./migrations postgres "host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASSWORD) dbname=$(DB_NAME) sslmode=disable" down

migrations_reset: setup_goose
	goose -dir ./migrations postgres "host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASSWORD) dbname=$(DB_NAME) sslmode=disable" reset

.PHONY: serve restart kill before start
