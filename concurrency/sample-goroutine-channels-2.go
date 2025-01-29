package concurrency

import (
	"fmt"
	"time"
)

func buyGlassesAtMart() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อแว่น : ร้านสะดวกซื้อ : สำเร็จ")
}

func buyWatchAtShop() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อนาฬิกา : ร้านขายนาฬิกา : สำเร็จ")
}

func buyFruitsAtGrocery() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อผลไม้ : ร้านขายของชำ : สำเร็จ")
}

func buyCarAtShop() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อรถ : ร้านขายรถยนต์ : สำเร็จ")
}

func sendMessageToReceiver(message chan<- string) {
	time.Sleep(1 * time.Second)
	targetMessage := "กำลังส่งของให้คุณ A"
	fmt.Println(targetMessage)
	message <- targetMessage
}

func ExecuteSampleGoroutinesAndChannels2() {
	start := time.Now()
	ch := make(chan string)

	go buyGlassesAtMart()
	go buyWatchAtShop()

	go sendMessageToReceiver(ch)

	buyFruitsAtGrocery()
	buyCarAtShop()
	messageFromSender := <-ch
	if messageFromSender == "กำลังส่งของให้คุณ A" {
		fmt.Println("คุณ A ได้รับของแล้ว")
		fmt.Println("ใช้เวลาในการทำงานทั้งหมด: ", time.Since(start))
	}
	close(ch)
}
