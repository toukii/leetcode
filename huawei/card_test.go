package huawei

import (
	"fmt"
	"sync"
	"testing"
)

var (
	cardMap    map[uint8]*Card
	input      string
	check3Once *sync.Once
	check4Once *sync.Once
)

func init() {
	input = "T1T2T3T4T5T6D2D2D2D2D1D2D3"
	// input = "T1T1T2T2T1T1T2T2T1T1T2T2T1T1T2T2T1T1T2T2T1T1T2T2T1T1T2T2"
	cardMap = make(map[uint8]*Card)
	check3Once = &sync.Once{}
	check4Once = &sync.Once{}
}

type Card struct {
	C     uint8
	P     uint8
	count int
	next  *Card
}

func (self *Card) equals(c *Card) bool {
	if self.C == c.C && self.P == c.P {
		return true
	}
	return false
}

func (c *Card) String() string {
	return fmt.Sprintf("%v%v %d", string(c.C), string(c.P), c.count)
}

func (c *Card) Travel() {
	ca := c
	for {
		if ca == nil {
			return
		}
		fmt.Println(ca)
		ca = ca.next
	}
}

func (c *Card) Check() int {
	c3 := c.Check3Func()
	if c3 >= 1 {
		fmt.Println("卡2条，一番；")
	}
	c4 := c.Check4Func()
	if c4 >= 1 {
		fmt.Println("四归一，两番；")
	}
	c7 := c.Check7Func()
	if c7 >= 1 {
		fmt.Println("巧七，两番。")
	}
	ret := c3 + c4*2 + c7*2
	return ret
}

func (c *Card) Check3Func() int {
	if c.C != 'T' {
		return 0
	}
	result := c.check(1, 3)
	ret := 0
	if result {
		check3Once.Do(func() {
			ret = 1
		})
	}
	return ret
}
func (c *Card) Check4Func() int {
	result := c.check(4, 1)
	ret := 0
	if result {
		check4Once.Do(func() {
			ret = 1
		})
	}
	return ret
}

func (c *Card) Check7Func() int {
	result := c.check(2, 7)
	if result {
		return 1
	}
	return 0
}

func (c *Card) check(count, sum int) bool {
	i := 0
	ca := c
	for {
		if i >= sum {
			return true
		}
		if ca == nil {
			break
		}
		if ca.count == count {
			i++
		}
		ca = ca.next
	}
	return false
}

func newCard(c, p uint8) *Card {
	return &Card{
		C:     c,
		P:     p,
		count: 1,
	}
}

func TestTwo(t *testing.T) {
	length := len(input)
	for i := length - 1; i >= 1; i -= 2 {
		card := newCard(input[i-1], input[i])
		if headcard, ok := cardMap[card.C]; !ok {
			cardMap[card.C] = card
		} else {
			if headcard.equals(card) {
				headcard.count++
			} else {
				card.next = headcard
				cardMap[card.C] = card
			}
		}
	}
	fmt.Println(input)
	if len(input) > 36 {
		return
	}
	var result int
	for k, v := range cardMap {
		fmt.Println(k)
		v.Travel()
		result += v.Check()
	}

	fmt.Printf("共翻%d番\n", result)
}
