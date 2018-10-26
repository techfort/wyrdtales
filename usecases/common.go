package usecases

import (
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic"
	"github.com/techfort/wyrdtales/models"
)

// UnmarshalPost unmarshales a post
func UnmarshalPost(source []byte) (models.Post, error) {
	var post models.Post
	err := json.Unmarshal(source, &post)
	return post, err
}

// UnmarshalPostSlice unmrashals an array of posts
func UnmarshalPostSlice(result *elastic.SearchResult) ([]models.Post, error) {
	var (
		posts = make([]models.Post, 0)
		err   error
	)
	for _, r := range result.Hits.Hits {
		if p, err := UnmarshalPost(*r.Source); err != nil {
			fmt.Println(fmt.Sprintf("Error unmarshaling post: %v", err.Error()))
		} else {
			posts = append(posts, p)
		}
	}
	return posts, err
}
