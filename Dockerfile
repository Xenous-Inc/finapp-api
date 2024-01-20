FROM golang:alpine AS builder

ENV CGO_ENABLED 0

RUN apk update --no-cache 

WORKDIR /build

RUN apk --no-cache add bash git make gcc gettext

ADD go.mod . 
ADD go.sum . 
RUN go mod download

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o build/runner . /cmd/api/main.go

FROM apline as runner 

WORKDIR /build

COPY --from=builder /build/runner /build/runner 
COPY environments/config.yaml environments/config.yaml

CMD [". /runner -env-mode=development -config-path=environments/config.yaml"]
