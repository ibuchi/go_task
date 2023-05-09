# Go Task Application

Go Task Application is a simple command-line application written in Go that demonstrates how to connect to a PostgreSQL database, create a table, insert records, query records, and return the results as JSON.

## Installation

1. Install Go by following the instructions at [https://golang.org/doc/install](https://golang.org/doc/install).
2. Install PostgreSQL by following the instructions at [https://www.postgresql.org/download/](https://www.postgresql.org/download/).
3. Clone or download the repository from GitHub.

## Usage

1. Open a terminal and navigate to the root directory of the project.
2. Run the following command to build the application:

   ```
   go build .
   ```
3. Run the following command to start the postgres server:

   ```
   pg_ctl -D <path-to-data-directory> start
   ```

4. Run the following command to start the application:

   ```
   go run .\task.go
   ```

5. The application will connect to the PostgreSQL database and perform the following operations:
   - Drop the `users` table if it already exists.
   - Create the `users` table.
   - Insert three records into the `users` table.
   - Query the `users` table and return the results as JSON.

6. The JSON result will be printed to the terminal.
