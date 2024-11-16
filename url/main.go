package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Dwaipayan URL")
	myUrl := "https://www.youtube.com/watch?v=vu6ZQ-t1sUk&list=PLzjZaW71kMwSEVpdbHPr0nPo5zdzbDulm&index=29"
	fmt.Printf("%T", myUrl)

	parsedURL, err := url.Parse(myUrl)

	if err != nil {
		fmt.Println("Can't parse URL", err)
		return
	}

	fmt.Printf("%T", parsedURL)

	fmt.Println("Scheme: ", parsedURL.Scheme)
	fmt.Println("Host: ", parsedURL.Host)
	fmt.Println("Path: ", parsedURL.Path)
	fmt.Println("RawQuery: ", parsedURL.RawQuery)

	//modifying
	parsedURL.Path = "/newPath"
	parsedURL.RawQuery = "username = dwaipayan"

	newUrl := parsedURL.String()

	fmt.Printf("newUrl: ", newUrl)

}
