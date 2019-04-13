FROM go-dev

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
