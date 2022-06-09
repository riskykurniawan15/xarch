# Working with elsa
### Build env if not exist
`elsa create_env`

### Create new domain
`elsa create_domain ${name}`

### Create migration schema
`elsa create_migration_schema ${name}`

### Run migration schema
`elsa run_migration_schema`
### Run specific migration schema
`elsa run_migration_schema ${name}`

### Rollback migration schema
`elsa rollback_migration_schema`
### Rollback specific migration schema
`elsa rollback_migration_schema ${name}`
### Refresh migration schema
`elsa refresh_migration_schema`
### Refresh specific migration schema
`elsa refresh_migration_schema ${name}`

### Create migration seeder
`elsa create_migration_seeder ${name}`

### Run migration seeder
`elsa run_migration_seeder`
### Run specific migration seeder
`elsa run_migration_seeder ${name}`