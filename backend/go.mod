module bookkeeper-backend

go 1.22

require (
	bookkeeper/controllers v0.0.0-00010101000000-000000000000
	bookkeeper/modules/database v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.1
	github.com/mattn/go-sqlite3 v1.14.22
	gorm.io/driver/sqlite v1.5.4
	gorm.io/gorm v1.25.7
)

require (
	bookkeeper/controllers/outlay v0.0.0-00010101000000-000000000000 // indirect
	bookkeeper/controllers/outlaycat v0.0.0-00010101000000-000000000000 // indirect
	bookkeeper/modules/book v0.0.0-00010101000000-000000000000 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
)

replace bookkeeper/modules/book => ./modules/book

replace bookkeeper/modules/database => ./modules/database

replace bookkeeper/controllers => ./controllers

replace bookkeeper/controllers/outlay => ./controllers/outlay

replace bookkeeper/controllers/outlaycat => ./controllers/outlaycat
