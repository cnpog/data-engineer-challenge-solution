detected_OS := $(shell uname)
ifeq ($(detected_OS),Darwin)
    LOCAL_IP=host.docker.internal:9092
endif
ifeq ($(detected_OS),Linux)
    LOCAL_IP=127.0.0.1:9092
endif


dockerBasic:
	docker build . -t data-engineer-challenge-basic -f cmd/basic/Dockerfile
dockerAdvanced:
	docker build . -t data-engineer-challenge-advanced -f cmd/advanced/Dockerfile
dockerBonus:
	docker build . -t data-engineer-challenge-bonus -f cmd/bonus/Dockerfile

dockerAll:
	docker build . -t data-engineer-challenge-basic -f cmd/basic/Dockerfile
	docker build . -t data-engineer-challenge-advanced -f cmd/advanced/Dockerfile
	docker build . -t data-engineer-challenge-bonus -f cmd/bonus/Dockerfile

runDockerBasic:
	docker run --env KAFKAIP=$(LOCAL_IP) --env KAFKATOPIC=mytopic --net=host -it data-engineer-challenge-basic
runDockerAdvanced:
	docker run --env KAFKAIP=$(LOCAL_IP) --env KAFKATOPIC=mytopic --env KAFKATOPICOUT=mytopicout --net=host -it data-engineer-challenge-advanced
runDockerBonus:
	docker run -it data-engineer-challenge-bonus

runTest:
	go test -v ./pkg/counting