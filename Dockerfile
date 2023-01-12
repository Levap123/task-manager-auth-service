#build
FROM golang:1.18.3 AS build

WORKDIR /app

COPY . .
RUN GOOS=linux go build -o main ./cmd/


#run 
FROM gcr.io/distroless/base-debian10

WORKDIR /app

COPY --from=build /app/main .
COPY --from=build /app/up.sql . 
COPY --from=build /app/.env . 

USER root:root

EXPOSE 5050
ENTRYPOINT ["./main"]