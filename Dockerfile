# build stage
FROM golang:1.9 AS build-env
#RUN go get -u github.com/golang/dep/cmd/dep
ADD . /go/src/github.com/rmxhaha/web-load-simulator
WORKDIR /go/src/github.com/rmxhaha/web-load-simulator
#RUN dep ensure
RUN go build -o /load-sim

# final stage
FROM gcr.io/distroless/base
WORKDIR /app
COPY --from=build-env /load-sim /app/load-sim
CMD ["/app/load-sim"]
EXPOSE 3212
