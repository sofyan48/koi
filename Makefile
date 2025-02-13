.PHONY: build clean build-darwin build-linux install

build:
	go build -o bin/koi .

clean:
	rm -rf ./bin
	
build-darwin:
	env CGO_ENABLED=0 GOOS=darwin  go build -a -o bin/koi .

build-linux:
	env CGO_ENABLED=1 GOOS=linux go build -v -ldflags '-s -w' -a -tags netgo -installsuffix netgo -o bin/koi .

install:
	cp bin/kaj /usr/local/bin/
	chmod +x /usr/local/bin/kaj