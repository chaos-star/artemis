GO := go
app_name := artemis
app := web
ifeq ($(app),web)
else ifeq ($(app),rpc)
else ifeq ($(app),console)
else
$(error Please specify the application name. Use: make [run|build|clean] app=[web|rpc|console])
endif
src := "main_${app}.go"
dict := "${app_name}_${app}"

.PHONY: run build clean

run:
ifeq ($(app),console)
	@$(GO) run $(src) -name start
else
	@$(GO) run $(src)
endif
build:
	$(GO) build -o $(dict) -tags $(app)
clean:
	@rm -rf $(dict)







