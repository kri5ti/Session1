package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

func getPicture(dest string, elem string, wg *sync.WaitGroup, ch *chan int) {
	defer func() { <-*ch }()
	defer wg.Done()
	out, _ := os.Create(dest)
	defer out.Close()
	resp, _ := http.Get(elem)
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(dest, b, 0644)

}
func main() {
	var wg sync.WaitGroup
	js := []byte(`{ "Urls": [
		"https://timedotcom.files.wordpress.com/2016/02/blue-sky-colorful-balloons.jpg?quality=85&w=1100",
        "http://www.sz0931.com/data/out/39/52585566-happiness-wallpaper.jpg",
        "https://pp.vk.me/c618425/v618425852/19565/IAett21zJsU.jpg",
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
		log.Fatal(err)
	}
	ch := make(chan int, 2)
	for i, _ := range t.Urls {
		wg.Add(1)
		ch <- 1
		dest := "D:/git/Session1/22.Semaphore/photo_" + strconv.Itoa(i) + ".jpg"
		go getPicture(dest, t.Urls[i], &wg, &ch)
	}
	wg.Wait()
}
