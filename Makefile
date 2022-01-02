.PHONY: generate lint lint-fix test

default: generate lint test

generate: clean-generate
	go generate -x internal/generator.go

clean-generate:
	@rm -f application.go
	@rm -f audio.go
	@rm -f font.go
	@rm -f image.go
	@rm -f message.go
	@rm -f model.go
	@rm -f multipart.go
	@rm -f text.go
	@rm -f video.go

test:
	go test ./... --cover

lint:
	golangci-lint run
