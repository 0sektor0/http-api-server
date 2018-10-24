package tests

import (
	"reflect"
	"testing"

	m "http-api-server/models"
)

func ConcludeTestResult(expected interface{}, result interface{}, unmarshalError error, t *testing.T) {
	if unmarshalError != nil {
		t.Error(unmarshalError)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: %v\n got: %v\n", result, expected)
	}
}

func TestErrorUnmarshal(t *testing.T) {
	json := `{
		"message": "Can't find user with id #42\n"
	  }`

	expected := &m.Error{
		Message: "Can't find user with id #42\n",
	}

	result, err := m.UnmarshalError([]byte(json))
	ConcludeTestResult(expected, result, err, t)
}

func TestForumUnmarshal(t *testing.T) {
	json := `{
		"posts": 200000,
		"slug": "pirate-stories",
		"threads": 200,
		"title": "Pirate stories",
		"user": "j.sparrow"
	  }`

	expected := &m.Forum{
		Posts:   200000,
		Slug:    "pirate-stories",
		Threads: 200,
		Title:   "Pirate stories",
		User:    "j.sparrow",
	}

	result, err := m.UnmarshalForum([]byte(json))
	ConcludeTestResult(expected, result, err, t)
}

func TestPostFullUnmarshal(t *testing.T) {
	json := `{
		"author": {
		  "about": "This is the day you will always remember as the day that you almost caught Captain Jack Sparrow!",
		  "email": "captaina@blackpearl.sea",
		  "fullname": "Captain Jack Sparrow",
		  "nickname": "j.sparrow"
		},
		"forum": {
		  "posts": 200000,
		  "slug": "pirate-stories",
		  "threads": 200,
		  "title": "Pirate stories",
		  "user": "j.sparrow"
		},
		"post": {
		  "author": "j.sparrow",
		  "created": "2018-09-10T10:26:41.188Z",
		  "forum": "string",
		  "id": 0,
		  "isEdited": true,
		  "message": "We should be afraid of the Kraken.",
		  "parent": 0,
		  "thread": 0
		},
		"thread": {
		  "author": "j.sparrow",
		  "created": "2017-01-01T00:00:00.000Z",
		  "forum": "pirate-stories",
		  "id": 42,
		  "message": "An urgent need to reveal the hiding place of Davy Jones. Who is willing to help in this matter?",
		  "slug": "jones-cache",
		  "title": "Davy Jones cache",
		  "votes": 0
		}
	  }`

	_, err := m.UnmarshalPostFull([]byte(json))

	if err != nil {
		t.Error(err)
	}
}

func TestPostUnmarshal(t *testing.T) {
	json := `{
		"author": "j.sparrow",
		"created": "2018-09-10T10:31:44.251Z",
		"forum": "string",
		"id": 0,
		"isEdited": true,
		"message": "We should be afraid of the Kraken.",
		"parent": 0,
		"thread": 0
	  }`

	expected := &m.Post{
		Author:   "j.sparrow",
		Created:  "2018-09-10T10:31:44.251Z",
		Forum:    "string",
		Id:       0,
		Isedited: true,
		Message:  "We should be afraid of the Kraken.",
		Parent:   0,
		Thread:   0,
	}

	result, err := m.UnmarshalPost([]byte(json))
	ConcludeTestResult(expected, result, err, t)
}

func TestPostUpdateUnmarshal(t *testing.T) {
	json := `{
		"message": "We should be afraid of the Kraken."
	  }`

	expected := &m.PostUpdate{
		Message: "We should be afraid of the Kraken.",
	}

	result, err := m.UnmarshalPostUpdate([]byte(json))
	ConcludeTestResult(expected, result, err, t)
}

func TestStatusUnmarshal(t *testing.T) {
	json := `{
		"forum": 100,
		"post": 1000000,
		"thread": 1000,
		"user": 1000
	}`

	expected := &m.Status{
		Forum:  100,
		Post:   1000000,
		Thread: 1000,
		User:   1000,
	}

	result, err := m.UnmarshalStatus([]byte(json))
	ConcludeTestResult(expected, result, err, t)
}

func TestThreadUnmarshal(t *testing.T) {
	json := `{
		"author": "j.sparrow",
		"created": "2017-01-01T00:00:00.000Z",
		"forum": "pirate-stories",
		"id": 42,
		"message": "An urgent need to reveal the hiding place of Davy Jones. Who is willing to help in this matter?",
		"slug": "jones-cache",
		"title": "Davy Jones cache",
		"votes": 0
	}`

	expected := &m.Thread{
		Author:  "j.sparrow",
		Created: "2017-01-01T00:00:00.000Z",
		Forum:   "pirate-stories",
		Id:      42,
		Message: "An urgent need to reveal the hiding place of Davy Jones. Who is willing to help in this matter?",
		Slug:    "jones-cache",
		Title:   "Davy Jones cache",
		Votes:   0,
	}

	result, err := m.UnmarshalThread([]byte(json))
	ConcludeTestResult(expected, result, err, t)
}

func TestThreadUpdateUnmarshal(t *testing.T) {
	json := `{
		"message": "An urgent need to reveal the hiding place of Davy Jones. Who is willing to help in this matter?",
		"title": "Davy Jones cache"
	  }`

	expected := &m.ThreadUpdate{
		Message: "An urgent need to reveal the hiding place of Davy Jones. Who is willing to help in this matter?",
		Title:   "Davy Jones cache",
	}

	result, err := m.UnmarshalThreadUpdate([]byte(json))
	ConcludeTestResult(expected, result, err, t)
}

func TestUserUnmarshal(t *testing.T) {
	json := `{
		"about": "This is the day you will always remember as the day that you almost caught Captain Jack Sparrow!",
		"email": "captaina@blackpearl.sea",
		"fullname": "Captain Jack Sparrow",
		"nickname": "j.sparrow"
	  }`

	expected := &m.User{
		About:    "This is the day you will always remember as the day that you almost caught Captain Jack Sparrow!",
		Email:    "captaina@blackpearl.sea",
		FullName: "Captain Jack Sparrow",
		NickName: "j.sparrow",
	}

	result, err := m.UnmarshalUser([]byte(json))
	ConcludeTestResult(expected, result, err, t)
}

func TestUserUpdateUnmarshal(t *testing.T) {
	json := `{
		"about": "This is the day you will always remember as the day that you almost caught Captain Jack Sparrow!",
		"email": "captaina@blackpearl.sea",
		"fullname": "Captain Jack Sparrow"
	  }`

	expected := &m.UserUpdate{
		About:    "This is the day you will always remember as the day that you almost caught Captain Jack Sparrow!",
		Email:    "captaina@blackpearl.sea",
		FullName: "Captain Jack Sparrow",
	}

	result, err := m.UnmarshalUserUpdate([]byte(json))
	ConcludeTestResult(expected, result, err, t)
}

func TestVoteUnmarshal(t *testing.T) {
	json := `{
		"nickname": "Sekougi",
		"voice": -1
	}`

	expected := &m.Vote{
		NickName: "Sekougi",
		Voice:    -1,
	}

	result, err := m.UnmarshalVote([]byte(json))
	ConcludeTestResult(expected, result, err, t)
}
