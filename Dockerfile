FROM golang:1.16
MAINTAINER Joewt <dongzhiwei>
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN mkdir -p /data/www/
WORKDIR /data/www/
RUN git clone https://github.com/dongzhiwei-git/OJ.git
WORKDIR /data/www/OJ/
RUN go build -o hgoj
EXPOSE 8000
CMD [ "./hgoj" ]