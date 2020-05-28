# gomysqlerror

A MySQL Error checker for golang

## SYNOPSIS

```

package main

import (
    database/sql

    _ "github.com/go-sql-driver/mysql"
    dberr "github.com/ytnobody/gomysqlerror/v57error"  // <- v57error is only for MySQL 5.7
)

type MyApp struct {
    DB *sql.DB
    ...
}

...

func (a MyApp) registerUser(email string, name string) error {
    db := a.DB
    _, err := db.Exec(`INSERT INTO user (email, name) VALUES (?, ?)`, email, name)
    if err != nil {
        if dberr.IsServerErrorDupEntry(err) {
            customErr := fmt.Errorf("specified email address is already exists")
            return customErr
        }
        ...
    }
    return nil
} 

```

## Compatible MySQL Versions

* 5.6 -> use `github.com/ytnobody/gomysqlerror/v56error`
* 5.7 -> use `github.com/ytnobody/gomysqlerror/v57error`
* 8.0 -> use `github.com/ytnobody/gomysqlerror/v80error`
