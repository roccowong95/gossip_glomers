FROM ubuntu:18.04
USER root
#ADD openjdk-19.0.2_linux-x64_bin.tar.gz go1.20.2.linux-amd64.tar.gz maelstrom.tar.bz2 /root
ADD https://download.java.net/java/GA/jdk19.0.2/fdb695a9d9064ad6b064dc6df578380c/7/GPL/openjdk-19.0.2_linux-x64_bin.tar.gz \
    https://dl.google.com/go/go1.20.2.linux-amd64.tar.gz \
    https://github.com/jepsen-io/maelstrom/releases/download/v0.2.3/maelstrom.tar.bz2 \
    /root
RUN apt update && apt install -y openssl ca-certificates graphviz gnuplot git
RUN echo 'export PATH="/root/maelstrom/:/root/go/bin:/root/jdk-19.0.2/bin:${PATH}"' >> /etc/bash.bashrc && \
    /root/go/bin/go env -w GOPROXY=https://goproxy.cn,direct