.PHONY: application rest

application:
	@echo "🛠️  Run 'make application rest' or other targets"

rest:
	@echo "🔄 Running REST client Makefile..."
	$(MAKE) -C cmd/rest-client