module controller

replace models v0.0.0 => ../models

replace utils v0.0.0 => ../utils

replace utils/db v0.0.0 => ../utils/db

require (
	github.com/gin-contrib/sse v0.0.0-20170109093832-22d885f9ecc7 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/golang/protobuf v1.2.0 // indirect
	github.com/jmoiron/sqlx v1.2.0
	github.com/mattn/go-isatty v0.0.4 // indirect
	github.com/ugorji/go/codec v0.0.0-20181209151446-772ced7fd4c2 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
	models v0.0.0
	utils v0.0.0
	utils/db v0.0.0
)
