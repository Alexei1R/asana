# Asana  
Interview Assignment

## Running

Before running the application you need to create a `.env` file with the following keys:

```env
ASANA_TOKEN="token taken from Asana"

# Polling intervals (default 5s for first, 10s for second)
FETCH_POLLING_INTERVAL="10s"
FETCH_SECOND_POLLING_INTERVAL="20s"

# Refresh interval (default 1h)
REFRESH_INTERVAL="1h"
```

You can also adjust the polling and refresh intervals in `configs/config.toml` under the `fetch` and `refresh` structs.

## Commands

```bash
go mod tidy        # install dependencies
go run cmd/api/main.go  # run the application
```

If you are using Linux and have `make` installed:

```bash
make tidy   # install dependencies
make run    # run the application
make test   # run all tests
```



```
cmd/
configs/
internal/
pkg/
docs/
Makefile
go.mod
go.sum
README.md
```

## Explanation of folders

- **cmd**: Entry point(s) for the application.  
  - `main.go` inside `cmd/api/` is the main executable that starts the Asana application.

- **configs**: Contains configuration files like `config.toml`.  
  - `config.toml` is the main configuration file for the app.

- **internal**: Core application logic and domain code.  
  - Includes application services, domain models, and infrastructure implementations.

- **pkg**: Reusable packages not tied to a single application, like `config`, `log`, or `storage`.

- **docs**: Documentation.  

- **Makefile**: Commands to build, run, tidy, and test the project.  

