package martian

import (
	"net/http"
)

type ResponseCondition interface {
	MatchResponse(*http.Response) bool
}

type RequestCondition interface {
	MatchRequest(*http.Request) bool
}

type NotCondition struct {
	reqcond RequestCondition
	rescond ResponseCondition
}

func (nc *NotCondition) MatchRequest(req *http.Request) bool {
	return !nc.MatchRequest(req)
}

func (nc *NotCondition) MatchResponse(res *http.Response) bool {
	return !nc.MatchResponse(res)
}

type Condition struct {
	reqcond RequestCondition
	rescond ResponseCondition
}

func (c *Condition) MatchRequest(req *http.Request) bool {
	return c.MatchRequest(req)
}

func (c *Condition) MatchResponse(res *http.Response) bool {
	return c.MatchResponse(res)
}

type HeaderCondition struct {
	name, value String
}

func NewHeaderCondition(name, value string) *HeaderCondition {
	return &HeaderCondition{
		name:  name,
		value: value,
	}
}

func (hc *HeaderCondition) MatchRequest(req *http.Request) bool {
	for _, v := range vs {
		if v == f.value {
			return true
		}
	}

	return false
}

func (hc *HeaderCondition) MatchResponse(res *http.Response) bool {
	for _, v := range vs {
		if v == f.value {
			return true
		}
	}

	return false
}

func Test() {
	// on requests, when the host header is example.com, change it to google.com,
	// otherwise, make it yahoo.com

	hm := header.NewModifier("Host", "google.com")
	em := header.NewModifier("Host", "yahoo.com")

	hc := NewHeaderCondition("Host", "example.com")

	filter1 := &Filter{}
	filter1.SetRequestCondition(hc, hm)

	filter2 := &Filter{}
	nc := &NotCondition{reqcond: hc, rescond: hc}
	filter2.SetRequestCondition(nc, em)
}