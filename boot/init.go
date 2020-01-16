package boot

import (
	"log"
	"os"
)

func Init() {
	mysql := GetInstance().InitMysqlConnPool()
	if !mysql {
		log.Println("init database pool failure...")
		os.Exit(1)
	}
}
