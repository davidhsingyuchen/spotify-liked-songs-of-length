# Spotify Liked Songs of Length

A CLI that prints the names and the URLs of all your Spotify Liked Songs that are of the specified length.

## Usage

1. To make authorized API calls to Spotify, you need to [create a Spotify application](https://developer.spotify.com/my-applications/) first.
    1. During the application creation process, you will be asked to provide a redirect URI. Enter `http://localhost:42000/callback` for that. The value refers to the URL Spotify's authorization server will redirect to after you allow the authorization request. See `redirectURL` in `client.go` for more details.
1. Find the client ID and client secret for your application and export the following environment variables:

    ```sh
    export SPOTIFY_ID=<client-ID-of-your-Spotify-application>
    export SPOTIFY_SECRET=<client-secret-of-your-Spotify-application>
    ```

1. `make build` to compile the binary.
1. `./spotify-liked-songs-of-length --seconds 260` to run the binary.

## Sample Output

```sh
./spotify-liked-songs-of-length
Please visit this URL to grant this app the access to your Liked Songs: https://accounts.spotify.com/authorize?client_id=<redacted>&redirect_uri=http%3A%2F%2Flocalhost%3A42000%2Fcallback&response_type=code&scope=user-library-read&state=show-me-your-music
2024/01/26 16:41:41 Total number of liked songs: 2780
2024/01/26 16:41:41 Getting all liked songs...
2024/01/26 16:42:42 Liked songs of length 4m20s:
南瓜妮歌迷俱樂部 - 山坡上的薩滿: https://open.spotify.com/track/3HMxNxvx91Efe836ZaztRu
白色海岸The White Coast - Shattered Star: https://open.spotify.com/track/3e04i22QhM4cppTMWB4zom
法老 - Afk: https://open.spotify.com/track/4z3yPNaMnZL3Kha12zSZo6
Lu1 - running: https://open.spotify.com/track/22u6Vpwm7gb2ON2aPKMHMv
WeiBird - I Wrote a Song for You: https://open.spotify.com/track/0aIfJQoEtwqQ7TMCZtNU5E
Lu1 - Highlight: https://open.spotify.com/track/1YdqRuWVomFFZpbsU0s7i8
wannasleep - 夜間限定: https://open.spotify.com/track/2496ZNU5RPV0vDjUwlrr94
WONFU - 小春日和: https://open.spotify.com/track/0613EA2Al83HP9LG75Ez9v
The Fur.​ - Movie Star: https://open.spotify.com/track/44TBCNp2SHKqOVurtFtg5C
Gowe - Jazz City Poets: https://open.spotify.com/track/0bKeNKTmH05yYQpBIFo2N4
Justice Der - Close to You: https://open.spotify.com/track/4CTgUidchJ71NULTY39lqk
Hwang Puha - 첫 마음 The First Heart: https://open.spotify.com/track/0rcx5iXlBn7nY47gHGRAvg
一種心情 - 離開雨季: https://open.spotify.com/track/38zg3shtCTg4Tid8qFLceo
DJ Didilong - 台北直直撞: https://open.spotify.com/track/41sJNiLh96NJPH6OGfytyX
```

## TODO

- Adopt [OAuth2 PKCE](https://www.oauth.com/oauth2-servers/pkce/) ([example code](https://github.com/zmb3/spotify/blob/master/examples/authenticate/pkce/pkce.go)), which is more secure.
- Store the access token in the built-in secure storage provided by the OS (e.g., macOS Keychain, Windows Credential Manager, etc.) and only ask the user to re-auth when the token expires.
