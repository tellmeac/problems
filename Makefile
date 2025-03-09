TEMPLATE_SERVICE=./cmd/example_std

start: ##@Start new problem from example.
	cp -r $(TEMPLATE_SERVICE) ./cmd/problem
