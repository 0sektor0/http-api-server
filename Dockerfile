FROM ubuntu:16.04

MAINTAINER  Egor Grigoryev


#postgresql
RUN apt-get -y update
RUN apt-get -y install apt-transport-https
RUN apt-get -y install wget
RUN echo 'deb http://apt.postgresql.org/pub/repos/apt/ xenial-pgdg main' >> /etc/apt/sources.list.d/pgdg.list
RUN wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add -
RUN apt-get -y update
ENV PGVER 10
RUN apt-get -y install postgresql-$PGVER

USER postgres

COPY data/forum-dump.sql forum-dump.sql

RUN /etc/init.d/postgresql start &&\
    psql --command "CREATE USER forum_admin WITH SUPERUSER PASSWORD 'forum_admin';" &&\
    createdb -O forum_admin forum &&\
    psql -d forum -f forum-dump.sql &&\
    /etc/init.d/postgresql 

RUN echo "host all  all    0.0.0.0/0  md5" >> /etc/postgresql/$PGVER/main/pg_hba.conf

RUN echo "listen_addresses='*'" >> /etc/postgresql/$PGVER/main/postgresql.conf

EXPOSE 5432

VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]

USER root


# GO
RUN apt install -y golang-1.10 git

ENV GOROOT /usr/lib/go-1.10
ENV GOPATH /opt/go
ENV PATH $GOROOT/bin:$GOPATH/bin:/usr/local/go/bin:$PATH

WORKDIR $GOPATH/src/github.com/0sektor0/http-api-server/
ADD ./ $GOPATH/src/github.com/0sektor0/http-api-server/

RUN go get \
    github.com/kataras/iris \
    github.com/lib/pq

RUN go install .

EXPOSE 5000

CMD service postgresql start &&\
    http-api-server