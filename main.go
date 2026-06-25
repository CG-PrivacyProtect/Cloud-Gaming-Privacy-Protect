package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"time"

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

	gmodFolder := filepath.Clean(s + `\steamapps\common\GarrysMod`)

	// check 1
	curTime := time.Now()
	os.Chtimes(filepath.Clean(gmodFolder+`\sourceengine\hl2_sound_vo_english_000`), curTime, curTime)

	// check 2
	os.Remove(filepath.Clean(gmodFolder + `\garrysmod\data\pcasino.jpg`))

	// check 3
	sqlDbPath := filepath.Clean(gmodFolder + `\garrysmod\cl.db`)
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
