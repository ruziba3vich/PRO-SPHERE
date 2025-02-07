// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"net/url"
// )

// // Replace these with your credentials provided by ProID
// const clientID = "67"
// const clientSecret = "K7mMNszuOjgHLB2iXOFTcn0n00Z67507135cb0c5"
// const redirectURI = "http://localhost:7777/v1/auth/oauth/callback"

// // ProID Authorization and Token URLs
// const authorizationURL = "https://id.sfere.pro/oauth"
// const tokenURL = "https://api.id.sfere.pro/api/v2/oauth/token"

// func main() {
// 	// Create a simple HTTP server to handle the callback and initiate OAuth flow
// 	http.HandleFunc("/v1/auth/oauth/callback", handleCallback)
// 	http.HandleFunc("/v1/auth/oauth/start", startOAuthFlow)
// 	//http://localhost:7777/v1/auth/oauth/start
// 	//http://localhost:7777/v1/auth/oauth/callback
// 	// Start the HTTP server
// 	fmt.Println("Server is running on http://localhost:7777")
// 	log.Fatal(http.ListenAndServe(":7777", nil))
// }

// // handleCallback is called when ProID redirects to your callback URL with the authorization code.
// func handleCallback(w http.ResponseWriter, r *http.Request) {
// 	// Get the authorization code from the query parameter
// 	code := r.URL.Query().Get("code")
// 	if code == "" {
// 		http.Error(w, "Code is missing in the request", http.StatusBadRequest)
// 		return
// 	}
// 	fmt.Println(code)
// 	// Exchange the authorization code for an access token
// 	token, err := exchangeCodeForToken(code)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Error exchanging code for token: %v", err), http.StatusInternalServerError)
// 		return
// 	}
// 	// newUuid := uuid.NewString()
// 	// Print the access token
// 	// fmt.Fprintf(w, "Successfully authenticated! Access token: %s", token)
// 	http.Redirect(w, r, "https://sfere.uz?code="+token, http.StatusFound)
// }

// // startOAuthFlow sends the user to the ProID authorization URL.
// func startOAuthFlow(w http.ResponseWriter, r *http.Request) {
// 	// Prepare the authorization URL
// 	authURL := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&response_type=code",
// 		authorizationURL, clientID, url.QueryEscape(redirectURI))
// 	// fmt.Println(authURL)
// 	// Redirect the user to the ProID authorization URL
// 	fmt.Println(authURL)
// 	http.Redirect(w, r, authURL, http.StatusFound)
// }

// // exchangeCodeForToken exchanges the authorization code for an access token from ProID.
// func exchangeCodeForToken(code string) (string, error) {
// 	// Prepare the form data for the POST request
// 	data := url.Values{}
// 	data.Set("grant_type", "authorization_code")
// 	data.Set("client_id", clientID)
// 	data.Set("client_secret", clientSecret)
// 	data.Set("redirect_uri", redirectURI)
// 	data.Set("code", code)

// 	// Make the HTTP POST request to exchange the code for a token
// 	resp, err := http.PostForm(tokenURL, data)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to send token request: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	// Read the response body
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to read response body: %v", err)
// 	}

// 	// Check if the response is successful
// 	if resp.StatusCode != http.StatusOK {
// 		return "", fmt.Errorf("token request failed with status %d: %s", resp.StatusCode, body)
// 	}

// 	// Parse the JSON response to extract the access token
// 	var response map[string]interface{}
// 	if err := json.Unmarshal(body, &response); err != nil {
// 		return "", fmt.Errorf("failed to parse token response: %v", err)
// 	}
// 	fmt.Println(response)
// 	// Return the access token
// 	expiresAt, ok := response["expires_in"].(string)
// 	if ok {
// 		fmt.Print(expiresAt)
// 	}
// 	token, ok := response["access_token"].(string)
// 	if !ok {
// 		return "", fmt.Errorf("access token not found in the response")
// 	}

//		return token, nil
//	}
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadFavicon(url, filename string) error {
	// Get the favicon file
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch favicon: %v", err)
	}
	defer resp.Body.Close()

	// Create the file on your server
	outFile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer outFile.Close()

	// Write the file data
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save favicon: %v", err)
	}

	return nil
}

func main() {
	err := downloadFavicon("https://www.make.com/favicon.ico", "./static/https://www.make.com/favicon.ico")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Favicon saved successfully!")
	}
}
