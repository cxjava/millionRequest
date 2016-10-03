package main

import (
	"fmt"
	"time"
)

// Job represents the job to be run
type Job struct {
	Payload Payload
}

type PayloadCollection struct {
	Payloads []Payload `json:"data" xml:"data" form:"data"`
}

type Payload struct {
	Username string `json:"name" xml:"name" form:"name"`
}

func (p *Payload) UploadToS3() error {
	storage_path := fmt.Sprintf("%v/%v", p.Username, time.Now().UnixNano())
	fmt.Println(storage_path)
	return nil
}
