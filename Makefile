run:
	@rm -r ./bin
	@go build -o bin/tcli .
	@./bin/tcli $(filter-out $@, $(MAKECMDGOALS))
%:
	@true
build:
	@rm -r ./bin
	@go build -o bin/tcli .
