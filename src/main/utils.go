package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
	//ID生成用
	"math/rand"
)

//ret:*Cookie
//UserInfoからクッキーを生成する.クッキーはポインタ．セッションkっキーである．
func UserInfo_to_Cookie(u UserInfo) *http.Cookie {
	s := strconv.Itoa(int(u.Id))
	c := http.Cookie{
		Name:  CookieName,
		Value: u.Name + "," + s,
	}
	return &c

}

//ret:*UserInfo
//クッキーからUserInfoを取り出す．返値はポインタである．
func Cookie_to_UserInfo(c *http.Cookie) *UserInfo {
	info_slice := strings.Split(c.Value, ",")

	i, err := strconv.Atoi(info_slice[1])
	ShowErr(err)

	return &UserInfo{
		Name: info_slice[0],
		Id:   int32(i),
	}

}

//ret:none
//エラーが非nilなら表示する．それだけ．エラー処理さぼってごめんなさい．
func ShowErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

//ret:string
//store at local directory
//保存した場所のパスを返す
func StoreContent(r *http.Request, date time.Time, user *UserInfo) string {
	//Todo　複数ファイルのパース
	//r.ParseMultipartForm(32<<20)
	//fileHeader:=r.MultipartForm.File["posted_file"]
	fileContent, fileHeader, err := r.FormFile(Tmpl_post_file)
	ShowErr(err)

	data, err := ioutil.ReadAll(fileContent)
	ShowErr(err)

	posted_filename := fileHeader.Filename
	local_filename := SetFileName(*user, date, posted_filename)
	err = ioutil.WriteFile(local_filename, []byte(data), 0600)
	ShowErr(err)

	return local_filename
}

//ret:string
//コンテンツのデータから画像ファイル保存する場所を決定する．もっと細かいディレクトリ構造にしたかった．．．
func SetFileName(u UserInfo, t time.Time, filename string) string {
	userIdstr := strconv.Itoa(int(u.Id))
	fname_slice := strings.Split(filename, "\\")
	//時間の情報をフォーマット
	layout := "2006_01_02_15_04_05"
	t_format := t.Format(layout)

	return Image_FolderPath + u.Name + "_" + userIdstr + "_" + t_format + "_" + fname_slice[len(fname_slice)-1]
}

//ret:UserInfo,bool
//  ユーザーを作成する．その"名前"のユーザーが既にある(第2返値がfalse)ならGet_User(username)[db.go参照]を使用してそのユーザー情報を取得する
func CreateUser(name string) (UserInfo, bool) {
	//db.goにいれるのさぼりましたごめんちょ
	rows, err := DB.Query("SELECT * FROM "+User_tbl_name+" WHERE name = ?", name)
	ShowErr(err)
	//データがあるならrows.Next()がtrueを返してくれる．
	if rows.Next() {
		return UserInfo{}, false
	}
	//いないなら作成
	rand.Seed(time.Now().UnixNano())
	u := UserInfo{
		Name: name,
		Id:   rand.Int31(),
	}
	return u, true
}

//なんか使い道考えてください．（適当）（丸投げ）
/*
func Get_Img_loc(data ContentData) string {
	location := SetFileName(data.Contributer, data.Date, data.Content)
	return location
}
*/
