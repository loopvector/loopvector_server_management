# Use an official Ansible base image
FROM ubuntu:24.04

# Install Ansible and dependencies
RUN apt-get update && apt-get install -y ansible wget software-properties-common && \
    add-apt-repository -y ppa:longsleep/golang-backports && \
    apt-get update && \
    apt-get install -y golang && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Set environment variables
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH

WORKDIR /app
CMD ["/bin/bash"]
