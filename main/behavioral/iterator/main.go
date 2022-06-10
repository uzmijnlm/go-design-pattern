package main

import "fmt"

func main() {
	user1 := &user{
		name: "a",
	}
	user2 := &user{
		name: "b",
	}

	users := []*user{user1, user2}

	userCollection := &userCollection{
		users: users,
	}

	iterator := userCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}
