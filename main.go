package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const migrationTableCreationQuery = `
CREATE TABLE IF NOT EXISTS migrations (
    migration_version INTEGER PRIMARY KEY,
    migration_date DATETIME DEFAULT CURRENT_TIMESTAMP
);
`

func main() {
	// Define a verbose flag with default value false
	verbose := flag.Bool("verbose", false, "Enable verbose output")
	flag.Parse()

	// Check for the correct number of arguments after parsing
	if flag.NArg() != 2 {
		fmt.Println("Usage: migration [-verbose] <database_file> <migration_file>")
		os.Exit(1)
	}

	// Get the non-flag command-line arguments
	dbFile := flag.Arg(0)
	migrationFile := flag.Arg(1)

	// If verbose flag is set, print the file names
	if *verbose {
		fmt.Printf("Database file: %s\n", dbFile)
		fmt.Printf("Migration file: %s\n", migrationFile)
	}

	// Open the database
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		fmt.Printf("Error opening database: %s\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// Ensure the migrations table exists
	_, err = db.Exec(migrationTableCreationQuery)
	if err != nil {
		fmt.Printf("Error creating migrations table: %s\n", err)
		os.Exit(1)
	}

	// Parse migration file name
	fileName := filepath.Base(migrationFile)
	if len(fileName) < 5 || (fileName[0] != 'f' && fileName[0] != 'r') || fileName[len(fileName)-4:] != ".sql" {
		fmt.Println("Invalid migration file name. Must match [f|r][1-9]+.sql")
		os.Exit(1)
	}

	direction := fileName[0]
	migVersion, err := strconv.Atoi(strings.TrimSuffix(fileName[1:len(fileName)-4], ".sql"))
	if err != nil {
		fmt.Printf("Error parsing version number: %s\n", err)
		os.Exit(1)
	}
	if *verbose {
		fmt.Printf("Direction: %c\n", direction)
		fmt.Printf("Migration file version: %d\n", migVersion)
	}

	// Get current version from the database
	var dbResult sql.NullInt64 // Use sql.NullInt64 to handle NULL values
	err = db.QueryRow("SELECT MAX(migration_version) FROM migrations").Scan(&dbResult)
	if err != nil {
		fmt.Printf("Error querying current migration version: %s\n", err)
		os.Exit(1)
	}

	// Check if dbResult is NULL
	var dbVersion int
	if dbResult.Valid {
		// If dbResult is not NULL, use the value
		dbVersion = int(dbResult.Int64)
	} else {
		// If dbResult is NULL, you can set a default value or handle it as needed
		dbVersion = 0
	}
	if *verbose {
		fmt.Printf("Db migration version: %d\n", dbVersion)
	}

	// Check if the migration can be applied
	if direction == 'f' && migVersion-1 != dbVersion {
		fmt.Println("Invalid forward migration version. Must be one greater than the current database version.")
		os.Exit(1)
	} else if direction == 'r' {
		if migVersion == 0 || migVersion != dbVersion {
			fmt.Println("Invalid reverse migration version. Must be equal to the current database version.")
			os.Exit(1)
		}
	}

	// Read migration file content
	migrationContent, err := ioutil.ReadFile(migrationFile)
	if err != nil {
		fmt.Printf("Error reading migration file: %s\n", err)
		os.Exit(1)
	}

	// Apply migration
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("Error beginning transaction: %s\n", err)
		os.Exit(1)
	}

	_, err = tx.Exec(string(migrationContent))
	if err != nil {
		tx.Rollback()
		fmt.Printf("Error executing migration: %s\n", err)
		os.Exit(1)
	}

	// Update migration migVersion
	if direction == 'f' {
		_, err = tx.Exec("INSERT INTO migrations (migration_version) VALUES (?)", migVersion)
	} else {
		_, err = tx.Exec("DELETE FROM migrations WHERE migration_version = ?", migVersion)
	}
	if err != nil {
		tx.Rollback()
		fmt.Printf("Error updating migration version: %s\n", err)
		os.Exit(1)
	}

	err = tx.Commit()
	if err != nil {
		fmt.Printf("Error committing transaction: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Migration applied successfully.")
}
