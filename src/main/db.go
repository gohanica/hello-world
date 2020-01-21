package main

import (
	"time"
)

//データベース操作をここに追加しました．
//これを使わない代替案を付けるのは僕のテストが終わった後にさせてください．
//ファイルシステムの部分が関数化できていればそちらとほぼほぼロジックは同じなはずなので許して．

//ret:none
//コンテンツのデータをデータベースに登録
func Insert_Content(data *ContentData) {
	ins, err := DB.Prepare("INSERT INTO " + Content_tbl_name + "(type,date,content,contributer_id) VALUES(?,?,?,?)")
	ShowErr(err)

	layout := "2006 JST Mon Jan 02 15:04:05"
	date := data.Date.Format(layout)
	_, err = ins.Exec(data.Type, date, data.Content, data.Contributer.Id)
	ShowErr(err)
}

//ret:none
//ユーザー情報をデータベースに登録
func Insert_User(user UserInfo) {

	ins, err := DB.Prepare("INSERT INTO " + User_tbl_name + "(id,name) VALUES(?,?)")
	ShowErr(err)
	_, err = ins.Exec(user.Id, user.Name)
	ShowErr(err)
}

//ret:DataSet
//データベースからコンテンツの情報を取得
func SelectData() DataSet {
	rows, err := DB.Query("SELECT type,date,content,contributer_id FROM " + Content_tbl_name)
	ShowErr(err)
	var ContentsData DataSet
	for rows.Next() {
		str := ""
		layout := "2006 MST Mon Jan 02 15:04:05"
		content := ContentData{}
		err := rows.Scan(&content.Type, &str, &content.Content, &content.Contributer.Id)
		ShowErr(err)

		t, err := time.Parse(layout, str)
		ShowErr(err)

		content.Date = t
		if content.Type == "image" {
			content.IsImage = true
		} else {
			content.IsImage = false
		}

		ContentsData.AllContents = append(ContentsData.AllContents, content)
		ContentsData.Num = len(ContentsData.AllContents)
		ContentsData.Name = "test"

	}
	return ContentsData
}

//ret:UserInfo
//データベースからユーザー情報を取得
func SelectUser(username string) UserInfo {
	rows, err := DB.Query("select name,id from " + DB_name + "." + User_tbl_name + " where name='" + username + "'")
	ShowErr(err)
	u := UserInfo{}
	for rows.Next() {
		rows.Scan(&u.Name, &u.Id)
	}
	return u
}
