TEMPLATE_SERVICE=./cmd/example

start: ##@Start new problem from example.
	cp -r $(TEMPLATE_SERVICE) ./cmd/problem
