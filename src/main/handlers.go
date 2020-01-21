package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"
)

//投稿時のハンドラ関数画像かテキストかを判定しそれぞれに合った挙動をする．最後にフォーラムへリダイレクトすることで更新してくれる．
func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.NotFound(w, r)
		return
	}
	c, err := r.Cookie(CookieName)
	if err != nil {
		fmt.Fprintln(w, "cannot get user info")
	}
	userInfo := Cookie_to_UserInfo(c)

	date := time.Now()
	contentType := r.Header["Content-Type"]
	contentdata := ContentData{}
	contype := strings.Split(contentType[0], ";")
	switch contype[0] {
	case "application/x-www-form-urlencoded":
		content := r.PostFormValue(Tmpl_post_text)
		contentdata = ContentData{
			Type:        "text",
			Date:        date,
			Content:     content,
			Contributer: *userInfo,
			IsImage:     false,
		}
	case "multipart/form-data":
		local_filename := StoreContent(r, date, userInfo)
		contentdata = ContentData{
			Type:        "image",
			Date:        date,
			Content:     local_filename, //file are stored at StoreContent()
			Contributer: *userInfo,
			IsImage:     true,
		}

	}

	Insert_Content(&contentdata)

	//リダイレクト
	http.Redirect(w, r, HTTP_localhost+Port+Forum_URL, 301)

}

//最初にフォーラムを訪れた際はここに飛ぶ．ユーザー情報を初期化してくれる
func Welcome(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue(Tmpl_user_name)
	var User UserInfo
	//このハンドラを実行したときにユーザ情報がないなら初期値を設定
	if username == "" {
		User = UserInfo{
			Name: "plz insert your name in the box!(eng only)",
			Id:   0,
		}
	} else {
		var unique_check bool
		//送られてきたユーザ情報がすでに登録済みか確認_名前が登録済みなら:unique_check=false
		User, unique_check = CreateUser(username)
		if unique_check == true {
			//初めてなら作成
			Insert_User(User)
		} else {
			//そうでないなら情報をDBから取得
			User = SelectUser(username)
		}
	}
	//クッキーをセット
	cookie := UserInfo_to_Cookie(User)
	http.SetCookie(w, cookie)
	http.Redirect(w, r, HTTP_localhost+Port+Forum_URL, 301)
}

//データ情報がないなら/welcomeへリダイレクトする．クッキーが生成されていない場合も同様
func Forum(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(CookieName)
	//クッキーがないなら作成
	if err != nil {
		ShowErr(err)
		http.Redirect(w, r, HTTP_localhost+Port+Welcome_URL, 301)
		return
	}
	//ユーザ情報がないなら作成
	if cookie.Value == ",0" {
		http.Redirect(w, r, HTTP_localhost+Port+Welcome_URL, 301)
		return
	}

	//DBからデータを持ってくる．
	contents := SelectData()
	userinfo := Cookie_to_UserInfo(cookie)
	tmpldata := TmplData{
		Contents: contents,
		User:     *userinfo,
	}
	//テンプレートを実行
	t := template.New(TmplFileName)
	t = template.Must(t.ParseFiles(Html_FolderPath + TmplFileName))
	t.Execute(w, tmpldata)

}
