# SQLite3 Migration CLI Utility
For Windows, Linux, or Mac

[![Go Build and Release](https://github.com/jameswdelancey/migratish/actions/workflows/build.yml/badge.svg)](https://github.com/jameswdelancey/migratish/actions/workflows/build.yml)

## Overview
This repository contains a simple Go-based SQLite3 migration CLI utility 
designed to manage database migrations. The utility is implemented in a 
single source file `main.go` and facilitates applying forward or reverse 
migrations to a SQLite3 database. 

## Features
- Support for forward (f) and reverse (r) migrations.
- Automated tracking of migration versions through a dedicated table 
  within the database. 
- Command-line interface accepting two arguments: the path to the 
  database file and the path to the migration SQL file. 

## Usage
To use the migration utility, compile the `main.go` file and execute the 
binary with the required arguments. 

```sh
go build -o migrate main.go
./migrate path/to/database.db path/to/migration.sql
```

The migration file name must follow the pattern `[f|r][1-9]+.sql`, where:
- `f` indicates a forward migration.
- `r` indicates a reverse migration.
- The number sequence (1-9 or higher) represents the migration version. 

## Migration Files
Migration files should be named according to the schema `[f|r][1-9]+.sql`:
- `f` for forward migrations.
- `r` for reverse migrations.
- A sequence number indicating the migration version.

For example:
- `f1.sql` would be the first forward migration.
- `r2.sql` would be the reverse migration for version 2 (downgrading from version 2 to 1).

## Migration Table
The utility will automatically create a `migration` table in the SQLite3 
database if it does not exist. This table tracks the current migration 
version of the database, containing a datetime column and a version 
column. 

## Assumptions
- The SQL migration files are correctly formatted and contain valid 
  SQLite3 SQL statements. 
- Error checking for SQL statements within the migration files is 
  minimal. 
- Migration file names are expected to follow the specified format and 
  the code does not handle versions higher than 9, though it is written to 
  support any number above 1. 

## Dependencies
This utility uses the `mattn/go-sqlite3` package for interacting with 
the SQLite3 database. Ensure that this dependency is installed before 
building the utility. 

```sh
go get github.com/mattn/go-sqlite3
```

## Limitations
- The utility assumes the presence of well-formed and valid SQL 
migration files. 
- Error handling for SQL execution is not extensive and should be 
improved for production use. 
- The utility does not support automated rollback of failed migrations. 

## Contributing
Contributions to improve the utility, including error handling, support 
for more complex migration patterns, and robustness, are welcome. Please 
submit a pull request or open an issue to discuss proposed changes. 

## License
This SQLite3 Migration CLI Utility is open-source software licensed 
under the MIT license. 
