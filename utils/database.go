package utils

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			Logger().Errorf("DB Tx Rollback: %s", errRollback.Error())
			panic(errRollback)
		}
		panic(err)
	} else {
		errCommit := tx.Commit()
		if errCommit != nil {
			Logger().Errorf("DB Tx Commit: %s", errCommit.Error())
		}
	}
}
