## The Time Tracker web app

A time tracker in the web. Deployed to timetracker.surge.sh

### Backend

To run the backend first you need all the dependencies. Go and restore the dependencies.

````
go get github.com/tools/godep
godep restore
````

Then you just run like this:

`go run main.go``

### Frontend

To deploy the Frontend (You need to have `surge`installed):

`sh deploy.sh`

Be aware to change the HOST `const` inside `App.jsx`

## Postgres

first build the container:

`docker build -t joaodias/postgres:latest .`

To fire up a database, run:

`docker run -p 5432:5432 -e POSTGRES_PASSWORD=dMzpXuzT -e POSTGR
ES_DB=TimeTracker joaodias/postgres`