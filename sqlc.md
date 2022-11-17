# sqlc integration

[sqlc](https://github.com/kyleconroy/sqlc) - Generate type-safe code from SQL

By default sqlc uses nullable types from standard library, but nan types have helpers to more easy usage.

To use nan with sqlc add this config to sqlc.yaml and use it as-is replacement.

```
overrides:
  # From https://github.com/kak-tus/nan/sqlc.md
  - db_type: "pg_catalog.int8"
    go_type: "github.com/kak-tus/nan.NullInt64"
    nullable: true
  - db_type: "pg_catalog.serial8"
    go_type: "github.com/kak-tus/nan.NullInt64"
    nullable: true
  - db_type: "pg_catalog.interval"
    go_type: "github.com/kak-tus/nan.NullInt64"
    nullable: true
  - db_type: "pg_catalog.int4"
    go_type: "github.com/kak-tus/nan.NullInt32"
    nullable: true
  - db_type: "pg_catalog.serial4"
    go_type: "github.com/kak-tus/nan.NullInt32"
    nullable: true
  - db_type: "pg_catalog.int2"
    go_type: "github.com/kak-tus/nan.NullInt16"
    nullable: true
  - db_type: "pg_catalog.serial2"
    go_type: "github.com/kak-tus/nan.NullInt16"
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
  - db_type: "pg_catalog.varchar"
    go_type: "github.com/kak-tus/nan.NullString"
    nullable: true
  - db_type: "pg_catalog.bpchar"
    go_type: "github.com/kak-tus/nan.NullString"
    nullable: true
  - db_type: "pg_catalog.numeric"
    go_type: "github.com/kak-tus/nan.NullString"
    nullable: true
  - db_type: "text"
    go_type: "github.com/kak-tus/nan.NullString"
    nullable: true
  - db_type: "pg_catalog.bool"
    go_type: "github.com/kak-tus/nan.NullBool"
    nullable: true
  - db_type: "pg_catalog.float8"
    go_type: "github.com/kak-tus/nan.NullFloat64"
    nullable: true
  - db_type: "pg_catalog.float4"
    go_type: "github.com/kak-tus/nan.NullFloat32"
    nullable: true
```
