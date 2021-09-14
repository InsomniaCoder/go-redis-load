BINARY=go-redis-load
test: 
	go test -v -cover -covermode=atomic ./...

local:
	go run main.go

lint:
	go fmt ./...

build:
	go build -o ${BINARY} *.go

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t ${BINARY} .

run:
	docker-compose up --build -d

stop:
	docker-compose down