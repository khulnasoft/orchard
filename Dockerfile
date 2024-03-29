FROM goreleaser/goreleaser:latest as builder

WORKDIR /tmp/orchard
ADD . /tmp/orchard/

RUN goreleaser build --single-target --snapshot --timeout 60m

FROM gcr.io/distroless/base

LABEL org.opencontainers.image.source=https://github.com/khulnasoft/orchard
ENV GIN_MODE=release
ENV ORCHARD_HOME=/data
EXPOSE 6120

COPY --from=builder /tmp/orchard/dist/orchard_linux_*/orchard /bin/orchard

ENTRYPOINT ["/bin/orchard"]

# default arguments to run controller
CMD ["controller", "run"]
