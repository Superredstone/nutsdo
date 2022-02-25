package db

import (
	"fmt"
	"os"

	"github.com/xujiajun/nutsdb"
)

func OpenDB() *nutsdb.DB {
	opt := nutsdb.DefaultOptions

	userDir, err := os.UserHomeDir()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	opt.Dir = userDir + "/.nutsdo"

	db, err := nutsdb.Open(opt)

	return db
}
