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
			err = BuildDomain(arg)
		} else if cmd == "create_migration_schema" {
			err = CreateMigrationSchema(arg)
		} else if cmd == "run_migration_schema" {
			err = RunMigrationSchema(arg, "up")
		} else if cmd == "rollback_migration_schema" {
			err = RunMigrationSchema(arg, "down")
		} else if cmd == "refresh_migration_schema" {
			err_down := RunMigrationSchema(arg, "down")
			if err_down != nil {
				log.Fatal(err_down)
			}
			err_up := RunMigrationSchema(arg, "up")
			if err_up != nil {
				log.Fatal("%w", err_up)
			}
		} else if cmd == "create_migration_seeder" {
			err = CreateMigrationSeeder(arg)
		} else if cmd == "run_migration_seeder" {
			err = RunMigrationSeeder(arg)
		} else {
			err = fmt.Errorf("Failed Run Command Elsa")
		}

		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Please type your command")
	}
}
