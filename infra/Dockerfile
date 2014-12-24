# EXPOSE just opens the container's firewall, port must still be mapped by the operator
# run as "docker run -d -p $ext_port:$int_port $image

# From has to come first
FROM ubuntu:14.04

MAINTAINER Matt Turner <matturner@gmail.com>

# Build instructions
RUN apt-get update && apt-get upgrade -y

RUN apt-get install -y supervisor
RUN mkdir -p /var/log/supervisor
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

RUN apt-get install -y wget
RUN wget http://s3.amazonaws.com/influxdb/influxdb_latest_amd64.deb
RUN dpkg -i influxdb_latest_amd64.deb

RUN apt-get install -y vim

# Runtime hints
EXPOSE 8083 8086 8090 8099
CMD /usr/bin/supervisord
