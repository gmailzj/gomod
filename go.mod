module gomod

replace models v0.0.0 => ./models

replace utils v0.0.0 => ./utils

replace utils/db v0.0.0 => ./utils/db

replace utils/uuid v0.0.0 => ./utils/uuid

replace controller => ./controller

require (
	controller v0.0.0
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/jmoiron/sqlx v1.2.0 // indirect
	github.com/json-iterator/go v1.1.5 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.2.2 // indirect
	golang.org/x/net v0.0.0-20181217023233-e147a9138326 // indirect
	golang.org/x/sync v0.0.0-20181108010431-42b317875d0f // indirect
	golang.org/x/sys v0.0.0-20181213200352-4d1cda033e06 // indirect
	google.golang.org/appengine v1.4.0 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	models v0.0.0
	utils v0.0.0
	utils/db v0.0.0 // indirect
	utils/uuid v0.0.0
)
