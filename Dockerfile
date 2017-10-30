FROM webhippie/alpine:latest
MAINTAINER Thomas Boerger <thomas@webhippie.de>

EXPOSE 8080 80 443
VOLUME ["/var/lib/umschlag"]

LABEL org.label-schema.version=latest
LABEL org.label-schema.name="Umschlag API"
LABEL org.label-schema.vendor="Thomas Boerger"
LABEL org.label-schema.schema-version="1.0"

RUN apk update && \
  apk add \
    ca-certificates \
    bash \
    sqlite && \
  rm -rf \
    /var/cache/apk/* && \
  addgroup \
    -g 1000 \
    umschlag && \
  adduser -D \
    -h /var/lib/umschlag \
    -s /bin/bash \
    -G umschlag \
    -u 1000 \
    umschlag && \
  mkdir -p \
    /usr/share/umschlag

COPY assets /usr/share/umschlag/
COPY templates /usr/share/umschlag/
COPY umschlag-api /usr/bin/

ENV UMSCHLAG_SERVER_STORAGE /var/lib/umschlag

USER umschlag
ENTRYPOINT ["/usr/bin/umschlag-api"]
CMD ["server"]
