default: bin/kbsc

bin:
	mkdir -p bin

bin/kbsc: bin
	GOOS=linux go build -o bin/kbsc -v .

.PHONY: quay.io/rhdevelopers/build-status-checker
quay.io/rhdevelopers/build-status-checker:
	docker build -t quay.io/rhdevelopers/build-status-checker --rm .

.PHONY: clean
clean:
	rm -rf bin 

.PHONY: test
test:
	kubectl apply -f test/deployment.yaml

.PHONY: dep
dep:
	dep ensure