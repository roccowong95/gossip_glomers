FROM ubuntu:18.04
RUN apt update && apt install -y openssl ca-certificates graphviz gnuplot git && \
    useradd --create-home --shell /bin/bash splatoon
# only when adding LOCAL files, docker build will attemp to untar
ADD --chown=splatoon \
    https://download.java.net/java/GA/jdk19.0.2/fdb695a9d9064ad6b064dc6df578380c/7/GPL/openjdk-19.0.2_linux-x64_bin.tar.gz \
    https://dl.google.com/go/go1.20.2.linux-amd64.tar.gz \
    https://github.com/jepsen-io/maelstrom/releases/download/v0.2.3/maelstrom.tar.bz2 \
    /home/splatoon/
RUN echo 'export PATH="/home/splatoon/maelstrom/:/home/splatoon/go/bin:/home/splatoon/jdk-19.0.2/bin:${PATH}"' >> /etc/bash.bashrc
USER splatoon
RUN \
    cd /home/splatoon && ls -lh /home/splatoon && \
    tar zxf openjdk-19.0.2_linux-x64_bin.tar.gz && tar zxf go1.20.2.linux-amd64.tar.gz && tar xopf maelstrom.tar.bz2 && \
    /home/splatoon/go/bin/go env -w GOPROXY=https://goproxy.cn,direct && \
    rm openjdk-19.0.2_linux-x64_bin.tar.gz go1.20.2.linux-amd64.tar.gz maelstrom.tar.bz2