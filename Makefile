DEST = $(GOBIN)

translate : clean
	go build -o bin/translate ./cmd/translate/main.go

install : translate
	cp bin/translate $(DEST)/

clean :
	rm -rf bin/
