FROM registry:2

WORKDIR /opt
RUN mkdir /opt/data
ADD templates /opt/templates
ADD static /opt/static
ADD config.yml /opt/config.yml
ADD docker-registry-ui /opt/
ADD registry.yml /etc/docker/registry/config.yml
COPY entrypoint.sh /

EXPOSE 8000