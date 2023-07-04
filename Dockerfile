FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /crud-tasks

EXPOSE 8081

# Run
CMD ["/crud-tasks"]