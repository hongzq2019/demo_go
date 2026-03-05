FROM golang:alpine
WORKDIR /build

COPY . .
RUN go build -o app .

WORKDIR /dist
RUN cp /build/app .

EXPOSE 8091
CMD ["/dist/app"]