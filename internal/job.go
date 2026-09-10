package internal

type Job struct {
	ID      string `json:"job_id"`
	Project string `json:"project"`
	Shots   int    `json:"shots"`
}
