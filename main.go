package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"seqUtil"
)

type JsonResult struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Seq     uint64 `json:"seq"`
}

type SeqServer struct {
}

func OutputJson(w http.ResponseWriter, r *http.Request, Seq string, Status int) {
	//	out := &JsonResult{}
}

func (server *SeqServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/getSeq" {
		generateSeq(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func generateSeq(w http.ResponseWriter, r *http.Request) {
	result := &JsonResult{}
	r.ParseForm()
	fmt.Println(r.Form["userid"])
	if _, ok := r.Form["userid"]; ok {
	} else {
		result.Message = "userid为空"
		result.Status = 1
		b, _ := json.Marshal(result)
		w.Write(b)
		return
	}

	if _, ok := r.Form["appid"]; ok {
	} else {
		result.Message = "appid为空"
		result.Status = 2
		b, _ := json.Marshal(result)
		w.Write(b)
		return
	}

	// 增加curSeq
	uid := r.Form["userid"][0] + r.Form["appid"][0]
	curseq := curSeq[uid] + 1

	//更新curSeq的map
	curSeq[uid] += 1

	// 计算出属于哪个号段
	section := seqUtil.BKDRHash(uid) % sectionNumber
	fmt.Println("section: ", section)
	// 计算出属于哪个号段
	sectionMaxSeq := maxSeq[section]

	fmt.Println("sectionMaxSeq:", sectionMaxSeq)
	if curseq > sectionMaxSeq {
		maxSeq[section] += step
		db.Query("insert into max_seq (secid, maxSeq) values (?,?); ", section, maxSeq[section])
	}

	result.Message = "ok"
	result.Status = 0
	result.Seq = curseq
	b, _ := json.Marshal(result)
	w.Write(b)

}

var (
	curSeq        map[string]uint64
	maxSeq        map[uint64]uint64
	sectionNumber uint64 = 100
	step          uint64 = 10
	db            *sql.DB
)

func main() {
	db, _ = sql.Open("mysql", "root:zk@/seq")

	curSeq = make(map[string]uint64)
	maxSeq = make(map[uint64]uint64)
	fmt.Println("Start SeqServer ----------------")
	server := &SeqServer{}
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
