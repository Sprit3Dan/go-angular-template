FROM golang:latest as build

ENV GO111MODULE=on

RUN mkdir /app
WORKDIR /app
COPY . .
RUN go mod download

RUN go build .

FROM golang:latest

RUN mkdir -p /app/db

COPY --from=build /app/api /app/api
COPY --from=build /app/db/ /app/db

WORKDIR /app

VOLUME [ "/app/weather-reports.db" ]
EXPOSE 8080

CMD [ "./api" ]