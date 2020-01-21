package main

import (
	"html/template"
	"net/http"
	"strconv"
)

//mani関数
func main() {

	http.HandleFunc("/suc", handleClockTpl)

	http.ListenAndServe(":8080", nil)
}

//クッキーやHTMLCSS等に関する関数宣言
func handleClockTpl(w http.ResponseWriter, r *http.Request) {

	//クッキーがなかったらここで作成
	count, err := r.Cookie("watta")
	if err != nil {
		count = &http.Cookie{Name: "watta", Value: "0", HttpOnly: true}
		http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
	}

	//css参照
	tpl := template.Must(template.ParseFiles("html/suiteki.html"))

	//訪問回数をクッキーから取得
	currentCount, err := strconv.ParseInt(count.Value, 0, 64)
	count.Value = strconv.Itoa(int(currentCount + 1))
	http.SetCookie(w, count)

	//HTMLへのテンプレート
	m := map[string]string{
		"Date": count.Value,
	}

	//HTML再描写
	tpl.Execute(w, m)

}
