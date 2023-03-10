FROM ubuntu:22.04

COPY ./dist/rpiFanControl_linux_arm64/rpiFanControl /FanControl

CMD ["/FanControl"]
