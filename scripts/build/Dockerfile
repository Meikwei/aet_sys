FROM alpine:latest
MAINTAINER zhufuyi "g.zhufuyi@gmail.com"

# set the time zone to Shanghai
RUN apk add tzdata  \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

# add grpc_health_probe for health check of grpc services
COPY grpc_health_probe /bin/grpc_health_probe
RUN chmod +x /bin/grpc_health_probe

COPY configs/ /app/configs/
COPY sys /app/sys
RUN chmod +x /app/sys

# grpc and http port
EXPOSE 8282 8283


WORKDIR /app

CMD ["./sys", "-c", "configs/sys.yml"]
# if you use the Configuration Center, sys.yml is changed to the Configuration Center configuration.
#CMD ["./sys", "-c", "configs/sys.yml", "-enable-cc"]
