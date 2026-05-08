# Gator

Gator is a CLI application built with Go that uses PostgreSQL to store and manage data.

## Requirements

Before running the program, make sure you have the following installed:

- Go
- PostgreSQL

## Install Go

Download and install Go from the official website:

[Go Downloads](https://go.dev/dl/?utm_source=chatgpt.com)

To verify the installation:

```bash
go version
```

## Install PostgreSQL

Download PostgreSQL from:

PostgreSQL Downloads

After installation, make sure your PostgreSQL server is running.

You can verify it with:

```bash
psql --version
```

## Install the Gator CLI

Use go install to install the CLI:
```bash
go install github.com/zepetriko/go_aggregator
```

After installation, make sure your Go bin directory is in your PATH.

Then verify the installation:

```bash
gator
```

## Database Setup

Create a PostgreSQL database for the application:

```sql
CREATE DATABASE gator;
```

---

## Configuration File

Create a configuration file named `.gatorconfig.json` in your home directory.

Example:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable"
}
```

Replace:

- `username` with your PostgreSQL username
- `password` with your PostgreSQL password

---

## Running the Program

Once configured, you can use the CLI directly from your terminal.

Example:

```bash
gator register johndoe
```

---

## Useful Commands

### Register a User

```bash
gator register <username>
```

Creates a new user.

### Login

```bash
gator login <username>
```

Logs in as an existing user.

### Add a Feed

```bash
gator addfeed "<feed_name>" "<feed_url>"
```

Example:

```bash
gator addfeed "Hacker News" "https://news.ycombinator.com/rss"
```

### Follow a Feed

```bash
gator follow "<feed_url>"
```

### Browse Posts

```bash
gator browse
```

Displays recent posts from followed feeds.

### Reset the Database

```bash
gator reset
```

Deletes all stored data from the database.

---

## Troubleshooting

### `command not found: gator`

Make sure your Go bin directory is included in your PATH.

### PostgreSQL Connection Errors

Verify that:

- PostgreSQL is running
- Your username and password are correct
- The database exists
- The connection string in `.gatorconfig.json` is valid

---

## Additional Resources

- Go Documentation: https://go.dev/doc/
- PostgreSQL Documentation: https://www.postgresql.org/docs/