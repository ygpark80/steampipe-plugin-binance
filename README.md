# Binance Plugin for Steampipe

## Installing and Testing the Plugin

To install the plugin, simple run the following command.

```
% make local
go build -o  ~/.steampipe/plugins/local/binance/binance.plugin *.go
```

Check your local plugin using the following command.

```
% steampipe plugin list
+--------------------------------------------------+---------+-------------+
| Name                                             | Version | Connections |
+--------------------------------------------------+---------+-------------+
| hub.steampipe.io/plugins/turbot/aws@latest       | 0.57.0  | aws         |
| hub.steampipe.io/plugins/turbot/steampipe@latest | 0.2.0   | steampipe   |
| local/binance                                    | local   |             |
+--------------------------------------------------+---------+-------------+
```

Copy the sample `binance.spc` file to `~/.steampipe/config` folder and change the name of the `plugin` from `binance` to `local/binance`.

```
% cat ~/.steampipe/config/binance.spc
connection "binance" {
    plugin = "binance"

    # api_key = "YOUR_API_KEY_HERE"
    # api_secret = "YOUR_API_SECRET_HERE"
}
```

Check and see if you have a valid connection.

```
% steampipe plugin list
+--------------------------------------------------+---------+-------------+
| Name                                             | Version | Connections |
+--------------------------------------------------+---------+-------------+
| hub.steampipe.io/plugins/turbot/aws@latest       | 0.57.0  | aws         |
| hub.steampipe.io/plugins/turbot/steampipe@latest | 0.2.0   | steampipe   |
| local/binance                                    | local   | binance     |
+--------------------------------------------------+---------+-------------+
```

Let's test the plugin.

```
% steampipe query "select count(*) from binance_staking" --timing
+-------+
| count |
+-------+
| 334   |
+-------+

Time: 355.912958ms
```

That's it.

## Testing

```
% go test
```

## Update Go dependencies

```
% go get -u
% go mod tidy
```
