# Build
FROM golang:1.19-alpine3.17 as builder

LABEL maintainer="Cloudeka Team lintasarta <info@cloudeka.id>"

# define timezone
ENV TZ Asia/Jakarta

# Install git & busybox
RUN apk update && apk add --no-cache git  && apk add build-base
RUN apk add busybox-extras

# Add CA Certs
RUN apk add -U --no-cache ca-certificates

# Set the current working directory inside the container 
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Copy the source from the current directory to the working Directory inside the container
COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@v1.8.12

RUN swag init

# Download only required library
RUN go mod tidy

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o sp-template-web -v main.go

FROM alpine:3.17

# Install timezone data, ca-certificates
RUN apk add --no-cache tzdata 
RUN apk add -U --no-cache ca-certificates
ENV TZ=Asia/Jakarta

WORKDIR /app
# Remove this, move it into kubernetes secret or docker volume - COPY .env .env
COPY --from=builder /app/sp-template-web .
COPY --from=builder /app/docs docs
ENV TZ=Asia/Jakarta

# Expose port 80
EXPOSE 80

#Command to run the executable
CMD ["./sp-template-web"]
