module gomod

replace utils v0.0.0 => ./utils

replace utils/uuid v0.0.0 => ./utils/uuid

replace controller => ./controller

require (
	controller v0.0.0
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gin-contrib/sse v0.0.0-20170109093832-22d885f9ecc7 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/json-iterator/go v1.1.5 // indirect
	github.com/mattn/go-isatty v0.0.4 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.2.2 // indirect
	github.com/ugorji/go/codec v0.0.0-20181022190402-e5e69e061d4f // indirect
	golang.org/x/net v0.0.0-20181217023233-e147a9138326 // indirect
	golang.org/x/sync v0.0.0-20181108010431-42b317875d0f // indirect
	golang.org/x/sys v0.0.0-20181213200352-4d1cda033e06 // indirect
	google.golang.org/appengine v1.4.0 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	gopkg.in/yaml.v2 v2.2.1 // indirect
	utils v0.0.0
	utils/uuid v0.0.0
)
