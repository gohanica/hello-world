package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
)

//
//
//!!!!!!!!!!!!!!!!!!!!!!!!!  plz read readme.txt  !!!!!!!!!!!!!!!!!!!!!!!!!!!!!
//
//

//定数たち
//ローカルに移動したときに変えてください
//ここを変えれば良いことにしたけど，さすがに見づらいかな？そしたらごめんなさい
const (
	//開くときのポート番号:25565
	//変更時htmlも変更
	Port = "25565"
	//データベース名:webapp
	DB_name = "webapp"
	//ユーザー情報を扱うテーブル名:user
	User_tbl_name = "user"
	//コンテンツの情報を扱うテーブル名:contents
	Content_tbl_name = "contents"
	//クッキーの名前:userinfo
	CookieName = "userinfo"
	//テンプレートファイルの名前:tmpl.html
	TmplFileName = "tmpl.html"

	//CSSフォルダのパス:css/
	//変更時htmlも変更
	Css_FolderPath = "css/"
	//HTMLフォルダのパス:html/
	Html_FolderPath = "html/"
	//画像フォルダのパス:img/

	Image_FolderPath = "img/"

	//htmlにおけるファイルを投稿するフォームの名前:posted_file
	//変更時htmlも変更
	Tmpl_post_file = "posted_file"
	//htmlにおける文章を投稿するフォームの名前:posted_text
	//変更時htmlも変更
	Tmpl_post_text = "posted_text"
	//htmlにおけるユーザー情報を投稿するフォームの名前:username
	//変更時htmlも変更
	Tmpl_user_name = "username"

	//postするときのURL:/post
	//変更時htmlも変更
	Post_URL = "/post"
	//ユーザー登録するときのURL:/welcome
	//変更時htmlも変更
	Welcome_URL = "/welcome"
	//フォーラム本体のURL:/forum
	//変更時htmlも変更
	Forum_URL = "/forum"

	//全ての源:http://localhost
	//変更時htmlも変更
	HTTP_localhost = "http://localhost:"
)

var (
	//DBへのポインタ
	DB *sql.DB
)

func init() {
	//gohanica_client
	sql_username := "gohanica_client"
	//gohanica_client
	sql_password := "gohanica_client"
	//tcp(27.95.148.210)
	sql_ip_address := "tcp(27.95.148.210)"

	var err error
	//DBに接続Go本とは記法が少し違うものの性質は全く同じ
	DB, err = sql.Open("mysql", sql_username+":"+sql_password+"@"+sql_ip_address+"/"+DB_name)
	ShowErr(err)
}

func main() {
	//cssファイルないなら作成(今回は使用せず)
	err := os.Mkdir(Css_FolderPath, 0600)
	ShowErr(err)
	//htmlファイルないなら作成（これはいらなかったかも？）
	err = os.Mkdir(Html_FolderPath, 0600)
	ShowErr(err)
	//imgファイルないなら作成
	err = os.Mkdir(Image_FolderPath, 0600)
	ShowErr(err)
	//css,html,imgファイルを全てハンドラとして登録
	cssFiles := http.FileServer(http.Dir(Css_FolderPath))
	http.Handle("/css/", http.StripPrefix("/css/", cssFiles))

	htmlFiles := http.FileServer(http.Dir(Html_FolderPath))
	http.Handle("/html/", http.StripPrefix("/html/", htmlFiles))

	imgFiles := http.FileServer(http.Dir(Image_FolderPath))
	http.Handle("/img/", http.StripPrefix("/img/", imgFiles))

	//handlers.goを参照のこと．
	http.HandleFunc(Post_URL, Post)
	http.HandleFunc(Welcome_URL, Welcome)
	http.HandleFunc(Forum_URL, Forum)

	//いざ開始
	http.ListenAndServe(":"+Port, nil)

}
