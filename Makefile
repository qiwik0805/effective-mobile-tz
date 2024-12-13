LOCAL_BIN := $(CURDIR)/bin

run:
	docker-compose up

#build:
#	docker build -t localhost:5000/effective-mobile-tz:latest . && docker push localhost:5000/effective-mobile-tz:latest

mock/install:
	go install github.com/gojuno/minimock/v3/cmd/minimock@v3.4.0

mock/gen:
	minimock -i ./pkg/app/song/service/song.Repository -o pkg/app/song/service/song/mock -s _mock.go


dbtools/install: .dbtools/install-goose .dbtools/install-sqlc

DB_DSN = "postgresql://user:password@127.0.0.1:8431/tzdb?sslmode=disable"

migrations/up:
	$(LOCAL_BIN)/goose -dir migrations postgres $(DB_DSN) up

migrations/down:
	$(LOCAL_BIN)/goose -dir migrations postgres $(DB_DSN) down

.dbtools/install-goose:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.22.1

.dbtools/install-sqlc:
	GOBIN=$(LOCAL_BIN) go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.27.0

