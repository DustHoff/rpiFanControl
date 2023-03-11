FROM ubuntu:22.04

COPY --chmod=755 dist/rpifancontrol_linux_arm64/rpifancontrol /

RUN apt update -y && \
    apt install -y curl && \
    apt clean -y && \
    rm -rf /var/lib/apt/lists/* /var/cache/apt/*

CMD ["/rpifancontrol"]
