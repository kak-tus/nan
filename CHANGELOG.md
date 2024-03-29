2023-03-13 v0.9.0
  - Support for uints.

2022-09-03 v0.8.0
  - Support for pointer (*type) helpers to create nan values.

2022-06-03 v0.7.0
  - Support for pointer (*type) helpers.

2022-02-08 v0.6.1
  - Update all dependencies.

2022-02-08 v0.6.0
- Replace markbates/pkger with native Go embed #11 (creker).
- Minimum go version changed to 1.16 (creker).

2022-01-25 v0.5.0
- Add version command.

2021-09-29 v0.4.0
- Add support for encoding.TextMarshaler and encoding.TextUnmarshaler (creker).

2021-08-14 v0.3.0
- Add support for go-json (creker).

2020-10-07 v0.2.3
- Fixed bug with generating marhallers for custom structs with omitempty.

2020-10-07 v0.2.2
- Include basic helpers in data generation.

2020-10-07 v0.2.1
- Fixed generation for "go generate" commend (exclude our go generate).

2020-10-07 v0.2.0
- Support for generate nullable tuper for custom structs.

2020-10-07 v0.1.12
- Fixed easyjson generation with -all option enabled for custom user-generated types.

2020-08-07 v0.1.11
- Omitempty support for jsoniter.

2020-08-07 v0.1.10
- implement easyjson.Optional

2020-08-07 v0.1.9
- Shorter helpers. Old helpers - deprecated.
- generate helpers

2020-08-06 v0.1.8
- Validator interface

2020-08-05 v0.1.6
- nan command that generates marshalers

2020-08-03 v0.1.5
- Remove unneded sql dependency.
- Add support for sql marhsalling (Scan and Value methods) (creker).
- Some refactoring - files merged by types (creker).

2020-07-31 v0.1.4
- Add int type (creker).
- Ability to marshal and unmarshal mixed sized integers like in qocql (creker).

2020-07-31 v0.1.3
- Add helpers for recently added types.

2020-07-30 v0.1.2
- Add float32, int8, int16, int32 (creker).
- Bugfix with processing null string value (creker).
- Bugfix with processing null/zero time value (creker).

2020-07-25 v0.1.1
  - Add float64.

2020-07-24 v0.1.0
  - First release.
