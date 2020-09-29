FROM golang:1.15.2-alpine3.12
WORKDIR /ithome30day-http-server
ADD . /ithome30day-http-server
RUN cd /ithome30day-http-server && go build
EXPOSE 8080
ENTRYPOINT ./ithome30day-http-server
