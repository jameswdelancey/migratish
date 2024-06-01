# SQLite3 Schema Management (Migration) CLI Utility

## Overview

This repository hosts a straightforward SQLite3 migration CLI utility developed in Go, compatible with Windows, Linux, or Mac environments. Utilizing a single source file `main.go`, the tool efficiently manages database migrations, enabling both forward and reverse migration operations.

[![Go Build and Release](https://github.com/jameswdelancey/migratish/actions/workflows/build.yml/badge.svg)](https://github.com/jameswdelancey/migratish/actions/workflows/build.yml)

## Features

- **Migration Direction**: Supports forward (f) and reverse (r) migrations.
- **Automated Version Tracking**: Automatically tracks migration versions via a dedicated database table.
- **Simplified CLI Interface**: Accepts two arguments via the command line: the path to the database file and the migration SQL file.

## Usage

### Download Latest Release

Download the latest release binary for your platform from the **Releases** section of this repository. Execute the binary with the required arguments:

```sh
# Linux and MacOS
./migratish -verbose ./db/database.db ./db_migrations/migration.sql

# Windows
./migratish.exe -verbose ./db/database.db ./db_migrations/migration.sql
```

You may also add the binary to your system's PATH for convenience.

### Usage with Compilation

To utilize the migration utility, compile the `main.go` file and execute the resulting binary along with the necessary arguments:

```sh
go build -o migrate main.go
./migrate path/to/database.db path/to/migration.sql
```

Ensure that your migration files follow the naming convention `[f|r][1-9][0-9]*.sql`, where:
- `f` denotes a forward migration.
- `r` denotes a reverse migration.
- The number sequence (1-10 or higher) indicates the migration version.

## Migration Files

Migration files should adhere to the schema `[f|r][1-9][0-9]*.sql`, where:
- `f` represents forward migrations.
- `r` represents reverse migrations.
- The number sequence signifies the migration version.

Example:
- `f1.sql` is the first forward migration.
- `r2.sql` is the reverse migration for version 2 (downgrading from version 2 to 1).

## Migration Table

The utility automatically creates a `migration` table within the SQLite3 database if it doesn't already exist. This table tracks the current migration version, containing datetime and version columns.

## Assumptions

- SQL migration files are correctly formatted and contain valid SQLite3 SQL statements.
- Error checking for SQL statements within migration files is minimal.
- Migration file names are expected to follow the specified format, and the code does not handle versions higher than 9, although it's designed to support any number above 1.

## Dependencies

This utility relies on the `mattn/go-sqlite3` package for SQLite3 database interaction. Ensure this dependency is installed before building the utility:

```sh
go install ./...
```

## Limitations

- Assumes the presence of well-formed and valid SQL migration files.
- Error handling for SQL execution is not extensive and should be enhanced for production use.
- Does not support automated rollback of failed migrations.

## Contributing

Contributions to enhance the utility, including improved error handling, support for complex migration patterns, and overall robustness, are encouraged. Feel free to submit pull requests or open issues to discuss proposed changes.

## License

This SQLite3 Migration CLI Utility is open-source software licensed under the MIT License.
