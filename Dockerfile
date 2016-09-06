FROM alpine:edge

RUN apk update && \
  apk add \
    ca-certificates \
    sqlite && \
  rm -rf \
    /var/cache/apk/*

ADD bin/umschlag-api /usr/bin/
ENTRYPOINT ["/usr/bin/umschlag-api"]
CMD ["server"]
