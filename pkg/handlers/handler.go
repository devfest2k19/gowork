package handlers

import (
	"github.com/devfest2k19/gowork/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/pickme-go/log"
	"io/ioutil"
	"net/http"
	"sync/atomic"
)

var counter int64
var PersonMap map[int64]models.Person

type HandlePost struct {

}

type PostResponse struct {
	Id int64 `json:"id"`
}

func (HandlePost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, err)
	}

	var p models.Person

	err = json.Unmarshal(data, &p)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, err)
	}

	p.Id = atomic.AddInt64(&counter, 1)
	PersonMap[p.Id] = p

	w.WriteHeader(http.StatusOK)
	pr := PostResponse{
		Id:p.Id,
	}
	b, _ := json.Marshal(pr)
	_, _ = w.Write(b)
}



