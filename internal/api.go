package internal

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"time"
)

type NullInt64 struct {
	sql.NullInt64
}

type NullString struct {
	sql.NullString
}

func (ni NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

func (ni NullInt64) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		ni.Valid = true
		return nil
	}
	return json.Unmarshal(data, &ni.Int64)
}

func (ni NullString) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.String)
}

func (ni NullString) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		ni.Valid = true
		return nil
	}
	return json.Unmarshal(data, &ni.String)
}

type MonitorRec struct {
	Name         string
	StatusCode   NullInt64
	ResponseTime NullInt64
	SSLError     NullString
	SSLExpire    time.Time
	Updated      time.Time
}

type WebsiteInfo struct {
	Id    int
	Url   string
	Nodes map[int]MonitorRec
}

type LatestMonitorInfo struct {
	NodeNames map[int]string
	Websites  []WebsiteInfo
}

type ProbeResult struct {
	WebsiteId    int
	Protocol     int
	StatusCode   NullInt64
	ResponseTime NullInt64
	SSLError     NullString
	SSLExpire    time.Time
}

type Website struct {
	Id  int
	Url string
}

type AllWebsites struct {
	Websites []Website
}
