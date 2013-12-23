package roon

import (
	"encoding/json"
	"net/http"
	"strconv"
)

const (
	RoonUrl    = "https://roon.io"
	RoonApiUrl = RoonUrl + "/api/v1"
)

type Blog struct {
	AccentColor      string   `json:"accent_color"`
	Addons           []string `json:"addons"`
	AnalyticsId      string   `json:"analytics_id"`
	AnalyticsService string   `json:"analytics_service"`
	CreatedAt        int      `json:"created_at"`
	CustomDomain     string   `json:"custom_domain"`
	FontStyle        string   `json:"font_style"`
	Id               int      `json:"id"`
	Language         string   `json:"language"`
	Subdomain        string   `json:"subdomain"`
	Title            string   `json:"title"`
	UpdatedAt        int      `json:"updated_at"`
	Url              string   `json:"url"`
	User             User     `json:"user"`
}

type User struct {
	AccentColor   string `json:"accent_color"`
	AvatarUrl     string `json:"avatar_url"`
	Bio           string `json:"bio"`
	BioHtml       string `json:"bio_html"`
	CreatedAt     int    `json:"created_at"`
	FamilyName    string `json:"family_name"`
	GivenName     string `json:"given_name"`
	Id            int    `json:"id"`
	UpdatedAt     int    `json:"updated_at"`
	Username      string `json:"username"`
	Website       string `json:"website"`
	WritingFormat string `json:"writing_format"`
}

// GetBlog() retrieves the blog specified by the blog id
func GetBlog(id string) (*Blog, error) {
	var url = RoonApiUrl + "/blogs/" + id
	var blog = new(Blog)

	if err := roonRequest(url, blog); err != nil {
		return nil, err
	}

	return blog, nil
}

// GetPost() retrieves the post specified by the post id
func (b *Blog) GetPost(id string) (*Post, error) {
	return GetPost(strconv.Itoa(b.Id), id)
}

// GetPosts() retrieves all posts in the blog
func (b *Blog) GetPosts() (*[]Post, error) {
	return GetPosts(strconv.Itoa(b.Id))
}

type Post struct {
	BlogId         int            `json:"blog_id"`
	CharacterCount int            `json:"character_count"`
	Content        string         `json:"content"`
	ContentHtml    string         `json:"content_html"`
	CreatedAt      int            `json:"created_at"`
	DeletedAt      int            `json:"deleted_at"`
	Excerpt        string         `json:"excerpt"`
	ExcerptHtml    string         `json:"excerpt_html"`
	FeaturedMedium FeaturedMedium `json:"featured_medium"`
	Id             int            `json:"id"`
	LikesCount     int            `json:"likes_count"`
	PublishedAt    int            `json:"published_at"`
	ReadTime       int            `json:"read_time"`
	Slug           string         `json:"slug"`
	Title          string         `json:"title"`
	UpdatedAt      int            `json:"updated_at"`
	Url            string         `json:"url"`
	User           User           `json:"user"`
	WordCount      int            `json:"word_count"`
}

type FeaturedMedium struct {
	Id             int    `json:"id"`
	Token          string `json:"token"`
	UserId         int    `json:"user_id"`
	BlogId         int    `json:"blog_id"`
	PostId         int    `json:"post_id"`
	TotalBytes     int    `json:"total_bytes"`
	MimeType       string `json:"mime_type"`
	UploadedAt     int    `json:"uploaded_at"`
	CreatedAt      int    `json:"created_at"`
	UpdatedAt      int    `json:"updated_at"`
	ProcessedAt    int    `json:"processed_at"`
	OriginalImage  Image  `json:"original"`
	GiantImage     Image  `json:"giant"`
	LargeImage     Image  `json:"large"`
	SmallImage     Image  `json:"small"`
	ThumbnailImage Image  `json:"Thumbnail"`
}

type Image struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Bytes  int    `json:"bytes"`
}

// GetPost() gets a single post from a blog, given the blog id and post id
func GetPost(blogId, postId string) (*Post, error) {
	var url = RoonApiUrl + "/blogs/" + blogId + "/posts/" + postId
	var post = new(Post)

	if err := roonRequest(url, post); err != nil {
		return nil, err
	}

	return post, nil
}

// GetPosts() gets all posts from a specified blog
func GetPosts(id string) (*[]Post, error) {
	var url = RoonApiUrl + "/blogs/" + id + "/posts/"
	var posts = new([]Post)

	if err := roonRequest(url, posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func roonRequest(url string, v interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return nil // TODO: return an error of some kind
	}

	dec := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	if err := dec.Decode(v); err != nil {
		return err
	}

	return nil
}
