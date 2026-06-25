package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/sys/windows/registry"
)

func main() {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Wow6432Node\Valve\Steam`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatalln(err)
	}
	defer k.Close()

	s, _, err := k.GetStringValue("InstallPath")
	if err != nil {
		log.Fatalln(err)
	}

	gmodFolder := filepath.Clean(s + `\steamapps\common\GarrysMod\garrysmod`)

	// check 1
	os.Remove(filepath.Clean(gmodFolder + `\data\pcasino.jpg`))

	// check 2
	sqlDbPath := filepath.Clean(gmodFolder + `\cl.db`)
	db, err := sql.Open("sqlite3", sqlDbPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// fuck you GoDofWaR and wilton
	// you're alright
	// you're just jewish
	db.Exec("DELETE FROM cookies")
}
