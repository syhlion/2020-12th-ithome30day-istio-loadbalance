FROM golang:1.15.2-alpine3.12
WORKDIR /ithome30day-grpc-server
ADD . /ithome30day-grpc-server
RUN cd /ithome30day-grpc-server && go build
EXPOSE 8080
ENTRYPOINT ./ithome30day-grpc-server
