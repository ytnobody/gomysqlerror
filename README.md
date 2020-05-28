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

