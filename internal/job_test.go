package internal

import "testing"

func TestJobSeed(t *testing.T) {
	if (Job{ID: "job-a"}).ID == "" {
		t.Fatal("job id should be retained")
	}
}
