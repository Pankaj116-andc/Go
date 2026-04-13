package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning URL")
	myURL := "https://www.youtube.com/watch?v=vu6ZQ-t1sUk&list=PLzjZaW71kMwSEVpdbHPr0nPo5zdzbDulm&index=26"
	fmt.Printf("Type of myURL is %T\n", myURL)
	parseurl, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("Error while parsing the url ", err)
		return
	}
	fmt.Printf("Type of parseurl is %T\n", parseurl)

	fmt.Println("Scheme: ", parseurl.Scheme)
	fmt.Println("Host: ", parseurl.Host)
	fmt.Println("Path: ", parseurl.Path)
	fmt.Println("RawQuery: ", parseurl.RawQuery)

	//modifying url
	parseurl.Path = "/newPath"
	parseurl.RawQuery = "username=iampankaj"

	newUrl := parseurl.String()
	fmt.Println("new URL: ", newUrl)
}
