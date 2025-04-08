.PHONY: all build clean run

all: build

build:
	go build -o stock-balancer

run: build
	@shifted=$$(echo $(MAKECMDGOALS) | sed 's/^run //'); \
	./stock-balancer $$shifted

clean:
	rm -f stock-balancer