package main

import (
	"time"
)

//ユーザー情報．名前とIDが格納されている.
//IDはランダム
type UserInfo struct {
	Name string
	Id   int32
}

//コンテンツの内容のデータ
////テンプレートに流し込むときに画像かどうかの判定のためのメンバが存在する．
//
type ContentData struct {
	//text/image
	Type string
	//投稿時の時間 ナノ秒まで格納
	Date time.Time
	//textなら本文,imageならローカルのパスを格納
	Content string
	//投稿者の情報
	Contributer UserInfo
	//画像ならtrue
	IsImage bool
}

//このフォーラムの情報を格納.
//コンテンツ数とコンテンツ本体,フォーラム名が格納されている
type DataSet struct {
	Name        string
	Num         int
	AllContents []ContentData
}

//テンプレートに流し込む情報
//Userはリクエスト送信者の情報である．cookieから取得する．
type TmplData struct {
	Contents DataSet
	User     UserInfo
}
