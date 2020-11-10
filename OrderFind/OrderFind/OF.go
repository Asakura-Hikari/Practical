package OrderFind

func OF(index string)bool  {
	o := false
	db := []string {"艾瑞莉娅","阿卡丽","乐芙兰","奥利安娜"}
	for i:=0;i<len(db);i++ {
		if index == db[i] {
			o = true
			break
		}
	}
	return o
}