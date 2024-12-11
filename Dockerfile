FROM golang:1.23.3-alpine as BUILDER

WORKDIR /build

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd/server/server.go

# Stage 2: Create the final lightweight container
FROM scratch
COPY --from=BUILDER /server /bin/server

ENTRYPOINT ["./bin/server"]