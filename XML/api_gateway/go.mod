module gateway/module

go 1.18

replace common/module => ./../common

require (
	common/module v0.0.0-00010101000000-000000000000
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/google/uuid v1.1.2
	github.com/gorilla/handlers v1.5.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.5.0
	golang.org/x/crypto v0.0.0-20220518034528-6f7dac969898
	google.golang.org/grpc v1.46.2
	gopkg.in/go-playground/validator.v9 v9.31.0
	gorm.io/driver/postgres v1.3.6
	gorm.io/gorm v1.23.5
)

require (
	github.com/eknkc/basex v1.0.0 // indirect
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/hako/branca v0.0.0-20200807062402-6052ac720505 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.12.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.11.0 // indirect
	github.com/jackc/pgx/v4 v4.16.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/juranki/branca v0.0.0-20181210182342-7f16f3130a36 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/trycourier/courier-go/v2 v2.5.0 // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220519153652-3a47de7e79bd // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
)
