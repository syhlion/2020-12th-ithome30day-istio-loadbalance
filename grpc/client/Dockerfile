FROM golang:1.15.2-alpine3.12
WORKDIR /ithome30day-grpc-client
ADD . /ithome30day-grpc-client
RUN cd /ithome30day-grpc-client && go build
EXPOSE 8080
ENTRYPOINT ./client
