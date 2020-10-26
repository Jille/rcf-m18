gen: parser.go
	go build -v

run: gen
	./gen 2>&1 | tee out.new
	diff -u out.prev out.new || true
	mv out.new out.prev
	goimports < out.prev > ../rcfm18/generated.go

fmt:
	goimports -w -d parser.go
