# sqlc integration

[sqlc](https://github.com/kyleconroy/sqlc) - Generate type-safe code from SQL

By default sqlc uses nullable types from standard library, but nan types have helpers to more easy usage.

To use nan with sqlc add this config to sqlc.yaml and use it as-is replacement.

```
overrides:
  # From https://github.com/kak-tus/nan/sqlc.md
  # Types https://github.com/kyleconroy/sqlc/blob/main/internal/codegen/golang/postgresql_type.go#L36
  - db_type: "serial"
    go_type: "github.com/kak-tus/nan.NullInt32"
    nullable: true
  - db_type: "serial4"
    go_type: "github.com/kak-tus/nan.NullInt32"
    nullable: true
  - db_type: "pg_catalog.serial4"
    go_type: "github.com/kak-tus/nan.NullInt32"
    nullable: true
  - db_type: "bigserial"
    go_type: "github.com/kak-tus/nan.NullInt64"
    nullable: true
  - db_type: "serial8"
    go_type: "github.com/kak-tus/nan.NullInt64"
    nullable: true
  - db_type: "pg_catalog.serial8"
    go_type: "github.com/kak-tus/nan.NullInt64"
    nullable: true
  - db_type: "smallserial"
    go_type: "github.com/kak-tus/nan.NullInt16"
    nullable: true
  - db_type: "serial2"
    go_type: "github.com/kak-tus/nan.NullInt16"
    nullable: true
  - db_type: "pg_catalog.serial2"
    go_type: "github.com/kak-tus/nan.NullInt16"
    nullable: true
  - db_type: "integer"
    go_type: "github.com/kak-tus/nan.NullInt32"
    nullable: true
  - db_type: "int"
    go_type: "github.com/kak-tus/nan.NullInt32"
    nullable: true
  - db_type: "int4"
    go_type: "github.com/kak-tus/nan.NullInt32"
    nullable: true
  - db_type: "pg_catalog.int4"
    go_type: "github.com/kak-tus/nan.NullInt32"
    nullable: true
  - db_type: "bigint"
    go_type: "github.com/kak-tus/nan.NullInt64"
    nullable: true
  - db_type: "int8"
    go_type: "github.com/kak-tus/nan.NullInt64"
    nullable: true
  - db_type: "pg_catalog.int8"
    go_type: "github.com/kak-tus/nan.NullInt64"
    nullable: true
  - db_type: "smallint"
    go_type: "github.com/kak-tus/nan.NullInt16"
    nullable: true
  - db_type: "int2"
    go_type: "github.com/kak-tus/nan.NullInt16"
    nullable: true
  - db_type: "pg_catalog.int2"
    go_type: "github.com/kak-tus/nan.NullInt16"
    nullable: true
  - db_type: "float"
    go_type: "github.com/kak-tus/nan.NullFloat64"
    nullable: true
  - db_type: "double precision"
    go_type: "github.com/kak-tus/nan.NullFloat64"
    nullable: true
  - db_type: "float8"
    go_type: "github.com/kak-tus/nan.NullFloat64"
    nullable: true
  - db_type: "pg_catalog.float8"
    go_type: "github.com/kak-tus/nan.NullFloat64"
    nullable: true
  - db_type: "real"
    go_type: "github.com/kak-tus/nan.NullFloat32"
    nullable: true
  - db_type: "float4"
    go_type: "github.com/kak-tus/nan.NullFloat32"
    nullable: true
  - db_type: "pg_catalog.float4"
    go_type: "github.com/kak-tus/nan.NullFloat32"
    nullable: true
  - db_type: "numeric"
    go_type: "github.com/kak-tus/nan.NullString"
    nullable: true
  - db_type: "pg_catalog.numeric"
    go_type: "github.com/kak-tus/nan.NullString"
    nullable: true
  - db_type: "money"
    go_type: "github.com/kak-tus/nan.NullString"
    nullable: true
  - db_type: "boolean"
    go_type: "github.com/kak-tus/nan.NullBool"
    nullable: true
  - db_type: "bool"
    go_type: "github.com/kak-tus/nan.NullBool"
    nullable: true
  - db_type: "pg_catalog.bool"
    go_type: "github.com/kak-tus/nan.NullBool"
    nullable: true
  # json
  # jsonb
  # bytea
  # blob
  # pg_catalog.bytea
  - db_type: "date"
    go_type: "github.com/kak-tus/nan.NullTime"
    nullable: true
  - db_type: "pg_catalog.time"
    go_type: "github.com/kak-tus/nan.NullTime"
    nullable: true
  - db_type: "pg_catalog.timetz"
    go_type: "github.com/kak-tus/nan.NullTime"
    nullable: true
  - db_type: "pg_catalog.timestamp"
    go_type: "github.com/kak-tus/nan.NullTime"
    nullable: true
  - db_type: "pg_catalog.timestamptz"
    go_type: "github.com/kak-tus/nan.NullTime"
    nullable: true
  - db_type: "timestamptz"
    go_type: "github.com/kak-tus/nan.NullTime"
    nullable: true
  - db_type: "text"
    go_type: "github.com/kak-tus/nan.NullString"
    nullable: true
  - db_type: "pg_catalog.varchar"
    go_type: "github.com/kak-tus/nan.NullString"
    nullable: true
  - db_type: "pg_catalog.bpchar"
    go_type: "github.com/kak-tus/nan.NullString"
    nullable: true
  - db_type: "string"
    go_type: "github.com/kak-tus/nan.NullString"
    nullable: true
  - db_type: "citext"
    go_type: "github.com/kak-tus/nan.NullString"
    nullable: true
  # uuid
  # inet
  # cidr
  # macaddr
  # macaddr8
  - db_type: "ltree"
    go_type: "github.com/kak-tus/nan.NullString"
    nullable: true
  - db_type: "lquery"
    go_type: "github.com/kak-tus/nan.NullString"
    nullable: true
  - db_type: "ltxtquery"
    go_type: "github.com/kak-tus/nan.NullString"
    nullable: true
  - db_type: "interval"
    go_type: "github.com/kak-tus/nan.NullInt64"
    nullable: true
  - db_type: "pg_catalog.interval"
    go_type: "github.com/kak-tus/nan.NullInt64"
    nullable: true
  # daterange
  # datemultirange
  # tsrange
  # tsmultirange
  # tstzrange
  # tstzmultirange
  # numrange
  # nummultirange
  # int4range
  # int4multirange
  # int8range
  # int8multirange
  # hstore
  # bit
  # varbit
  # pg_catalog.bit
  # pg_catalog.varbit
  # box
  # cid
  # oid
  # tid
  # circle
  # line
  # lseg
  # path
  # point
  # polygon
  # void
  # any
```
