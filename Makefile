include .env

test:
	NETFLIX_USERNAME=${NETFLIX_USERNAME} \
	NETFLIX_PASSWORD=${NETFLIX_PASSWORD} \
		go clean -testcache && go test ./... -v

.PHONY: test