# Changelog

## [Unreleased]

### Changed
- Added `Wait()` call to wait for RethinkDB to be ready before returning client
- Added port exposure in container config
- Added `Port` variable (default: 28015) for the service port
- Added `RethinkUpWaitTime` variable (default: 10 seconds) to configure wait timeout
- Updated default image from `rethinkdb` to `rethinkdb:2`
- Fixed connection address to include port
- Updated to use Go modules
