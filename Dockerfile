FROM ubuntu:22.04

COPY dist/rpifancontrol_linux_arm64/rpifancontrol /FanControl

CMD ["/FanControl"]
