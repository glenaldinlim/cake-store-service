package utils

import "database/sql"

func CommitOrRollback(location string, tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		PanicIfError(location, errRollback)
		panic(err)
	} else {
		errCommit := tx.Commit()
		PanicIfError(location, errCommit)
	}
}
