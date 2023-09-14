package doh

import "errors"

const DOH_ENDPOINT string = "https://cloudflare-dns.com/dns-query"

type Record string

func (r Record) String() string {
	return string(r)
}

var ErrEmptyResponse = errors.New("empty response")

const (
	RECORD_A     Record = "A"
	RECORD_AAAA  Record = "AAAA"
	RECORD_CNAME Record = "CNAME"
	RECORD_MX    Record = "MX"
	RECORD_NS    Record = "NS"
	RECORD_PTR   Record = "PTR"
	RECORD_SOA   Record = "SOA"
	RECORD_TXT   Record = "TXT"
)

type DOHQuestion struct {
	Name string `json:"name"`
	Type int    `json:"type"`
}

type DOHAnswer struct {
	Name string `json:"name"`
	Type int    `json:"type"`
	TTL  int    `json:"TTL"`
	Data string `json:"data"`
}

type DOHResponse struct {
	Status   int           `json:"Status"`
	TC       bool          `json:"TC"`
	RD       bool          `json:"RD"`
	RA       bool          `json:"RA"`
	AD       bool          `json:"AD"`
	CD       bool          `json:"CD"`
	Question []DOHQuestion `json:"Question"`
	Answer   []DOHAnswer   `json:"Answer"`
}

type DOHError struct {
	Error string `json:"error"`
}
