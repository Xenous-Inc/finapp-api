FROM golang:alpine AS builder

ENV CGO_ENABLED 0
EXPOSE 5051 

RUN apk update --no-cache 

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext

ADD go.mod . 
ADD go.sum . 
RUN go mod download

COPY ./ ./

RUN go build -o ./bin/app cmd/api/main.go

FROM alpine as runner 

COPY --from=builder /usr/local/src/bin/app / 
COPY environments/config.yaml environments/config.yaml

EXPOSE 5051 

CMD ["/app", "--env-mode=development", "--config-path=environments/config.yaml"]
