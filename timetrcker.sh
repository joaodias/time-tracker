echo "Initializing frontend"
cd frontend
webpack-dev-server --port 4002 &
echo "Initializing db"
cd ../postgres
docker run -p 5432:5432 -e POSTGRES_PASSWORD=dMzpXuzT -e POSTGRES_DB=TimeTracker joaodias/postgres:latest &
echo "Initializing backend"
cd ../backend
go run main.go &
echo "Done!"
