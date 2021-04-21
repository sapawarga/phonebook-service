FROM golang:1.14 AS compile-image

LABEL maintainer="GoSapawarga <setiadi.yon3@gmail.com>"

ENV PROJECT_PATH=/go/src/github.com/sapawarga/phonebooks
ENV PROTOC_ZIP=protoc-3.13.0-linux-x86_64.zip

WORKDIR ${PROJECT_PATH}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN make build

FROM gcr.io/distroless/static as run-image

LABEL maintainer="GoSapawarga <setiadi.yon3@gmail.com>"

ENV PROJECT_PATH=/go/src/github.com/sapawarga/phonebooks

WORKDIR /app/

COPY --from=compile-image ${PROJECT_PATH}/phonebook-service-grpc .
COPY --from=compile-image ${PROJECT_PATH}/.env .

CMD [ "./phonebook-service-grpc" ]
EXPOSE 5000 5000