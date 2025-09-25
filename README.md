# Asana 
Inteview assigment


## Running


## Before running the application you need to create a .env file with keys:
```env
ASANA_TOKEN = "token take from asana"
```


### To configure the polling interval you can set the FETCH_POLLING_INTERVAL env variable in seconds, default is 5seconds ex 
or you can change that in the configs/config.toml file under fetch struct
same for the second poll interval

```env
FETCH_POLLING_INTERVAL = "10s"
FETCH_SECOND_POLLING_INTERVAL = "20s"
```

```bash
go run cmd/api/main.go  # to run the application
go mod tidy    # to install libs
```
###If you ar useing linux ur have make installed
```bash
make run # to run the application
make tidy  # to install libs
```




