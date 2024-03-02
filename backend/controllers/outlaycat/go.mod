module bookkeeper/controllers/outlaycat

go 1.22

require (
	bookkeeper/modules/book v0.0.0-00010101000000-000000000000
	bookkeeper/modules/database v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.6.0
	github.com/gorilla/mux v1.8.1
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	gorm.io/gorm v1.25.7 // indirect
)

replace bookkeeper/modules/book => ../../modules/book

replace bookkeeper/modules/database => ../../modules/database
