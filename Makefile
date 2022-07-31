DATA_DIR ?= data

dep:

etl: dep
	@go build -v .
	@./buying-frenzy etl -d $(DATA_DIR)

clean: