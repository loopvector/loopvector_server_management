.PHONY: modsync
modsync:
	go get -u
	go mod tidy
	go mod vendor

.PHONY: clean
clean:
	@if exist vendor (rmdir /S /Q vendor) else (echo Vendor driectory not found, nothing to clean)

.PHONY: build
build: clean modsync
	docker build -t loopvector/lsm:v0.0.2 .
	make clean

.PHONY: push
push: 
	docker tag loopvector/lsm:v0.0.2 loopvector/lsm:v0.0.2
	docker push loopvector/lsm:v0.0.2	

.PHONY: run
run: build push	