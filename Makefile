.PHONY: api
api:
	@echo "Generating api..."
	goctl api go --api api/general.api --dir . 

.PHONY: curd
curd:
	@echo "Generating crud"
	goctl model mysql ddl -src internal/sql/score.sql -dir ./internal/model -c --cache=false