package db
/*

Example usage: 
	
	ctx, _ := context.WithCancel(context.Background())
	dba = &db.Actor{
		DB:         s.db,
		ActionChan: make(chan db.Func),
	}
	
	go dba.Run(ctx)

*/