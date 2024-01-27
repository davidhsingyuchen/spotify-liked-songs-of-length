package main

import (
	"fmt"
	"net/http"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

const (
	addr             = "localhost:42000"
	callbackEndpoint = "/callback"
	redirectURL      = "http://" + addr + callbackEndpoint
)

// NewClient returns an authorized Spotify client via implementing OAuth 2.0 Authorization Code Flow.
func NewClient() *spotify.Client {
	state := genState()
	auth := spotifyauth.New(
		spotifyauth.WithRedirectURL(redirectURL),
		spotifyauth.WithScopes(spotifyauth.ScopeUserLibraryRead),
	)
	ch := make(chan *spotify.Client)
	server := newServer(auth, state, ch)
	go server.ListenAndServe()

	url := auth.AuthURL(state)
	fmt.Printf("Please visit this URL to grant this app the access to your Liked Songs: %s\n", url)

	client := <-ch
	// We probably don't need to use Shutdown() here
	// because we've already got the client,
	// which is all we want from this server.
	server.Close()
	return client
}

// We can probably just use a hard-coded string here because:
//   - The redirection callback server is hosted on localhost, which is not accessible to the public.
//   - The server will only be there for a very short time span
//     (i.e., it will be shut down right after the auth ends).
//   - Even if someone manages to call the redirection URL with their own authorization code in that short time span,
//     the only consequence is that we'll see the their liked songs that are of the length specified by us,
//     which can be kind of interesting I suppose.
func genState() string {
	return "show-me-your-music"
}

// newServer returns an HTTP server that contains a callback handler for OAuth.
// After the callback is called, it passes the authorized client to ch.
func newServer(auth *spotifyauth.Authenticator, state string, ch chan<- *spotify.Client) *http.Server {
	server := http.NewServeMux()
	server.HandleFunc(callbackEndpoint, func(w http.ResponseWriter, r *http.Request) {
		redirectHandler(w, r, auth, state, ch)
	})
	return &http.Server{Addr: addr, Handler: server}
}

func redirectHandler(w http.ResponseWriter, r *http.Request, auth *spotifyauth.Authenticator, state string, ch chan<- *spotify.Client) {
	token, err := auth.Token(r.Context(), state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusNotFound)
		return
	}
	ch <- spotify.New(auth.Client(r.Context(), token))
}
