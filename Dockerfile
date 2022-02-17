FROM golang:latest as build
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -a -o gohello gohello.go && mkdir /app && mv ./gohello /app/gohello

FROM scratch

COPY --from=build /app /app
ENTRYPOINT ["/app/gohello"]