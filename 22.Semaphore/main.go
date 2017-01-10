package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type locker struct {
	ch     chan int
	status bool
}

func (l *locker) lock() {
	l.status = true
	l.ch <- 1
}

func (l *locker) unlock() {
	if l.status {
		l.status = false
		<-l.ch
	}
}
func getPicture(dest string, elem string, lock locker) {
	lock.lock()
	defer lock.unlock()
	out, _ := os.Create(dest)
	defer out.Close()
	resp, _ := http.Get(elem)
	defer resp.Body.Close()
	_, err := io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(1 * time.Second)

}
func main() {
	ch := make(chan int, 2)
	js := []byte(`{ "Urls": [
		"https://timedotcom.files.wordpress.com/2016/02/blue-sky-colorful-balloons.jpg?quality=85&w=1100",
        "http://www.sz0931.com/data/out/39/52585566-happiness-wallpaper.jpg",
        "https://timedotcom.files.wordpress.com/2016/02/blue-sky-colorful-balloons.jpg?quality=85&w=1100",
        "http://www.sz0931.com/data/out/39/52585566-happiness-wallpaper.jpg",
        "https://timedotcom.files.wordpress.com/2016/02/blue-sky-colorful-balloons.jpg?quality=85&w=1100",
        "http://www.sz0931.com/data/out/39/52585566-happiness-wallpaper.jpg",
        "https://timedotcom.files.wordpress.com/2016/02/blue-sky-colorful-balloons.jpg?quality=85&w=1100",
        "http://www.sz0931.com/data/out/39/52585566-happiness-wallpaper.jpg",
        "https://timedotcom.files.wordpress.com/2016/02/blue-sky-colorful-balloons.jpg?quality=85&w=1100",
        "http://www.sz0931.com/data/out/39/52585566-happiness-wallpaper.jpg",
        "https://timedotcom.files.wordpress.com/2016/02/blue-sky-colorful-balloons.jpg?quality=85&w=1100",
        "http://www.sz0931.com/data/out/39/52585566-happiness-wallpaper.jpg"
        ] } 
        `)
	type st struct {
		Urls []string
	}
	var t st
	err := json.Unmarshal(js, &t)
	if err != nil {
		fmt.Println("error:", err)
	}
	for i := 0; i < len(t.Urls)-1; i++ {

		dest := "D:/git/Session1/22.Semaphore/photo_" + strconv.Itoa(i) + ".jpg"
		go getPicture(dest, t.Urls[i], locker{ch, false})

	}
	time.Sleep(15 * time.Second)
}
