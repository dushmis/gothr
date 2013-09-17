package main

import (
	"call"
	"fmt"
	//"time"
	"io/ioutil"
	"net/http"
)

var (
	count int
)

func main() {
	c := &call.Gothr{}
	c.Init(20)
	c.Execute(Scan)

}

func Scan() int {
	count += 1
	fmt.Printf("%d\n", count)

	resp, _ := http.Get("https://twitter.com/users/username_available?scribeContext%5Bcomponent%5D=form&scribeContext%5Belement%5D=screen_name&username=abcd")
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", body)

	//time.Sleep(1*time.Millisecond)
	return 9
}
