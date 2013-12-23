# roon

Just another [Roon API](https://roon.io/developer) wrapper. [Roon](https://roon.io) is one of the slickest blog farms on the internet. Go get a blog, then `go get` this api wrapper and start hacking!

## Install
Do the go get dance:

`go get github.com/dahenson/roon`

## Example
It's fairly simple. You can get my blog, then get a post from my blog like this:


```
package main

import (
	"fmt"
	"github.com/dahenson/roon"
	"log"
)

func main() {
	b, err := roon.GetBlog("dane") // This is my blog!
	if err != nil {
		log.Fatal(err)
	}

	posts, err := b.GetPosts() // Get all of my posts
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Posts from %s\n", b.User.Username)
	for _, post := range *posts {
		fmt.Println(post.Title)
	}
}
```

I might make it more sophisticated in the future, but this works really well for now. Don't be afraid to submit a pull request. I don't bite.

## Documentation

Simply run `godoc -http=":6060"` and navigate to `localhost:6060/pkg/github.com/dahenson/roon`

Or

Go to [godoc.org](http://godoc.org/github.com/dahenson/roon).

A special thanks to Go for writing my documentation for me.