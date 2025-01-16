# Stage 1: Build
FROM golang:1.23.4 AS builder
WORKDIR /app
COPY . .
RUN go mod download && go build -o lsm

# Stage 2: Runtime
FROM ubuntu:24.04
RUN apt-get update && \
    apt-get install -y sshpass && \
    apt-get install -y ansible wget software-properties-common && \
    apt-get clean
WORKDIR /app
COPY --from=builder /app/lsm /usr/local/bin/lsm
COPY ansible /app/ansible
COPY .env /app/.env
RUN mkdir -p /app/config
CMD ["tail", "-f", "/dev/null"]

# # Use an official Ansible base image
# FROM ubuntu:24.04

# # Install Ansible and dependencies
# RUN apt-get update && apt-get install -y ansible wget software-properties-common && \
#     add-apt-repository -y ppa:longsleep/golang-backports && \
#     apt-get update && \
#     apt-get install -y golang && \
#     apt-get clean && \
#     rm -rf /var/lib/apt/lists/*

# # Set environment variables
# ENV GOPATH=/go
# ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH

# WORKDIR /app
# CMD ["/bin/bash"]
