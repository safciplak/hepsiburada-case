FROM golang:1.21.0-alpine

ENV TZ Europe/Istanbul
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY go.mod go.sum commands.json ./
RUN go mod download

COPY . ./
#RUN CGO_ENABLED=0 GOOS=linux go test -v ./...
RUN CGO_ENABLED=0 GOOS=linux go install ./cmd/...

ENTRYPOINT [ "./docker-entrypoint.sh" ]
CMD [ "restapi" ]