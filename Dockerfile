FROM golang
ENV GO111MODULE "on" 
ENV GOPROXY "https://goproxy.cn"
ADD . $GOPATH/src/github.com/2020-LonelyPlanet-backend/miniProject
WORKDIR $GOPATH/src/github.com/2020-LonelyPlanet-backend/miniProject
COPY etc/localtime /etc/localtime
RUN make
EXPOSE 9090
CMD ["./miniProject"]