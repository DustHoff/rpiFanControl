FROM ubuntu:22.04
LABEL org.opencontainers.image.source https://github.com/DustHoff/rpifancontrol
COPY dist/rpifancontrol_linux_arm64/rpifancontrol /FanControl

CMD ["/FanControl"]
