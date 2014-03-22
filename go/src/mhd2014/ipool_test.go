package main

import (
	"testing"
	"time"
)

func TestIPoolQuery(t *testing.T) {

	t0 := time.Date(2014, 3, 21, 0, 0, 0, 0, time.UTC)
	SearchIPoolArticles(t0, t0.Add(5*time.Hour*24), []string{"www.welt.de", "www.abendblatt.de"})
}
