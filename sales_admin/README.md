# Sales Admin Portal
Welcome to the Acme Cult Sales Admin Portal!

## Setup For Local Development

### Install Postgres
This can be done either with **Docker**, or through a **manual installation**.

With **Docker**, you can use the following command
```
docker run --rm -d  -e POSTGRES_PASSWORD=password \
                    -e POSTGRES_USER=portal-admin \
                    -e POSTGRES_DB=portal \
                    -e LC_ALL=C.UTF-8\
                    -p 5432:5432 \
                    postgres:11-alpine
```

With a **manual installation**, either find the appropriate [installation](https://www.postgresql.org/download/) for your environment

*Mac Users* can use `homebrew` to install postgres: 
```bash
# Install postgres
brew install postgres

# Start the postgres service and automatically set to start on boot
brew services start postgres 
```

### Install Go
Download and install the appropriate Go distribution for your environment [here](https://golang.org/)

*Mac Users*

You can install Go using `homebrew`. i.e. `brew install go`

### Install NodeJS
Download and install the appropriate NodeJS distribution for your environment [here](https://nodejs.org/)

*Mac Users*

You can install NodeJS using `homebrew` and Node Version Manager `nvm`.
```bash
# Install NVM
brew install nvm

# Install NodeJS v11-LTS
nvm install 11 --lts

# Verify that the current version of node being used is correct
nvm current
```

### Install Angular CLI
The Angular CLI can be installed using `npm`. i.e. `npm i -g @angular/cli`

### Environment Variables
The `docker-compose.yml` file contains a set of existing environment variables that can be altered if you wish.

**backend**

The backend expects to have the following variables defined prior to running the API:
```bash
DB_USER: portal-admin           # Database user
DB_PASS: password           # Database user's password
DB_NAME: portal         # Name of database being used on the cluster
DB_HOST: db         # Hostname of the postgres cluster
DB_PORT: 5432           # TCP port being used by the postgres cluster for connections
DB_MAX_LIFETIME: 60         # Maximum db connection lifetime in seconds
DB_MAX_OPEN_CONNECTIONS: 50         # Maximum number of open connections allowed by the application
DB_MAX_IDLE_CONNECTIONS: 5          # Maximum number of idle connections allowed by the applicatio
DB_MAX_ATTEMPTS: 10         # Maxmimum number of retries if initial db connection fails
API_PORT: 9090          # TCP port to be used by the API
LOG_LEVEL: info         # Level of logging detail to be sent to stdout
ACCESS_CONTROL_ALLOW_ORIGINS: "*"           # CORS policy to be used by the API
JWT_SECRET: ZPDDvgd2hnKJzV42cRUB7pmfpUxsXY          # JWT secret used for signing tokens
ADMIN_USER: cultleader22@thecult.com            # Admin user email that can be used for authentication. This user is automatically added by the API
ADMIN_PASS: Sup3r-Cult!st           # Admin user password that can be used for authentication. This user is automatically added by the API
```

**frontend**

The frontend expects to have an `environment.ts` file within `src/environments` for each environment being deployed to.

Currently the only variables that are expected:
```TypeScript
export const environment = {
  production: false,    // [true/false]: A boolean flag that tells Angular that this is the production environment
  apiBaseUrl: 'http://localhost:9090'   // [string]: The base URL of the backend API server. e.g. `http://localhost:9090`
};
```

**postgres**

When starting a postgres cluster for the first time, you can specify a variety of environment variables that will tell postgres to automatically set certain things on init.

```bash
LC_ALL = C.UTF-8            # Default locale to be used
POSTGRES_PASSWORD = password            # If provided, Postgres will create a superuser with the given credentials
POSTGRES_USER = portal-admin            # If provided, Postgres will create a superuser with the given credentials
POSTGRES_DB = portal            # If provided, Postgres will create a database with the given name on init
```

## Deployment Instructions

### Docker
The easiest way to get up and running is make sure you have [Docker](https://docs.docker.com/install/) installed. Once you do, local deployment is as easy as running `docker-compose up`. You can also run it detached from your terminal by adding the `-d` flag.

If your stack is already up and running via `docker-compose up`, just use `docker-compose down` to seamlessly shut down your containers.

### Manual Deployment
If you'd like to deploy without Docker, you may use the following steps

1. Be sure to set ALL of the environment variables described above under `Setup For Local Development` before moving to the next step
2. Ensure that a local instance of Postgres is running `psql -U $DB_USER -h $DB_HOST -d $DB_NAME`. If you do not have Postgres installed, please see the instructions under the `Local Setup For Development` section above
3. Navigate to the backend source code located at `src/api`
4. Compile the backend with `go build -i` and either run the binary in your terminal window or use `nohup`
5. Navigate to the frontend source code located at `src/ui`
6. Either create a build with `ng build --prod --build-optimizer`, or use a live development server with `ng serve --open` and your browser will load a window automatically

### Check Out The [App](http://localhost:4200)

## Testing

### Frontend Testing
Frontend end-to-end tests are run via `protractor`. Unit tests are run via `karma`. 

To run all e2e tests, you may use the command `ng e2e`.

To run all unit tests, you may use the command `ng test`.

### Backend Testing
Go's testing framework is built-in to the language runtime and can be called with `go test`

To differentiate between **unit** and **integration** tests, you can specify a build flag on the first line of your `_test` files.

e.g. `my_file_test.go` 
```Go
// +build integration

package my_package_test
```

The above `_test` file will only run when given the additional build flag for `integration`. e.g. `go test -v --tag integration`

This can be useful if you wish to have unit tests run before integration tests.

**Authentication**

The API uses JWT authentication with Bearer tokens.
 
In order to authenticate, you must send a POST to `/login` with your user's email and password as form values.

If your login attempt is successful you will be returned the following JSON
```javascript
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImN1bHRsZWFkZXIyMkB0aGVjdWx0LmNvbSIsInJvbGUiOiJhZG1pbiIsImV4cCI6MTU2NTEwNzczOSwiaXNzIjoiQWNtZSBDdWx0IEhlcm8gU3VwcGxpZXMgSW5jLiJ9.UmDpUDHmYfnWQJ9mT-lMvPOc-aEDdbkzUxlP4yL3vBI"
}
```

Handling authorization works in the same way as many other JWT authenticated APIs.

In all subsequent requests made to the API after `/login`, you must provide your bearer token in an Authorization header
```
{
    "Authorization": "Bearer {token}",
    ...
}
```


### Database Migrations
We are using an automated database migration framework provided by [mattes](http://github.com/mattes/migrate)

The migrations will run after the API has started and has established a connection to the postgres cluster.

The migration logic is stored in `sql` files within the `api/migrations` directory.
Each migration step is prefixed based on its chronological step in the database's version history and is comprised of an `up` and `down` file.
Depending on the available migration files in your deployed backend, the migrations will move either `up` or `down` as many versions as needed to match the latest migration in `api/migrations`.