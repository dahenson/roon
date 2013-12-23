package main

import (
	"fmt"
	"github.com/dahenson/roon"
	"log"
)

func main() {
	b, err := roon.GetBlog("dane")
	if err != nil {
		log.Fatal(err)
	}
	posts, err := b.GetPosts()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Posts from %s %s\n", b.User.GivenName, b.User.FamilyName)
	for _, post := range *posts {
		fmt.Println(post.Title)
	}
}
