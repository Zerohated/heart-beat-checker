PROJECTNAME=$(shell basename "$(PWD)")

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

define RECOMMENDED_TAG=
$(git tag -l | sort -V | tail -n 1 | sed 's#\.# #g' | awk '{ print $1"."$2 + 1".0" }')
endef
define NOW=
$(date "+%Y-%m-%d %H:%M:%S")
endef

name:
	@echo $(PROJECTNAME)

install:
	go mod download

start:
	go build -o $(GOBIN)/$(PROJECTNAME) ./cmd/$(PROJECTNAME)/main.go || exit
	./bin/$(PROJECTNAME)

test:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(GOBIN)/run ./cmd/$(PROJECTNAME)/main.go || exit
	docker build -t registry.cn-hangzhou.aliyuncs.com/metro/$(PROJECTNAME) .
	docker push registry.cn-hangzhou.aliyuncs.com/metro/$(PROJECTNAME)

docker:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(GOBIN)/run ./cmd/$(PROJECTNAME)/main.go || exit
	@echo "Creating tags... Current tag is " $$(git tag -l | sort -V | tail -n 1)
	@printf "Input new tag [%s]: " $(value RECOMMENDED_TAG)
	@read tag;\
	tag="$${tag:-$(value RECOMMENDED_TAG)}";\
	git tag -a $$tag -m "Published at $(value NOW)";\
	docker build -t registry.cn-hangzhou.aliyuncs.com/metro/$(PROJECTNAME):$$tag -t registry.cn-hangzhou.aliyuncs.com/metro/$(PROJECTNAME):latest .
	docker push registry.cn-hangzhou.aliyuncs.com/metro/$(PROJECTNAME)