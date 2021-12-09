#  Monitor tool
Logic is in internal/monitor/

monitor.go has got the necessary functions to retrieve stats using docker API.
helper.go has functions to calculate resource usage.
fields.go contains structs to parse JSON i.e. the stats/json received from the API.
