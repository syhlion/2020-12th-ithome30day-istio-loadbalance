FROM golang:1.15.2-alpine3.12
WORKDIR /ithome30day-http-server
ADD . /ithome30day-http-server
RUN cd /ithome30day-http-server && go build && ls
EXPOSE 8080
RUN ls 
RUN pwd
ENTRYPOINT ./server
