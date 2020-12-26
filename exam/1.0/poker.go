package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

const (
	//黑桃
	Spade = 0
	//红桃
	Hearts = 1
	//梅花
	Club = 2
	//方块
	Diamond = 3
)

type Poker struct {
	Num int
	Flower int
}
type Pokers []Poker

func (p Poker)PokerSelf()string  {
	var buffer string

	switch p.Flower {
	case Spade:
		buffer += "♤"
	case Hearts:
		buffer += "♡"
	case Club:
		buffer += "♧"
	case Diamond:
		buffer += "♢"
	}
	switch p.Num {
	case 13:
		buffer += "2"
	case 12:
		buffer += "A"
	case 11:
		buffer += "K"
	case 10:
		buffer += "Q"
	case 9:
		buffer += "J"
	default:
		buffer += strconv.Itoa(p.Num+2)
	}

	return buffer
}

func CreatePokers()(pokers Pokers)  {
	for i := 1; i < 14; i++ {
		for j := 0; j < 4; j++ {
			pokers = append(pokers,Poker{
				Num:    i,
				Flower: j,
			})
		}
	}
	return
}
func (p Pokers)Print()  {
	for _, i2 := range p {
		fmt.Print(i2.PokerSelf(),"")
	}
	fmt.Println()
}

type p interface {
	Len()
	Swap()
	Less()
}

func (p Pokers) Len() int { return len(p) }
func (p Pokers) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p Pokers) Less(i, j int) bool { return p[i].Num < p[j].Num }
//copy的代码
func (p Pokers)xipai() Pokers {
	rand.Seed(time.Now().UTC().UnixNano())//这是一个根据时间来生成的一个随机数保证每次洗牌的顺序不一样
	for i := len(p); i > 0; i-- {
		last := i - 1
		idx := rand.Intn(i)//获取一个随机数
		p[last], p[idx] = p[idx], p[last]//交换切片中两个元素的位置从而打乱顺序进行洗牌
	}
	return p
}

func main()  {
	var p Pokers
	p=CreatePokers()
	p.Print()
	p=p.xipai()//洗牌
	p.Print()
	sort.Stable(p)//排序
	p.Print()
}
//接口一类方法的集合