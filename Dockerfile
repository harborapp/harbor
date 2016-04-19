FROM alpine:edge

RUN apk update && \
  apk add \
    ca-certificates && \
  rm -rf \
    /var/cache/apk/*

ADD bin/harbor-api /usr/bin/
ENTRYPOINT ["/usr/bin/harbor-api"]
CMD ["server"]
