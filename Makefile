build:
	go run /app/main.go

run: build
	./server

watch:
	reflex -s -r '\.go$$' make run