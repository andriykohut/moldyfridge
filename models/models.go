// Package models provides provides models for database acces
package models

import (
	"fmt"
	"strings"
	"time"
)

type Food struct {
	Name   string
	Amount int
	Added  int64
}

func (f *Food) ToString() string {
	return fmt.Sprintf("%s: %d, age - %s", f.Name, f.Amount, f.StringAge())
}

func (f *Food) Age() int64 {
	return time.Now().Unix() - f.Added
}

func (f *Food) StringAge() string {
	duration := f.Age()
	age := ""
	days := int64(duration / 86400)
	hours := int64((duration - days*86400) / 3600)
	minutes := int64((duration - days*86400 - hours*3600) / 60)
	if days > 0 {
		age += fmt.Sprintf("%dd ", days)
	}
	if hours > 0 {
		age += fmt.Sprintf("%dh ", hours)
	}
	if minutes > 0 {
		age += fmt.Sprintf("%dm ", minutes)
	}
	if age == "" {
		age = "just now"
	} else {
		age = strings.TrimRight(age, " ")
	}
	return age
}
