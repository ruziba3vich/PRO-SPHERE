/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-11-19 20:07:00
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-04 15:56:43
 * @FilePath: /searching/internal/items/models/search.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package models

type SearchElementsAPIResponse struct {
	SearchInformation SearchInfo     `json:"searchInformation"`
	Items             []SearchResult `json:"items"`
}

type SearchInfo struct {
	TotalResults string  `json:"totalResults"`
	SearchTime   float64 `json:"searchTime"`
}

type SearchResult struct {
	Kind         string     `json:"kind"`
	Title        string     `json:"title"`
	Link         string     `json:"link"`
	DisplayLink  string     `json:"displayLink"`
	Snippet      string     `json:"snippet"`
	HtmlSnippet  string     `json:"htmlSnippet"`
	FormattedUrl string     `json:"formattedUrl"`
	PageMap      ApiPageMap `json:"pagemap"` // Changed to ApiPageMap
	MetaTags     []MetaTag  `json:"metatags"`
}

// API-specific structs
type ApiPageMap struct {
	CseThumbnail []ApiThumbnail `json:"cse_thumbnail"`
	CseImage     []ImageApi     `json:"cse_image"`
}

type ApiThumbnail struct {
	Src    string `json:"src"`
	Width  string `json:"width"`  // String type for API response
	Height string `json:"height"` // String type for API response
}

// Original structs remain unchanged
type PageMap struct {
	CseThumbnail []Thumbnail `json:"cse_thumbnail"`
	CseImage     []ImageApi  `json:"cse_image"`
}

type Thumbnail struct {
	Src    string `json:"url"`
	Width  int    `json:"width"`  // Remains as int
	Height int    `json:"height"` // Remains as int
}

type ImageApi struct {
	Src string `json:"src"`
}

type MetaTag struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}
