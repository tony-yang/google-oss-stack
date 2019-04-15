FROM ubuntu
  
RUN export DEBIAN_FRONTEND=noninteractive \
 && apt-get update && apt-get install -y \
    build-essential \
    ca-certificates \
    curl \
    git \
    openssl \
    python3 \
    python3-pip \
    unzip \
    vim \
    wget \
 && ln -s /usr/bin/python3 /usr/bin/python

ENV HOME /root

WORKDIR /root

RUN curl -O https://dl.google.com/go/go1.12.4.linux-amd64.tar.gz \
 && tar -C /usr/local -xzf go1.12.4.linux-amd64.tar.gz \
 && mkdir -p ${HOME}/dev/src \
 && mkdir -p ${HOME}/dev/bin \
 && mkdir -p ${HOME}/dev/pkg

ENV GOPATH="${HOME}/dev"
ENV PATH="${PATH}:/usr/local/go/bin:${GOPATH}/bin"

RUN mkdir protobuf \
 && cd protobuf \
 && wget https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/protoc-3.7.1-linux-x86_64.zip \
 && unzip protoc-3.7.1-linux-x86_64.zip \
 && mv bin/protoc /usr/local/bin/protoc \
 && mv include/google /usr/local/include/google \
 && cd ~/dev \
 && go get -u github.com/golang/protobuf/protoc-gen-go \
 && cd ~ \
 && rm -rf protobuf

RUN apt-get update \
 && apt-get install -y openjdk-8-jdk \
 && echo "deb [arch=amd64] http://storage.googleapis.com/bazel-apt stable jdk1.8" | tee /etc/apt/sources.list.d/bazel.list \
 && curl https://bazel.build/bazel-release.pub.gpg | apt-key add - \
 && apt-get update \
 && apt-get install -y bazel

ADD ./addressbook /root/dev/src/addressbook

CMD ["bash"]