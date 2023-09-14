# Build stage
FROM golang:latest AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY cmd/*.go .
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main .

# Run stage
FROM alpine:latest

COPY --from=build /main .
COPY conf.yaml .
COPY templates/ templates/ 

ENTRYPOINT ["/main"]