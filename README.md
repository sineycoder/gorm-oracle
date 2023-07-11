# GORM Oracle Driver
GORM Oracle driver for connect Oracle DB and Manager Oracle DB, 
Based on [CengSin/oracle](https://github.com/CengSin/oracle) 
and [sijms/go-ora](https://github.com/sijms/go-ora)(pure go oracle client), 
not recommended for use in a production environment.

> by the way: fix some bugs in [CengSin/oracle](https://github.com/CengSin/oracle) version,
> support **batch insert** / **conflict insert**.
> 
> but not support returning when using batch insert/update/delete.

# Required dependency
- Oracle 12c+
- Golang 1.18+
- gorm 1.25+

# Quick Start
## How to install
```go
go get -d github.com/sineycoder/gorm-oracle
```
## How to Use
```go
package main

import (
	oracle "github.com/sineycoder/gorm-oracle"
	"gorm.io/gorm"
)

func main() {
	// oracle://user:password@127.0.0.1:1521/service
	url := oracle.BuildUrl("127.0.0.1", "1521", "service", "user", "password", nil)
	db, err := gorm.Open(oracle.Open(url), &gorm.Config{})
	if err != nil {
		// panic error or log error info
	}

	// do somethings
}
```