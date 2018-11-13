module gomod

replace utils/demo v0.0.0 => ./utils/demo

replace utils/uuid v0.0.0 => ./utils/uuid

require (
	github.com/eiblog/utils v0.0.0-20180918123929-fcdc03d4d492
	github.com/google/uuid v1.0.0
	utils/uuid v0.0.0
)
