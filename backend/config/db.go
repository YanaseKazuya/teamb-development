import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(aws-rds-endpoint:3306)/dbname")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// ここでDB操作を行う
}
