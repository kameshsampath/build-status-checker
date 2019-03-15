DESTINATION_NAME := 'quay.io/rhdevelopers/build-status-checker'

default: bin/kbsc

bin:
	mkdir -p bin

bin/kbsc: bin
	GOOS=linux go build -o bin/kbsc -v .

.PHONY: $(DESTINATION_NAME)
$(DESTINATION_NAME): bin/kbsc
	docker build -t $(DESTINATION_NAME) --rm . && docker push $(DESTINATION_NAME)

.PHONY: clean
clean:
	rm -rf bin 

.PHONY: test
test:
	kubectl apply -f test/deployment.yaml

.PHONY: dep
dep:
	dep ensure

.PHONY: all
all: clean $(DESTINATION_NAME)