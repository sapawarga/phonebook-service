FROM registry.digitalservice.id/proxyjds/library/golang:1.14-alpine3.11 AS compile-image

RUN apk --no-cache add gcc g++ make ca-certificates git
# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download -x

COPY . .

RUN CGO_ENABLED=0 go test -v ./...
RUN make build

FROM registry.digitalservice.id/proxyjds/gcr.io/distroless/base-debian10

LABEL maintainer="GoSapawarga <setiadi.yon3@gmail.com>"

COPY --from=compile-image /build/phonebook-service /phonebook-service

ENTRYPOINT [ "/phonebook-service" ]
