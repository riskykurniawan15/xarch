package elsa

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func RunElsa() {
	var cmd, arg string = "", ""
	if len(os.Args) > 2 {
		cmd = strings.ToLower(os.Args[2])

		if len(os.Args) > 3 {
			arg = strings.ToLower(os.Args[3])
		}

		var err error = nil

		if cmd == "create_env" {
			err = BuildEnvirontment()
		} else if cmd == "flush_log" {
			err = FlushLog()
		} else if cmd == "create_domain" {
			err = CreateDomain(arg)
		} else if cmd == "create_migration_schema" {
			err = CreateMigrationSchema(arg)
		} else if cmd == "run_migration_schema" {
			err = AutoMigrationSchema(arg, "up")
		} else if cmd == "rollback_migration_schema" {
			err = AutoMigrationSchema(arg, "down")
		} else if cmd == "refresh_migration_schema" {
			err = AutoMigrationSchema("", "refresh")
		} else if cmd == "create_migration_seeder" {
			err = CreateMigrationSeeder(arg)
		} else if cmd == "run_migration_seeder" {
			err = AutoMigrationSeeder(arg)
		} else {
			err = fmt.Errorf("failed run command elsa")
		}

		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Please type your command")
	}
}
