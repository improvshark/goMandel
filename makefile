
build:
	go build  -a -tags netgo -ldflags '-w -extldflags "-static"' -o goMandel *.go

docker: build
	docker build -t go-mandel . --no-cache

run: docker
	docker run -p 3000:3000 go-mandel
