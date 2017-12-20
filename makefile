
build:
	go build  -a -tags netgo -ldflags '-w -extldflags "-static"' -o goMandel *.go

serve: build
	./goMandel

clean:
	rm ./goMandel

pack:
	docker build -t improvshark/go-mandel .

upload: pack
	docker tag go-mandel improvshark/go-mandel; docker push improvshark/go-mandel

deploy:
	kubectl apply -f k8s/deployment.yml

run: pack
	docker run -p 3000:3000 go-mandel
