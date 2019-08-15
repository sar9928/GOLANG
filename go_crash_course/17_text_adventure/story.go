package main

import (
	"fmt"
)

// Recursion creating linked lists
type storyPage struct {
	text     string
	nextPage *storyPage
}

func playStory(page *storyPage) {
	if page == nil {
		return
	}
	fmt.Println(page.text)
	playStory(page.nextPage)
}

func deletePage(page *storyPage) {
	page.text = ""
}

func main() {
	//scanner := bufio.NewScanner(os.Stdin)

	// Linked list of pages
	page1 := storyPage{"You walk into a cavern", nil}
	page2 := storyPage{"You are alone, and you need to find a chair", nil}
	page3 := storyPage{"Your legs hurt so much from standing", nil}
	page4 := storyPage{"Do you want to continue", nil}
	page5 := storyPage{"That's weird.", nil}

	page1.nextPage = &page2
	page2.nextPage = &page3
	page3.nextPage = &page4
	page4.nextPage = &page5

	deletePage(&page4)
	playStory(&page1)
}
