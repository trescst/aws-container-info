FROM golang:1.6
MAINTAINER @steventrescinski
COPY aws-container-info /usr/bin/aws-container-info
COPY index.html /index.html
WORKDIR /
EXPOSE 9000
ENTRYPOINT ["aws-container-info"]
