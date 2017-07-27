# Boilerplate to build Go web applications

1. Starts a HTTP server at port 80
2. Autoloads controllers and their methods. see `site/routes.go`
3. Loads all templates in the specified template directory.
4. Serves static files on `/assets/filePath`. Place all static files in the folder specified.

## Usage

1. Clone the repository
2. Create a settings file `settings.go` using `settings.go.example` as a guide.
3. Specify the path to your assets folder
4. Specify your templates directory
5. Install using `go install` then run `site`(or another package name if your change it)


### Using Docker (Optional)

1. Create a start.sh
