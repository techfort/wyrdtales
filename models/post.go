package models

import (
	"time"
)

// Post models
type Post struct {
	PostID   string    `json:"postid"`
	AuthorID string    `json:"authorid"`
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	Tags     []string  `json:"tags"`
	Posted   time.Time `json:"posted"`
	Status   string    `json:"status"`
	Category string    `json:"category"`
}

// PostMapping is the mapping of elasticsearch properties
const PostMapping = `
	{
		"settings": {
			"number_of_shards": 1,
			"number_of_replicas": 0
		},
		"mappings": {
			"post": {
				"properties": {
					"title": {
						"type": "keyword"
					},
					"authorid": {
						"type": "keyword"
					},
					"body": {
						"type": "text",
						"store": true,
						"fielddata": true
					},
					"tags": {
						"type": "keyword"
					},
					"posted": {
						"type": "date"
					},
					"status": {
						"type": "keyword"
					},
					"category": {
						"type": "keyword"
					}
				}
			}
		}
	}
`
