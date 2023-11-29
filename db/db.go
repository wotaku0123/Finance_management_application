package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// gormを用いたデータベースの接続を表す変数
var DB *gorm.DB

// データベースへの接続を初期化するもの
// gorm.Openを使用してMySQLデータベースへの接続を開始する
// 成功した場合はその接続は DB変数に割り当てられる
func Init() {
	log.Println("func Init() is called")
	db, err := gorm.Open(mysql.Open("TakumiNakagawara:@Seraio12@(127.0.0.1:3306)/user_information?charset=utf8mb4&parseTime=True&loc=Local"))

	var result int
	if err := db.Raw("SELECT 1").Scan(&result).Error; err != nil {
		log.Fatal("Database test query failed: %v", err)
	} else if result != 1 {
		log.Fatal("Database test query returned unexpected result")
	} else {
		log.Println("Database test query succeeded")
	}

	if err != nil {
		log.Fatal("Failed to open database connection: %v", err)
	}

	sqlDB, err := db.DB()
	if sqlDB == nil {
		log.Fatal("Database connection is nil")
	} else {
		log.Println("Database connection successfully established")
	}

	DB = db
	//Model自動生成
	//genModel()
}
func genModel() {
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./query",
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		WithUnitTest:      true,
		FieldNullable:     false,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})
	g.UseDB(DB) // reuse your gorm db
	// Generate structs from all tables of current database
	allModel := g.GenerateAllTable()

	g.ApplyBasic(allModel...)
	// Generate the code
	g.Execute()
}
