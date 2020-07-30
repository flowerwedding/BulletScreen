package service

import (
	"BulletScreen/model"
	"log"
	"sync"
	"time"
)

type User struct {
	UserId string
	GoodsId  uint
}

var OrderChan = make(chan User, 1024)

var ItemMap = make(map[uint]*Item)

type Item struct {
	ID        uint   // 商品id
	Name      string // 名字
	Total     int    // 商品总量
	Left      int    // 商品剩余数量
	IsSoldOut bool   // 是否售罄
	leftCh    chan int
	sellCh    chan int
	done      chan struct{}
	Lock      sync.Mutex
}

func initMap() {
	item := &Item{
		ID:        1,
		Name:      "测试",
		Total:     100,
		Left:      100,
		IsSoldOut: false,
		leftCh:    make(chan int),
		sellCh:    make(chan int),
	}
	ItemMap[item.ID] = item
}

func getItem(itemId uint) *Item{
	return ItemMap[itemId]
}

func order() {
	for {
		user := <- OrderChan
		item := getItem(user.GoodsId)
		item.SecKilling(user.UserId)
	}
}

func (item *Item) SecKilling(userId string) {

	item.Lock.Lock()
	defer item.Lock.Unlock()

	if item.IsSoldOut {
		return
	}
	item.BuyGoods(1)

	MakeOrder(userId, item.ID,1)


}

func MakeOrder(userId string, goodsId uint, num int) {

	order := model.Order{
		UserID:  userId,
		GoodsID: goodsId,
		Num:     num,
	}
	err := order.MakeOrder()
	if err != nil {
		log.Printf("Error make an order. Error: %s",err)
	}
	log.Println("success")
}


func (item *Item) OffShelve() {
	beginTime := time.Now()
	offShelveTime := beginTime.Add(time.Second*5)
	timer := time.NewTimer(offShelveTime.Sub(beginTime))

	<-timer.C
	delete(ItemMap, item.ID)
	close(item.done)

}

func (item *Item) SalesGoods() {
	for {
		select {
		case num := <-item.sellCh:
			if item.Left -= num; item.Left <= 0 {
				item.IsSoldOut = true
			}

		case item.leftCh <- item.Left:
		case <-item.Done():
			log.Println("我自闭了")
			return
		}
	}
}

func (item *Item) Done() <-chan struct{} {
	if item.done == nil {
		item.done = make(chan struct{})
	}
	d := item.done
	return d
}

func (item *Item) Monitor() {
	go item.SalesGoods()
}

// 获取剩余库存
func (item *Item) GetLeft() int {
	var left int
	left = <-item.leftCh
	return left
}

// 购买商品
func (item *Item) BuyGoods(num int) {
	item.sellCh <- num
}

func InitService() {
	initMap()
	for _,item := range ItemMap{
		item.Monitor()
		go item.OffShelve()
	}
	for i := 0; i < 10; i++ {
		go order()
	}
}