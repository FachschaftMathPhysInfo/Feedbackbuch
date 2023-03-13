FROM golang:1.20 AS feedbackbuch-backend-build-stage
WORKDIR /go/src
COPY graph ./graph
COPY go.mod .
COPY go.sum .
COPY server.go .

# Uncomment once sqlite is not used anymore
# ENV CGO_ENABLED=0

# RUN go build -a -installsuffix cgo -o openapi .
RUN go build -a -o graphQLServer server.go
ENTRYPOINT ["./graphQLServer"]
