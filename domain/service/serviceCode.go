package service

// 在这个里创建出函数，并编号

// 入库
func LoadService() error {
	return nil
}

// 上班打卡提醒
func server01() error {
	// 遍历所有订阅了该信息的人
	//ctx := context.Background()
	//auids, err := Aggregates.R.SubscribeAgg.Repo.FindUserInfo(ctx, 1)
	//if err != nil {
	//	log.Errorf("1号上班打卡服务异常 %s", err)
	//	return err
	//}
	// 遍历这个auid,给所有人都发消息
	return nil
}

// 上班打卡提醒
func server02() error {
	return nil
}
