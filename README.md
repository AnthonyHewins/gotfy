# gotfy

GoTFY is a `ntfy` client for interacting with NTFY servers

## Install as a dependency

```shell
go get github.com/AnthonyHewins/gotfy
```

## Example usage

```go
server, _ := url.Parse("http://server.com")
customHTTPClient := http.DefaultClient

tp, err := gotfy.NewTopicPublisher(server, customHTTPClient)
if err != nil {
    panic("bad config:"+err.Error())
}

tp.SendMessage(&gotfy.Message{
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