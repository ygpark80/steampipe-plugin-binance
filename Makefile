STEAMPIPE_INSTALL_DIR ?= ~/.steampipe

install:
	go build -o ${STEAMPIPE_INSTALL_DIR}/plugins/hub.steampipe.io/plugins/ygpark80/binance@latest/steampipe-plugin-binance.plugin *.go

local:
	go build -o ${STEAMPIPE_INSTALL_DIR}/plugins/local/binance/binance.plugin *.go
