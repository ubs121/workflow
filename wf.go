// Copyright 2013 ubs121
package comm

import (
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"lerp/db"
	"memex/rpc"
	"net/http"
)

type (
	WfArgs struct {
		Id       string ",omitempty"
		WfId     string ",omitempty"
		Action   string ",omitempty"
		Data     bson.M ",omitempty"
		FromNode string ",omitempty"

		Collection string ",omitempty"
		ObjectId   string ",omitempty"
	}
)

// процесс ажиллуулах
func Run(w http.ResponseWriter, r *http.Request) {
	args := WfArgs{}
	rpc.ReadJson(r, &args)

	err, runResult := db.ExecJS(args.WfId, args)

	rpc.WriteJson(r, w, runResult, err)
}

// э-мэйл илгээх
func SendEmail(w http.ResponseWriter, r *http.Request) {
	var err error
	m := Message{}
	rpc.ReadJson(r, &m)

	err = _sendEmail(&m)

	rpc.WriteJson(r, w, "OK", err)
}

func Register() *mux.Router {
	UrlPrefix := "/wf"
	// create router
	rtr := mux.NewRouter()
	s := rtr.PathPrefix(UrlPrefix).Subrouter()
	s.HandleFunc("/process", Run)
	s.HandleFunc("/email", SendEmail)

	// чөлөөний процесс
	s.HandleFunc("/leave/start/{id}", LeaveStart)
	s.HandleFunc("/leave/approve/{id}", LeaveApprove)
	s.HandleFunc("/leave/reject/{id}", LeaveReject)

	return rtr
}
