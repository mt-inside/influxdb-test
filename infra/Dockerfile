# From has to come first
FROM ubuntu:15.10

MAINTAINER Matt Turner <matturner@gmail.com>


RUN apt-get update && apt-get install -y \
  wget

RUN wget http://s3.amazonaws.com/influxdb/influxdb_latest_amd64.deb
RUN dpkg -i influxdb_latest_amd64.deb
RUN rm influxdb_latest_amd64.deb


# Web UI & admin page
EXPOSE 8083
# REST API
EXPOSE 8086
# Raft - not needed externally?
#EXPOSE 8090
# Replication (protobuffers) - not needed externally?
#EXPOSE 8099


ENTRYPOINT /opt/influxdb/influxd
