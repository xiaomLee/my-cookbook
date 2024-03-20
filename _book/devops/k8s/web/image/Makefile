# run in docker alpine image need set CGO_ENABLED=0 and  add "-a -installsuffix cgo" build flag
build:
	CGO_ENABLED=0 go build -a -installsuffix cgo

image:
	docker build -t 850278995/go-web:latest .
	docker push 850278995/go-web:latest