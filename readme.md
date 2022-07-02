# XARCH

Xarch is a concept backend architecture to implement domain driven design principles and 12 factor app methodology for build a backend application.
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
    \|__|         \/__/         \|__|         \/__/         \/__/    V 1.0
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
go run main.go -engine=${engine}
```
you can type engine to (*) for running all engine or you can type the engine name for run the specific engine. Example:
```bash
go run main.go -engine=http
go run main.go -engine=consumer
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

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://github.com/riskykurniawan15/xarch/blob/main/LICENCE.md/)