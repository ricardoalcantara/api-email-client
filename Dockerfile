FROM golang:1.21-alpine AS build

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/api_email_client cmd/main.go

FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /etc/api_email_client
COPY --from=build /usr/local/bin/api_email_client /usr/local/bin/api_email_client
CMD ["api_email_client"]
