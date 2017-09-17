// Code generated by "bitfanDoc "; DO NOT EDIT
package httpserverprocessor

import "github.com/vjeantet/bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
  Name:       "httpserverprocessor",
  ImportPath: "github.com/vjeantet/bitfan/processors/input-httpserver",
  Doc:        "Listen and read a http request to build events with it.\n\nProcessor respond with a HTTP code as :\n\n* `202` when request has been accepted, in body : the total number of event created\n* `500` when an error occurs, in body : an error description\n\nUse codecs to process body content as json / csv / lines / json lines / ....\n\nURL is available as http://webhookhost/pluginLabel/URI\n\n* webhookhost is defined by bitfan at startup\n* pluginLabel is defined in pipeline configuration, it's the named processor if you put one, or `input_httpserver` by default\n* URI is defined in plugin configuration (see below)",
  DocShort:   "Reads events from standard input",
  Options:    &doc.ProcessorOptions{
    Doc:     "",
    Options: []*doc.ProcessorOption{
      &doc.ProcessorOption{
        Name:           "processors.CommonOptions",
        Alias:          ",squash",
        Doc:            "",
        Required:       false,
        Type:           "processors.CommonOptions",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "Codec",
        Alias:          "",
        Doc:            "The codec used for input data. Input codecs are a convenient method for decoding\nyour data before it enters the input, without needing a separate filter in your bitfan pipeline\n\nDefault decode http request as plain data, response is json encoded.\nSet multiple codec with role to customize",
        Required:       false,
        Type:           "codec",
        DefaultValue:   "\"plain\"",
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "Uri",
        Alias:          "",
        Doc:            "URI path",
        Required:       false,
        Type:           "string",
        DefaultValue:   "\"events\"",
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "Headers",
        Alias:          "headers",
        Doc:            "Headers to send back into each outgoing response\n@LSExample {\"X-Processor\" => \"bitfan\"}",
        Required:       false,
        Type:           "hash",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "Body",
        Alias:          "body",
        Doc:            "What to send back to client ?",
        Required:       false,
        Type:           "array",
        DefaultValue:   "[\"uuid\"]",
        PossibleValues: []string{},
        ExampleLS:      "",
      },
    },
  },
  Ports: []*doc.ProcessorPort{},
}
}