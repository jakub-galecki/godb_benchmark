package main

import (
	"database/sql"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/dustin/go-humanize"
	_ "github.com/mattn/go-sqlite3"

	"github.com/jakub-galecki/godb"
)

func printFileSizes(dirPath string) {
	filepath.Walk(dirPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
			return filepath.SkipDir
		}
		log.Printf("file: %s,  size: %v\n", path, humanize.Bytes(uint64(info.Size())))
		return nil
	})
}

func main() {
	db, err := godb.Open("thesis_test", godb.WithDbPath("/tmp"))
	if err != nil {
		log.Fatalln(err)
	}
	start := time.Now()
	// insert 100 000 keys to database
	for i := 0; i < 100000; i++ {
		err := db.Set([]byte(fmt.Sprintf("foo.%d", i)), []byte(fmt.Sprintf("bar.%d", i)))
		if err != nil {
			log.Fatalln(err)
		}
	}
	fmt.Printf("inserting 100000 keys to my implementation took: %v\n", time.Since(start))
	fmt.Println("stats for my implementation files")
	printFileSizes("/tmp/thesis_test")

	// sqlite
	dbDir, err := os.MkdirTemp("", "test")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creating sqlite file in: " + dbDir)

	defer os.RemoveAll(dbDir)

	dbsql, err := sql.Open("sqlite3", dbDir+"/test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = dbsql.Exec("CREATE TABLE test (id TEXT, name TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	start = time.Now()
	for i := 0; i < 100000; i++ {
		_, err := dbsql.Exec("INSERT INTO test VALUES (?, ?)", fmt.Sprintf("foo.%d", i), fmt.Sprintf("bar.%d", i))
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("inserting 100000 keys to my sqlite took: %v\n", time.Since(start))
	fmt.Println("stats for sqlite files")
	printFileSizes(dbDir)
}
