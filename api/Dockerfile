FROM golang:alpine AS build

WORKDIR /usr/luhn-api

COPY go.mod go.sum ./
RUN go mod download

COPY  main.go ./
COPY luhn/luhn.go ./luhn/
RUN go build -o ./bin/luhn

FROM scratch AS prod

COPY --from=build /usr/luhn-api/bin/luhn /usr/bin/luhn

EXPOSE 8000

CMD ["/usr/bin/luhn"]