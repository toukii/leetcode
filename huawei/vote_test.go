package huawei

import (
	"fmt"
	"testing"
)

type Candi struct {
	Name  string
	Count int
	next  *Candi
}

func NewCandi(name string) *Candi {
	return &Candi{
		Name: name,
	}
}

var (
	CandiMap              map[string]*Candi
	FirstCandi, LastCandi *Candi
)

func init() {
	CandiMap = make(map[string]*Candi)
}

func add(name string) {
	candi := NewCandi(name)
	CandiMap[name] = candi
	if FirstCandi == nil {
		FirstCandi, LastCandi = candi, candi
	} else {
		if FirstCandi == LastCandi {
			FirstCandi.next = candi
			LastCandi = candi
		} else {
			LastCandi.next = candi
			LastCandi = candi
		}
	}
}

func vote(name string) {
	if candi, ok := CandiMap[name]; !ok {
		// fmt.Println("%s no exist.")
		return
	} else {
		candi.Count++
	}
}

func result() {
	candi := FirstCandi
	for ; candi != nil; candi = candi.next {
		fmt.Printf("%s %d\n", candi.Name, candi.Count)
	}
}

func TestVote(t *testing.T) {
	add("xx1")
	add("xx2")
	add("xx3")
	add("xx4")
	add("xx5")
	add("xx6")

	vote("xx1")
	vote("xx3")
	vote("xx4")
	vote("xx1")
	vote("xx2")
	vote("xx7")
	vote("xx4")
	vote("xx5")
	vote("xx3")
	vote("xx2")
	vote("xx1")
	vote("xx7")

	result()
}
