module bookkeeper/controllers

go 1.22

require (
	bookkeeper/controllers/outlay v0.0.0-00010101000000-000000000000
	bookkeeper/controllers/outlaycat v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.1
)

require (
	bookkeeper/modules/book v0.0.0-00010101000000-000000000000 // indirect
	bookkeeper/modules/database v0.0.0-00010101000000-000000000000 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	gorm.io/gorm v1.25.7 // indirect
)

replace bookkeeper/controllers/outlay => ./outlay

replace bookkeeper/controllers/outlaycat => ./outlaycat

replace bookkeeper/modules/book => ../modules/book

replace bookkeeper/modules/database => ../modules/database
