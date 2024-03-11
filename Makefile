.PHONY: test bench

test:  ## Run test
	go test -p 8 -v --cover $(GOPKGS)

bench: ## Run benchmark
	go test -bench=. -benchmem -run=^$$ $(GOPKGS)