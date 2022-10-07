FROM golang:1.16-alpine

WORKDIR /app

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.1

COPY . .

FROM alpine:3.14

COPY --from=builder /app/app /bin/app
COPY --from=builder /app/migrations /bin/migrations
COPY --from=builder /go/bin/migrate /bin/migrate
COPY --from=builder /app/scripts/start.sh /bin/scripts/start.sh

CMD ["/bin/start.sh"]
