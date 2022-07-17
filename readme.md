# XARCH

Xarch is a concept backend architecture that applies implement domain driven design principles and 12 factor app methodology for build a backend application.
```bash
     ___           ___           ___           ___           ___     
    |\__\         /\  \         /\  \         /\  \         /\__\    
    |:|  |       /::\  \       /::\  \       /::\  \       /:/  /    
    |:|  |      /:/\:\  \     /:/\:\  \     /:/\:\  \     /:/__/     
    |:|__|__   /::\~\:\  \   /::\~\:\  \   /:/  \:\  \   /::\  \ ___ 
____/::::\__\ /:/\:\ \:\__\ /:/\:\ \:\__\ /:/__/ \:\__\ /:/\:\  /\__\
\::::/~~/~    \/__\:\/:/  / \/_|::\/:/  / \:\  \  \/__/ \/__\:\/:/  /
 ~~|:|~~|          \::/  /     |:|::/  /   \:\  \            \::/  / 
   |:|  |          /:/  /      |:|\/__/     \:\  \           /:/  /  
   |:|  |         /:/  /       |:|  |        \:\__\         /:/  /   
    \|__|         \/__/         \|__|         \/__/         \/__/    V 1.0.0
By: Risky Kurniawan | https://risoftinc.com | mailto:riskykurniawan@risoftinc.com
```

## Installation Requirement

Install golang v 1.8 and install all dependencies
```bash
go get ./...
```

## Running Xarch

For running xarch program open the console, bash or terminal
```bash
go run main.go xarch -engine=${engine}
```
you can type engine to (*) for running all engine or you can type the engine name for run the specific engine. Example:
```bash
go run main.go xarch -engine=http
go run main.go xarch -engine=consumer
```

## Working With Elsa (Electronic Smart Assistant)
Elsa is a tools for making your happy and enjoy to developing the program. Elsa have a many function for assist developing a program. You can say and use elsa with type command understandable in elsa program. Example you can command:
```bash
# Build env if not exist
go run main.go elsa create_env

# Flush all log file
go run main.go elsa flush_log

# Create new domain
go run main.go elsa create_domain ${domain_name}

# Create migration schema
go run main.go elsa create_migration_schema ${schema_name}
# Run migration schema
go run main.go elsa run_migration_schema
# Run specific migration schema
go run main.go elsa run_migration_schema ${schema_name}
# Rollback migration schema
go run main.go elsa rollback_migration_schema
# Rollback specific migration schema
go run main.go elsa rollback_migration_schema ${schema_name}
# Refresh migration schema
go run main.go elsa refresh_migration_schema
# Refresh specific migration schema
go run main.go elsa refresh_migration_schema ${schema_name}

# Create migration seeder
go run main.go elsa create_migration_seeder ${seeder_name}
# Run migration seeder
go run main.go elsa run_migration_seeder
# Run specific migration seeder
go run main.go elsa run_migration_seeder ${seeder_name}
```

## Folder Structure
```tree
.
├───app                     # Application layer contains all program in architecture
│   ├───elsa                # Elsa Application (tools)
│   ├───xarch               # Xarch Application (services)
│   └───app.go              # Program for handling switching program application
├───config                  # All configuration in application
├───domain                  # Domain Layers is containing a models, repository and services
│   ├───domain_name         # Name of domain 
│   │   ├───models          # Contains many file for modeling data
│   │   ├───repository      # All file to comunicated all backing services
│   │   └───services        # Services file to processing a bussiness application
│   └───domain.go           # For publishing a domain and inject all dependencies to domain
├───driver                  # Dir to defined and registered all driver (backing services)
├───helpers                 # Dir to contains all helpering program
│   ├───bcrypt              # Hash and compare bcrypt
│   ├───env                 # Program to get environtment to specifics format data type
│   ├───errors              # Defined a error code to communicated domain and interfaces
│   ├───jwt                 # Jwt token helper
│   ├───mail                # Mail sender helper
│   └───md5                 # Converted plain text to MD5
├───interfaces              # Interfaces Layer contains all engine method for user access a application
│   └───xarch               # Group interfaces by app
│       ├───consumer        # Consumer Engine
│       │   ├───engine      # Define a engine script for running a consumer interfaces
│       │   ├───handlers    # All handler in consumer interfaces
│       │   └───consumer.go # Routing and defined all job in consumer interfaces
│       └───http            # Http Engine
│           ├───engine      # Define a engine script for running a http server interfaces
│           ├───entities    # Define to entities models contains a response formater
│           ├───handlers    # All handler in http interfaces
│           ├───middleware  # All handler middleware in http interfaces
│           └───routers     # Routing http server
├───logger                  # Contains a logger program and dir to save all log file
├───migration               # Migration layers can acessed with elsa program
│   ├───schema              # Containing script schema in sql format
│   ├───seeder              # Containing script seeder in sql format
│   └───migrate.go          # For registered schema and seeder
└───main.go                 # Main Application
````
## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://github.com/riskykurniawan15/xarch/blob/main/LICENCE.md/)
