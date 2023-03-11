FROM ubuntu:22.04

COPY --chmod 755 dist/rpifancontrol_linux_arm64/rpifancontrol /

CMD ["/FanControl"]
