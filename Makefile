all: build-container run

build-container:
	docker build -t google-stack .

run: build-container
	docker run -itd --rm google-stack

test: build-container
	# docker run --rm google-stack bash -c 'cd ~/dev/src/addressbook/proto && protoc --go_out=. ./addressbook.proto && cd ~/dev/src/addressbook && go test -v'
	docker run --rm google-stack bash -c 'cd ~/dev/src/addressbook && bazel run :gazelle && bazel test --test_output all //:go_default_test'
