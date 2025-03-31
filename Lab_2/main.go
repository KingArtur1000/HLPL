package main

import (
	"fmt"
	"sort"
)

// IGift интерфейс определяет базовый функционал для подарков
type IGift interface {
	GetInformation()
}

// GiftItem базовая структура для подарков
type GiftItem struct {
	name     string
	quantity uint
	price    float64
}

var totalItems int // статическая переменная для отслеживания общего количества

// Конструктор для GiftItem
func NewGiftItem(name string, quantity uint, price float64) *GiftItem {
	totalItems += int(quantity)
	return &GiftItem{
		name:     name,
		quantity: quantity,
		price:    price,
	}
}

func (g *GiftItem) GetInformation() {
	fmt.Printf("Название: %s\n", g.name)
	fmt.Printf("Количество: %d\n", g.quantity)
	fmt.Printf("Цена: %.2f\n", g.price)
}

func (g *GiftItem) UpdateInfo(name string, quantity uint, price float64, _ string) {
	g.name = name
	g.quantity = quantity
	g.price = price
}

func (g *GiftItem) UpdateQuantity(amount int) {
	g.quantity = uint(int(g.quantity) + amount)
	totalItems += amount
}

func GetTotalItems() int {
	return totalItems
}

// Toy структура наследуется от GiftItem
type Toy struct {
	GiftItem
	typeToy string
}

func NewToy(name string, quantity uint, price float64, typeToy string) *Toy {
	return &Toy{
		GiftItem: *NewGiftItem(name, quantity, price),
		typeToy:  typeToy,
	}
}

func (t *Toy) GetInformation() {
	t.GiftItem.GetInformation()
	fmt.Printf("Тип игрушки: %s\n", t.typeToy)
}

func (t *Toy) UpdateInfo(name string, quantity uint, price float64, typeToy string) {
	t.GiftItem.UpdateInfo(name, quantity, price, "")
	t.typeToy = typeToy
}

// GiftSet структура наследуется от GiftItem
type GiftSet struct {
	GiftItem
	contents string
}

func NewGiftSet(name string, quantity uint, price float64, contents string) *GiftSet {
	return &GiftSet{
		GiftItem: *NewGiftItem(name, quantity, price),
		contents: contents,
	}
}

func (gs *GiftSet) GetInformation() {
	gs.GiftItem.GetInformation()
	fmt.Printf("Содержимое набора: %s\n", gs.contents)
}

func (gs *GiftSet) UpdateInfo(name string, quantity uint, price float64, contents string) {
	gs.GiftItem.UpdateInfo(name, quantity, price, "")
	gs.contents = contents
}

// GiftStorage структура для хранения подарков
type GiftStorage struct {
	gifts []IGift
}

func (gs *GiftStorage) AddGift(gift IGift) {
	gs.gifts = append(gs.gifts, gift)
}

func (gs *GiftStorage) ShowAllGiftItems() {
	for _, gift := range gs.gifts {
		gift.GetInformation()
		fmt.Println("-----------------")
	}
}

func (gs *GiftStorage) UpdateGiftInfo(pos uint, name string, quantity uint, price float64, extraContent string) {
	if int(pos) >= len(gs.gifts) {
		fmt.Println("Ошибка: некорректный индекс!")
		return
	}

	switch gift := gs.gifts[pos].(type) {
	case *Toy:
		gift.UpdateInfo(name, quantity, price, extraContent)
	case *GiftSet:
		gift.UpdateInfo(name, quantity, price, extraContent)
	}
}

func (gs *GiftStorage) UpdateGiftQuantityByIndex(pos uint, quantity int) {
	if int(pos) >= len(gs.gifts) {
		fmt.Println("Ошибка: некорректный индекс!")
		return
	}

	switch gift := gs.gifts[pos].(type) {
	case *Toy:
		gift.UpdateQuantity(quantity)
	case *GiftSet:
		gift.UpdateQuantity(quantity)
	}
}

func (gs *GiftStorage) FindGiftByName(name string) {
	found := false
	for _, gift := range gs.gifts {
		var giftName string
		switch g := gift.(type) {
		case *Toy:
			giftName = g.name
		case *GiftSet:
			giftName = g.name
		}

		if giftName == name {
			fmt.Println("\tПодарок найден:")
			gift.GetInformation()
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Подарок по названию: %s не найден.\n\n", name)
	}
}

func (gs *GiftStorage) SortGiftsByName() {
	sort.Slice(gs.gifts, func(i, j int) bool {
		var name1, name2 string

		switch g1 := gs.gifts[i].(type) {
		case *Toy:
			name1 = g1.name
		case *GiftSet:
			name1 = g1.name
		}

		switch g2 := gs.gifts[j].(type) {
		case *Toy:
			name2 = g2.name
		case *GiftSet:
			name2 = g2.name
		}

		return name1 < name2
	})
}

func main() {
	giftStorage := &GiftStorage{}

	giftStorage.AddGift(NewToy("Мишка", 10, 500.0, "Мягкая игрушка"))
	giftStorage.AddGift(NewToy("Конструктор", 15, 1200.0, "Развивающая игрушка"))
	giftStorage.AddGift(NewToy("Кукла", 8, 800.0, "Кукла"))
	giftStorage.AddGift(NewGiftSet("Детский набор", 5, 1500.0, "Шоколад, игрушка, книга"))
	giftStorage.AddGift(NewGiftSet("Сладкий набор", 7, 1000.0, "Конфеты, мягкая игрушка"))

	fmt.Println("\tПервоначальный список подарков:")
	giftStorage.ShowAllGiftItems()
	fmt.Printf("\tИТОГО: Общее кол-во подарков: %d\n", GetTotalItems())

	giftStorage.UpdateGiftInfo(2, "Робот", 12, 1500.0, "Электронная игрушка")
	giftStorage.UpdateGiftQuantityByIndex(0, 5)
	giftStorage.UpdateGiftQuantityByIndex(2, 3)

	fmt.Println("\n\tОбновленный список подарков:")
	giftStorage.ShowAllGiftItems()
	fmt.Printf("\tИТОГО: Общее кол-во подарков: %d\n", GetTotalItems())

	giftStorage.UpdateGiftQuantityByIndex(4, 2)

	fmt.Println("\n\tФинальный список подарков:")
	giftStorage.ShowAllGiftItems()
	fmt.Printf("\tИТОГО: Общее кол-во подарков: %d\n", GetTotalItems())

	fmt.Println("\n\tПоиск подарка по имени:")
	giftStorage.FindGiftByName("Чика")
	giftStorage.FindGiftByName("Робот")

	fmt.Println("\n\tСортировка подарков по имени:")
	giftStorage.SortGiftsByName()
	giftStorage.ShowAllGiftItems()
}
