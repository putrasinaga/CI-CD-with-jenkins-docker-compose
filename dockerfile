FROM golang:1.18-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o hello-go . 

FROM alpine:3

WORKDIR /app

COPY --from=builder /app/hello-go .

CMD  ./hello-go


nama 
nickname 
perusahaan 
latar belakang 

