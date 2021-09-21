include .env

test:
	# clear cache given website may change and cache only considers code
	go clean -testcache

	NETFLIX_USERNAME=${NETFLIX_USERNAME} \
	NETFLIX_PASSWORD=${NETFLIX_PASSWORD} \
		go test ./netflix -v

.PHONY: test