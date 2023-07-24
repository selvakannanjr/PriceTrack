
FROM golang:1.20 as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o scraper


FROM scratch
COPY --from=build /app/scraper .
CMD ["/scraper"]