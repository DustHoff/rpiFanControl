FROM ubuntu:22.04

COPY --chmod=755 dist/rpifancontrol_linux_arm64/rpifancontrol /

RUN apt-get install -y curl && \
    apt-cache clean && \
    rm -rf /var/lib/apt/lists/* /var/cache/apt/* \

CMD ["/rpifancontrol"]
