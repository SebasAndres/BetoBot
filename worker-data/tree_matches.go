package main

func isInLevel1(rdb redis.Client, ctx Context, t1 string){
	return rdb.SIsMember(ctx, PREFIX+"TREE_MATCHES_LEVEL_1", t1).Val()
}

func isChildOf(rdb redis.Client, ctx Context, child string, parent string){
	return 
}


func agregarNuevoHSET(t2 string){

func registerMatchInRedisTree(m *Match, r* mRedis) {

	t1 := m.HomeTeam
	t2 := m.AwayTeam

	if isInLevel1(rdb ctx, t1) && isChildOf(rdb, ctx, t2, t1){
		updateHSET in tree.level_1[m.t1][m.t2]
		return
	}
	else if isInLevel1(rdb, ctx, t2) && isChildOf(rdb, ctx, t1, t2){
		updateHSET in tree.level_1[m.t2][m.t1]
		return
	}
	else if isInLevel1(rdb, ctx, t1){
		tree.level_1[m.t1].agregarNuevoHSET_Level2(m.t2)
	}
	else if isInLevel1(rdb, ctx, t2){
		tree.level_1[m.t2].agregarNuevoHSET_Level2(m.t1) 
	}
	else{
		tree.level_1.agregarNuevoHSET_Level1(m.t1)
		tree.level_1[m.t1].agregarNuevoHSET_Level2(m.t2)
	}
}
