package database

/*

Example usage:

	ctx, _ := context.WithCancel(context.Background())
	dba = &db.Actor{
		DB:         s.db,
		ActionChan: make(chan database.Func),
	}

	go dba.Run(ctx)

*/
