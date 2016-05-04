FROM go:latest

ENV REFRESHED_AT 2016-05-04

ADD . /go/src/UniversalOSSWebFileService

RUN go install UniversalOSSWebFileService

RUN cd /go/bin
RUN cp /go/src/UniversalOSSWebFileService/configuration.json .
RUN cp /go/src/UniversalOSSWebFileService/configurationFile.sql .

ENTRYPOINT /go/bin/UniversalOSSWebFileService