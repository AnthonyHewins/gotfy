# gotfy

GoTFY is a `ntfy` API client for interacting with NTFY servers.
Push messages to your NTFY server for simple notifications

## Install as a dependency

```shell
go get github.com/AnthonyHewins/gotfy
```

## Example usage

```go
server, _ := url.Parse("https://server.com")

// Create a publisher with default HTTP client
tp, err := gotfy.NewPublisher(server)
if err != nil {
    panic("bad config:"+err.Error())
}

// Or with custom HTTP client
customHTTPClient := http.DefaultClient
tp, err := gotfy.NewPublisher(server, gotfy.WithHTTPClient(customHTTPClient))

// Or with basic authentication
tp, err := gotfy.NewPublisher(server, gotfy.WithAuth("username", "password"))

// Or combine multiple options
tp, err := gotfy.NewPublisher(
    server, 
    gotfy.WithHTTPClient(customHTTPClient),
    gotfy.WithAuth("username", "password"),
)
if err != nil {
    panic("bad config:"+err.Error())
}

pubResp, err := tp.SendMessage(context.Background(), &gotfy.Message{
    Topic:   "topic",
    Message: "message",
    Title: "title",
    Tags:    []string{"emoji1","emoji2","some text"},
    Priority: gotfy.High,
    Actions: []gotfy.ActionButton{
	    Label: "label",
	    Link: "http://link.sh",
	    Clear: true,
    },
    ClickURL: "http://click.com",
    IconURL: "http://icon.com",
    Delay:   time.Minute * 5,
    Email:   "email@domain.com",
})

if err != nil {
    panic("something happened "+err.Error())
}

fmt.Println(pubResp)
// Takes form of:
// type PublishResp struct {
// 	ID      string `json:"id"`      // :"bUhbhgmmbeW0"
// 	Time    int    `json:"time"`    // :1685150791
// 	Expires int    `json:"expires"` // :1685193991
// 	Event   string `json:"event"`   // :"message"
// 	Topic   string `json:"topic"`   // :"TopicName"
// 	Message string `json:"message"` // :"triggered"
// }
```
