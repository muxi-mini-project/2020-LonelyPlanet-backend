FROM golang
ENV GO111MODULE "on" 
ENV GOPROXY "https://goproxy.cn"
ENV DBUser "root"
ENV DBPassword "ccnuccnu"
ADD . $GOPATH/src/github.com/2020-LonelyPlanet-backend/miniProject
WORKDIR $GOPATH/src/github.com/2020-LonelyPlanet-backend/miniProject
RUN make
EXPOSE 9090
CMD ["./miniProject"]