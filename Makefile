build: main.go
	@mkdir -p target
	@go build -o target/connect-kafka

clean:
	@rm -rf target

.PHONY: clean
