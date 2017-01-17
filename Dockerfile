# Building on top of Ubuntu 14.04. The best distro around.
FROM ubuntu:14.04

RUN mkdir /opt/tubackend
COPY . /opt/tubackend
EXPOSE 8080

ENTRYPOINT ["/opt/tubackend"]