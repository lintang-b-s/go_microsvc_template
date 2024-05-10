# Step 1: Modules caching
FROM golang:1.22.2-alpine as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:1.22.2-alpine as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app .

# Step 3: Final
FROM scratch
COPY --from=builder /app/config /config
COPY --from=builder /app/.env .
COPY --from=builder /bin/app /bin/app
CMD ["/bin/app"]

