module bookkeeper/modules/database

go 1.22

require (
	bookkeeper/modules/book v0.0.0-00010101000000-000000000000
	gorm.io/gorm v1.25.7
)

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
)

replace bookkeeper/modules/book => ../book
