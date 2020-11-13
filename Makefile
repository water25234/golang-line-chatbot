.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/LiaoLiao app/main.go

clean:
	rm -rf bin

deploy: 
	make clean 
	make build
	sls deploy --verbose