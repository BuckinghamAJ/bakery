echo "Building..."
templ generate
./tailwindcss -i cmd/web/styles/input.css -o cmd/web/assets/css/output.css
go run cmd/api/main.go
