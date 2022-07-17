package elsa

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/riskykurniawan15/xarch/config"
	"github.com/riskykurniawan15/xarch/driver"
	migrate "github.com/riskykurniawan15/xarch/migration"
)

const (
	SchemaLoc string = "migration/schema/"
	SeederLoc string = "migration/seeder/"
)

type Migration struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"column:migration_name;unique;size:100"`
	Status    bool      `gorm:"column:migration_status;type:bool;default:false"`
	CreatedAt time.Time `gorm:"column:created_at;"`
	UpdatedAt time.Time `gorm:"column:updated_at;"`
}

func (Migration) TableName() string {
	return "migration_histories"
}

func Prefix() string {
	return time.Now().Format("20060102150405") + "_"
}

func DB_Driver() *gorm.DB {
	cfg := config.Configuration()
	DB := driver.ConnectDB(cfg.DB)

	return DB
}

func MigrationTableCheck(DB *gorm.DB, method string) []*Migration {
	var model []*Migration

	exist := DB.Migrator().HasTable(&Migration{})
	if exist != true {
		DB.AutoMigrate(&Migration{})
	}

	var order string
	where := &Migration{}
	if method == "up" {
		order = "asc"
	} else {
		where.Status = true
		order = "desc"
	}

	result := DB.
		Model(&Migration{}).
		Where(where).
		Order("updated_at " + order).
		Find(&model)

	if result.Error != nil {
		panic(result.Error)
	}

	return model
}

func ReadScript(str string, method string) (out string) {
	var begin, end string = "", ""
	if method == "up" {
		begin = "--Up"
		end = "--EndUp"
	} else {
		begin = "--Down"
		end = "--EndDown"
	}

	scriptFirst := strings.Index(str, begin)
	if scriptFirst == -1 {
		return ""
	}
	scriptLast := strings.Index(str, end)
	if scriptLast == -1 {
		return ""
	}
	scriptFirstAdjusted := scriptFirst + len(begin)
	if scriptFirstAdjusted >= scriptLast {
		return ""
	}

	return str[scriptFirstAdjusted:scriptLast]
}

func RunScript(DB *gorm.DB, files, str, tipe, method string) (bool, error) {
	run_script := true
	histories := &Migration{}
	model := &Migration{}

	if tipe == "schema" {
		run_script = false

		if DB.Model(&Migration{}).Where("migration_name = ?", files).Find(histories).RowsAffected == 0 {
			histories.Name = files
			histories.Status = false
			DB.Model(&Migration{}).Create(histories)
			run_script = true
		}

		if method == "up" {
			model.Status = true
		} else {
			model.Status = false
		}

		if histories.Status != model.Status {
			run_script = true
		}
	}

	if run_script {
		if err := DB.Exec(str).Error; err != nil {
			return false, err
		}
	}

	if tipe == "schema" {
		DB.Model(&Migration{}).Where("migration_name = ?", files).Update("migration_status", model.Status)
	}
	return run_script, nil
}

/* Schema Script */
func CreateMigrationSchema(schema_name string) error {
	if schema_name == "" {
		return fmt.Errorf("failed create migration schema")
	}

	files, err := os.OpenFile(SchemaLoc+Prefix()+schema_name+".sql", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer files.Close()
	if _, err := files.WriteString("--Up\n/* Type Your SQL Code in Here */\n--EndUp\n--Down\n/* Type Your SQL Code in Here */\n--EndDown"); err != nil {
		return err
	}

	log.Println("Success create migration schema " + schema_name)

	return nil
}

func AutoMigrationSchema(schema_name string, method string) error {
	DB := DB_Driver()

	MigrationTable := MigrationTableCheck(DB, method)

	if schema_name == "" {
		if method == "down" || method == "refresh" {
			fmt.Print(method)
			for _, schema := range MigrationTable {
				if err := execMigration(DB, schema.Name, "down", "schema"); err != nil {
					return err
				}
			}
		}
		if method == "up" || method == "refresh" {
			for _, schema := range migrate.ExecSchema() {
				if err := execMigration(DB, schema, "up", "schema"); err != nil {
					return err
				}
			}
		}
	} else {
		if err := execMigration(DB, schema_name, method, "schema"); err != nil {
			return err
		}
	}

	return nil
}

/* End Schema Script */

/* Seeder Script */
func CreateMigrationSeeder(seeder_name string) error {
	if seeder_name == "" {
		return fmt.Errorf("failed create migration seeder")
	}

	files, err := os.OpenFile(SeederLoc+Prefix()+seeder_name+".sql", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer files.Close()
	if _, err := files.WriteString("--Up\n/* Type Your SQL Code in Here */\n--EndUp"); err != nil {
		return err
	}

	log.Println("Success create migration schema " + seeder_name)

	return nil
}

func AutoMigrationSeeder(seeder_name string) error {
	DB := DB_Driver()

	if seeder_name == "" {
		for _, seeder := range migrate.ExecSeeder() {
			if err := execMigration(DB, seeder, "up", "seeder"); err != nil {
				return err
			}
		}
	} else {
		if err := execMigration(DB, seeder_name, "up", "seeder"); err != nil {
			return err
		}
	}

	return nil
}

/* End Seeder Script */

/* Exec Migration Script */
func execMigration(DB *gorm.DB, name, method, types string) error {
	var location string
	if types == "schema" {
		location = SchemaLoc
	} else {
		location = SeederLoc
	}
	content, err := os.ReadFile(location + name)
	if err != nil {
		return err
	}
	script := ReadScript(string(content), method)
	isRun, err := RunScript(DB, name, script, types, method)
	if err != nil {
		return err
	}
	if isRun {
		if method == "up" {
			log.Println(fmt.Sprintf("Success run migration %s %s", types, name))
		} else {
			log.Println(fmt.Sprintf("Success rollback migration %s %s", types, name))
		}
	}

	return nil
}

/* End Exec Migration Script */
