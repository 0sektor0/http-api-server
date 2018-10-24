FROM ubuntu:16.04
MAINTAINER  Egor Grigoryev



RUN apt-get -y update
RUN apt-get -y install apt-transport-https git wget
RUN echo 'deb http://apt.postgresql.org/pub/repos/apt/ xenial-pgdg main' >> /etc/apt/sources.list.d/pgdg.list
RUN wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add -
RUN apt-get -y update
ENV PGVER 10
RUN apt-get -y install postgresql-$PGVER



# GO
USER root

RUN git clone https://github.com/0sektor0/http-api-server
RUN ls -l
WORKDIR http-api-server
RUN ls -l
RUN mv http-api-server /usr/bin/
RUN ls -l | grep http-api-server

EXPOSE 5000



#postgres
USER postgres

RUN /etc/init.d/postgresql start &&\
    psql --command "CREATE USER forum_admin WITH SUPERUSER PASSWORD 'forum_admin';" &&\
    createdb -O forum_admin forum &&\
    psql -d forum -a -f ./data/db-init.psql &&\
    /etc/init.d/postgresql stop

EXPOSE 5432

VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]



CMD service postgresql start &&\
    http-api-server