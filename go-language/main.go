package main

import (
	"log"
	"time"
)

func printDenganKey(key string, value string) {
	log.Printf("[%s]: %s\n", key, value)
}

func nantiJalanDiThreadLain(nama string, milisecond time.Duration) {
	printDenganKey(nama, "thread mulai berjalan")
	printDenganKey(nama, "thread mulai menunggu")
	time.Sleep(milisecond)
	printDenganKey(nama, "thread selesai")
}

func main() {
	printDenganKey("main", "program mulai dijalankan")
	go nantiJalanDiThreadLain("thread 1", 500*time.Millisecond)
	go nantiJalanDiThreadLain("thread 2", 400*time.Millisecond)
	time.Sleep(450 * time.Millisecond)
	printDenganKey("main", "lanjut sebentar")
	time.Sleep(100 * time.Millisecond)
	printDenganKey("main", "program selesai")
}
