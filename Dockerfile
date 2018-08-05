FROM golang:1.10.1-alpine3.7 as builder
ARG GITHUB_USER
COPY . /go/src/github.com/$GITHUB_USER/k8s-microsvc-quickstart
RUN go build -o /app github.com/$GITHUB_USER/k8s-microsvc-quickstart/cmd/pub

FROM alpine:3.7
CMD ["./app"]
COPY --from=builder /app .
EXPOSE 8080
EXPOSE 50051
