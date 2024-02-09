package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	message := recover()
	if message != nil {
		errRollback := tx.Rollback()
		PanicIfError(errRollback)
		panic(message)
	} else {
		errCommit := tx.Commit()
		PanicIfError(errCommit)
	}
}
