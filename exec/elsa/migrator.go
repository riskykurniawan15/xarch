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
func CreateMigrationSchema(arg string) error {
	if arg == "" {
		return fmt.Errorf("Failed create migration schema")
	}

	files, err := os.OpenFile(SchemaLoc+Prefix()+arg+".sql", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer files.Close()
	if _, err := files.WriteString("--Up\n/* Type Your SQL Code in Here */\n--EndUp\n--Down\n/* Type Your SQL Code in Here */\n--EndDown"); err != nil {
		return err
	}

	log.Println("Success create migration schema " + arg)

	return nil
}

func RunMigrationSchema(arg string, method string) error {
	var FileList []string

	DB := DB_Driver()

	MigrationTable := MigrationTableCheck(DB, method)

	if method == "up" {
		if arg == "" {
			for _, f := range migrate.ExecSchema() {
				FileList = append(FileList, f)
			}
		} else {
			FileList = append(FileList, arg)
		}
	} else {
		if arg == "" {
			for _, f := range MigrationTable {
				FileList = append(FileList, f.Name)
			}
		} else {
			FileList = append(FileList, arg)
		}
	}

	for _, f := range FileList {
		content, err := os.ReadFile(SchemaLoc + f)
		if err != nil {
			return err
		}
		script := ReadScript(string(content), method)
		isRun, err := RunScript(DB, f, script, "schema", method)
		if err != nil {
			return err
		}
		if isRun {
			if method == "up" {
				log.Println("Success run migration schema " + f)
			} else {
				log.Println("Success rollback migration schema " + f)
			}
		}
	}
	return nil
}

/* End Schema Script */

/* Seeder Script */
func CreateMigrationSeeder(arg string) error {
	if arg == "" {
		return fmt.Errorf("Failed create migration seeder")
	}

	files, err := os.OpenFile(SeederLoc+Prefix()+arg+".sql", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer files.Close()
	if _, err := files.WriteString("--Up\n/* Type Your SQL Code in Here */\n--EndUp"); err != nil {
		return err
	}

	log.Println("Success create migration schema " + arg)

	return nil
}

func RunMigrationSeeder(arg string) error {
	var FileList []string

	DB := DB_Driver()

	if arg == "" {
		for _, f := range migrate.ExecSeeder() {
			FileList = append(FileList, f)
		}
	} else {
		FileList = append(FileList, arg)
	}

	for _, f := range FileList {
		content, err := os.ReadFile(SeederLoc + f)
		if err != nil {
			return err
		}
		script := ReadScript(string(content), "up")
		isRun, err := RunScript(DB, f, script, "seeder", "up")
		if err != nil {
			return err
		}
		if isRun {
			log.Println("Success run migration seeder " + f)
		}
	}
	return nil
}

/* End Seeder Script */
