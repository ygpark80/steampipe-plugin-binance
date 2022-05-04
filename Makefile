install:
	go build -o  ~/.steampipe/plugins/hub.steampipe.io/plugins/ygpark80/binance@latest/steampipe-plugin-binance.plugin *.go

local:
	go build -o  ~/.steampipe/plugins/local/binance/binance.plugin *.go
