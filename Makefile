.PHONY: install ask build

build:
	go build -o chatbotcli main.go

install:
	@if [ ! -f chatbotcli ]; then \
		echo "Building chatbotcli..."; \
		make build; \
	fi
	./chatbotcli install --token=$(token)

ask:
	@if [ ! -f chatbotcli ]; then \
		echo "Building chatbotcli..."; \
		make build; \
	fi
	./chatbotcli ask
