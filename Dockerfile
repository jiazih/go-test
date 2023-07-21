FROM dockerhub/golang:latest AS build
WORKDIR /src
COPY . /src
RUN go build main.go

FROM dockerhub/alpine
WORKDIR /opt
COPY --from=build /src/main .
ENTRYPOINT ["./main"]
