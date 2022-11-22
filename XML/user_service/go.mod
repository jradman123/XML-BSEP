module user/module

go 1.18

replace (
	common/module => ../common
	monitoring/module => ../monitoring
)

require (
	common/module v0.0.0-00010101000000-000000000000
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.6.0
	github.com/google/uuid v1.3.0
	github.com/lib/pq v1.10.6
	github.com/mattevans/pwned-passwords v0.5.0
	github.com/microcosm-cc/bluemonday v1.0.18
	github.com/opentracing/opentracing-go v1.2.0
	github.com/sirupsen/logrus v1.8.1
	github.com/trycourier/courier-go/v2 v2.5.0
	golang.org/x/crypto v0.0.0-20220315160706-3147a52a75dd
	google.golang.org/grpc v1.46.2
	gopkg.in/go-playground/validator.v9 v9.31.0
	gorm.io/driver/postgres v1.3.5
	gorm.io/gorm v1.23.5
	monitoring/module v0.0.0-00010101000000-000000000000
)

require (
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gorilla/css v1.0.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.9.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.12.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.11.0 // indirect
	github.com/jackc/pgx/v4 v4.16.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/klauspost/compress v1.15.6 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/minio/highwayhash v1.0.2 // indirect
	github.com/nats-io/jwt/v2 v2.2.0 // indirect
	github.com/nats-io/nats.go v1.16.0 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/snowzach/rotatefilehook v0.0.0-20220211133110-53752135082d // indirect
	github.com/tamararankovic/microservices_demo/common v0.0.0-20220326142530-97bfd7810e53 // indirect
	github.com/uber/jaeger-client-go v2.25.0+incompatible // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.mongodb.org/mongo-driver v1.9.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20220111092808-5a964db01320 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20220609170525-579cf78fd858 // indirect
	google.golang.org/genproto v0.0.0-20220519153652-3a47de7e79bd // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)
