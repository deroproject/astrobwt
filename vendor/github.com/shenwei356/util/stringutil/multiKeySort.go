package stringutil

import (
	"strconv"
	"strings"

	"github.com/shenwei356/natsort"
)

// SortType defines the sort type
type SortType struct {
	Index       int
	IgnoreCase  bool
	Natural     bool // natural order
	Number      bool
	UserDefined bool
	Reverse     bool
	Levels      map[string]int
}

// MultiKeyStringSlice sort [][]string by multiple keys
type MultiKeyStringSlice struct {
	SortTypes *[]SortType
	Value     []string
}

// MultiKeyStringSliceList is slice of MultiKeyStringSlice
type MultiKeyStringSliceList []MultiKeyStringSlice

func (list MultiKeyStringSliceList) Len() int { return len(list) }
func (list MultiKeyStringSliceList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list MultiKeyStringSliceList) Less(i, j int) bool {
	var err error
	var v int
	for _, t := range *list[i].SortTypes {
		if t.Natural {
			if t.IgnoreCase {
				v = strings.Compare(strings.ToLower(list[i].Value[t.Index]),
					strings.ToLower(list[j].Value[t.Index]))
			} else {
				v = strings.Compare(list[i].Value[t.Index], list[j].Value[t.Index])
			}
			if v == 0 {
				continue
			}

			if natsort.Compare(list[i].Value[t.Index], list[j].Value[t.Index], t.IgnoreCase) {
				v = -1
			} else {
				v = 1
			}
		} else if t.Number {
			var a, b float64
			a, err = strconv.ParseFloat(removeComma(list[i].Value[t.Index]), 64)
			if err != nil {
				a = 0
			}
			b, err = strconv.ParseFloat(removeComma(list[j].Value[t.Index]), 64)
			if err != nil {
				b = 0
			}
			if a < b {
				v = -1
			} else if a == b {
				v = 0
			} else {
				v = 1
			}
		} else if t.UserDefined {
			var a, b int
			var okA, okB bool
			if t.IgnoreCase {
				a, okA = t.Levels[strings.ToLower(list[i].Value[t.Index])]
				b, okB = t.Levels[strings.ToLower(list[j].Value[t.Index])]
			} else {
				a, okA = t.Levels[list[i].Value[t.Index]]
				b, okB = t.Levels[list[j].Value[t.Index]]
			}
			if okA {
				if okB {
					if a < b {
						v = -1
					} else if a == b {
						v = 0
					} else {
						v = 1
					}
				} else {
					v = -1
				}
			} else if okB {
				v = 1
			} else {
				v = strings.Compare(list[i].Value[t.Index], list[j].Value[t.Index])
			}
		} else {
			if t.IgnoreCase {
				v = strings.Compare(strings.ToLower(list[i].Value[t.Index]),
					strings.ToLower(list[j].Value[t.Index]))
			} else {
				v = strings.Compare(list[i].Value[t.Index], list[j].Value[t.Index])
			}
		}

		if v == 0 {
		} else if v < 0 {
			if t.Reverse {
				return false
			}
			return true
		} else {
			if t.Reverse {
				return true
			}
			return false
		}
	}
	return true
}

func removeComma(s string) string {
	newSlice := []byte{}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ',':
		default:
			newSlice = append(newSlice, s[i])
		}
	}
	return string(newSlice)
}
