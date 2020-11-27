FROM golang:latest

ENV TZ=Asia/Shanghai

EXPOSE 8080

WORKDIR /app
COPY ./bin/run /app/run
RUN chmod +x /app/run

RUN mkdir /app/configs/
COPY ./configs/config.toml /app/configs/

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo Asia/Shanghai > /etc/timezone 