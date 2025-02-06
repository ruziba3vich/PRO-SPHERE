/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-09-29 08:50:11
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-11-20 22:47:22
 * @FilePath: /sfere_backend/internal/items/models/search/images.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package models

type ImageSearchAPIResponse struct {
	Kind       string          `json:"kind"`
	Url        URL             `json:"url"`
	Queries    Queries         `json:"queries"`
	Context    Context         `json:"context"`
	SearchInfo ImageSearchInfo `json:"searchInformation"`
	Items      []Item          `json:"items"`
}

// URL Struct
type URL struct {
	Type     string `json:"type"`
	Template string `json:"template"`
}

// Queries Struct for requests and nextPage
type Queries struct {
	Request  []Request `json:"request"`
	NextPage []Request `json:"nextPage"`
}

// Request Struct inside Queries
type Request struct {
	Title          string `json:"title"`
	TotalResults   string `json:"totalResults"`
	SearchTerms    string `json:"searchTerms"`
	Count          int    `json:"count"`
	StartIndex     int    `json:"startIndex"`
	InputEncoding  string `json:"inputEncoding"`
	OutputEncoding string `json:"outputEncoding"`
	Safe           string `json:"safe"`
	Cx             string `json:"cx"`
	SearchType     string `json:"searchType"`
}

// Context Struct
type Context struct {
	Title string `json:"title"`
}

// SearchInfo Struct for search information
type ImageSearchInfo struct {
	SearchTime            float64 `json:"searchTime"`
	FormattedSearchTime   string  `json:"formattedSearchTime"`
	TotalResults          string  `json:"totalResults"`
	FormattedTotalResults string  `json:"formattedTotalResults"`
}

// Item Struct for each search result item
type Item struct {
	Kind        string `json:"kind"`
	Title       string `json:"title"`
	HtmlTitle   string `json:"htmlTitle"`
	Link        string `json:"link"`
	DisplayLink string `json:"displayLink"`
	Snippet     string `json:"snippet"`
	HtmlSnippet string `json:"htmlSnippet"`
	Mime        string `json:"mime"`
	FileFormat  string `json:"fileFormat"`
	Image       Image  `json:"image"`
}

// Image Struct for image-related fields
type Image struct {
	ContextLink     string `json:"contextLink"`
	Height          int    `json:"height"` // Change to int
	Width           int    `json:"width"`  // Change to int
	ByteSize        int    `json:"byteSize"`
	ThumbnailLink   string `json:"thumbnailLink"`
	ThumbnailHeight int    `json:"thumbnailHeight"`
	ThumbnailWidth  int    `json:"thumbnailWidth"`
}
