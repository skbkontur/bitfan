// Code generated by "bitfanDoc -codec jsonlinesDecoder"; DO NOT EDIT
package jsonlinescodec

import "github.com/vjeantet/bitfan/processors/doc"

func (p *jsonlinesDecoder) Doc() *doc.Codec {
	return &doc.Codec{
  Name:       "jsonlinescodec",
  ImportPath: "/Users/sodadi/go/src/github.com/vjeantet/bitfan/codecs/jsonlinescodec",
  Doc:        "",
  DocShort:   "",
  Options:    &doc.CodecOptions{
    Doc:     "",
    Options: []*doc.CodecOption{
      &doc.CodecOption{
        Name:           "Delimiter",
        Alias:          "",
        Doc:            "Change the delimiter that separates lines",
        Required:       false,
        Type:           "string",
        DefaultValue:   "\"\\n\"",
        PossibleValues: []string{},
        ExampleLS:      "",
      },
    },
  },
}
}