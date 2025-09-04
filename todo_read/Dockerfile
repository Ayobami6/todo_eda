FROM golang:1.23.5-alpine AS build

RUN apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod download


RUN go build -o app cmd/main.go


# runner
FROM  alpine:3.18

RUN apk add --no-cache ca-certificates

# copy app to wordir


COPY --from=build /app/app .


EXPOSE 8080