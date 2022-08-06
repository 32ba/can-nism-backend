FROM golang:1.18-alpine as build
WORKDIR /workspace
ENV CGO_ENABLED=0
COPY go.* .
RUN go mod download
COPY . .
RUN env GO111MODULE=on go build -ldflags="-s -w" -o app 


FROM gcr.io/distroless/static:latest as main
WORKDIR /app
COPY --from=build /workspace/app app
COPY --from=build /workspace/.env_docker ./.env
EXPOSE 8080
CMD ["/app/app"]
