test:
	go test -v --cover -count=1 ./...

bench:
	go test -v ./... -bench=. -benchmem