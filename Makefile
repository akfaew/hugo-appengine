PROJECT = XXX

## development
.PHONY: hugo
hugo:
	(cd hugo && hugo server --disableFastRender -D)

build:
	rm -rf hugo/public
	(cd hugo && hugo -D)

run:
	go run *.go

update:
	go get -u
	go mod tidy
	go mod verify

## release
deploy: build
	echo y | gcloud app deploy app.yaml --project=$(PROJECT)

push:
	git push origin

release: deploy push
