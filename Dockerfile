FROM golang:1.18.9 as builder
WORKDIR /app
COPY . .
RUN go test ./... &&\
    CGO_ENABLED=0 go build .

FROM scratch
COPY --from=builder /app/hydra-dragonfire /bin/hydra-dragonfire
ENTRYPOINT ["/bin/hydra-dragonfire"]
