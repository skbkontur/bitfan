// Code generated by "bitfanDoc "; DO NOT EDIT
package twitter

import "github.com/vjeantet/bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
  Name:       "twitter",
  ImportPath: "/Users/sodadi/go/src/github.com/vjeantet/bitfan/processors/input-twitter",
  Doc:        "",
  DocShort:   "",
  Options:    &doc.ProcessorOptions{
    Doc:     "",
    Options: []*doc.ProcessorOption{
      &doc.ProcessorOption{
        Name:         "Add_field",
        Alias:        "",
        Doc:          "If this filter is successful, add any arbitrary fields to this event.",
        Required:     false,
        Type:         "hash",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Consumer_key",
        Alias:        "",
        Doc:          "",
        Required:     false,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Consumer_secret",
        Alias:        "",
        Doc:          "",
        Required:     false,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Oauth_token",
        Alias:        "",
        Doc:          "",
        Required:     false,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Oauth_token_secret",
        Alias:        "",
        Doc:          "",
        Required:     false,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Keywords",
        Alias:        "",
        Doc:          "Any keywords to track in the Twitter stream. For multiple keywords,\nuse the syntax [\"foo\", \"bar\"]. There’s a logical OR between each keyword\nstring listed and a logical AND between words separated by spaces per keyword string.\nSee https://dev.twitter.com/streaming/overview/request-parameters#track for more details.",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Follows",
        Alias:        "",
        Doc:          "A comma separated list of user IDs, indicating the users to return statuses for\nin the Twitter stream.\nSee https://dev.twitter.com/streaming/overview/request-parameters#follow for more details.",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Full_tweet",
        Alias:        "",
        Doc:          "Record full tweet object as given to us by the Twitter Streaming API",
        Required:     false,
        Type:         "bool",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Ignore_retweets",
        Alias:        "",
        Doc:          "Lets you ingore the retweets coming out of the Twitter API. Default false",
        Required:     false,
        Type:         "bool",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Languages",
        Alias:        "",
        Doc:          "A list of BCP 47 language identifiers corresponding to any of the languages\nlisted on Twitter’s advanced search page will only return tweets that have been\ndetected as being written in the specified languages",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Locations",
        Alias:        "",
        Doc:          "A comma-separated list of longitude, latitude pairs specifying a set of bounding boxes\nto filter tweets by. See\nhttps://dev.twitter.com/streaming/overview/request-parameters#locations for more details",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Tags",
        Alias:        "",
        Doc:          "If this filter is successful, add arbitrary tags to the event. Tags can be dynamic\nand include parts of the event using the %{field} syntax.",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
    },
  },
  Ports: []*doc.ProcessorPort{},
}
}