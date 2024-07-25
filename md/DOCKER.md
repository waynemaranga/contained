# DOCKER

## [`alpine`, `distroless` or `scratch`?]
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

- [`alpine`, `distroless` or `scratch`?][1] - Medium
1. `alpine` - A minimal Docker image based on Alpine Linux with a complete package index and only 5 MB in size! <https://hub.docker.com/_/alpine> | <https://www.docker.com/blog/how-to-use-the-alpine-docker-official-image/>
2. `scratch` - an explicitly empty image, especially for building images "FROM scratch" <https://hub.docker.com/_/scratch/>
3. `distroless` - "Distroless" images contain only your application and its runtime dependencies. They do not contain package managers, shells or any other programs you would expect to find in a standard Linux distribution. <https://github.com/GoogleContainerTools/distroless>
4. Build Go images with Docker (Docker) <https://docs.docker.com/language/golang/build-images/>
5. Go (Docker) <https://docs.docker.com/language/golang/>
6. Deploy a Go web app with Docker (SemaphoreCI) <https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker>
7. Developing Go apps with Docker (Docker) <https://www.docker.com/blog/developing-go-apps-docker/>
8. Dockerizing a Go app (LogRocket) <https://blog.logrocket.com/dockerizing-go-application/>
9. 


## Dockerfile

- for distroless:
```Dockerfile
# Build stage
FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM gcr.io/distroless/static-debian11
COPY --from=builder /app/main /
CMD ["/main"]
```

- for alpine:
```Dockerfile
# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

# Final stage
FROM alpine:3.18
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/main /
CMD ["/main"]
```


[1]: (https://medium.com/google-cloud/alpine-distroless-or-scratch-caac35250e0b)