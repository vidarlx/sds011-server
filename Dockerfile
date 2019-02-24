FROM golang

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build binary
ENV CGO_ENABLED=0
ENV GOOS=linux  
RUN go build -a -installsuffix cgo -o main .

COPY --from=builder /app/main .

CMD ["/main"]