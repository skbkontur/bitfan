// Code generated by "bitfanDoc "; DO NOT EDIT
package httpoutput

import "github.com/vjeantet/bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
  Name:       "httpoutput",
  ImportPath: "/Users/sodadi/go/src/github.com/vjeantet/bitfan/processors/output-http",
  Doc:        "",
  DocShort:   "",
  Options:    &doc.ProcessorOptions{
    Doc:     "",
    Options: []*doc.ProcessorOption{
      &doc.ProcessorOption{
        Name:         "AddField",
        Alias:        "add_field",
        Doc:          "Add a field to an event. Default value is {}",
        Required:     false,
        Type:         "hash",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "URL",
        Alias:        "url",
        Doc:          "This output lets you send events to a generic HTTP(S) endpoint\nThis setting can be dynamic using the %{foo} syntax.",
        Required:     true,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Headers",
        Alias:        "headers",
        Doc:          "Custom headers to use format is headers => {\"X-My-Header\", \"%{host}\"}. Default value is {}\nThis setting can be dynamic using the %{foo} syntax.",
        Required:     false,
        Type:         "hash",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "HTTPMethod",
        Alias:        "http_method",
        Doc:          "The HTTP Verb. One of \"put\", \"post\", \"patch\", \"delete\", \"get\", \"head\". Default value is \"post\"",
        Required:     false,
        Type:         "string",
        DefaultValue: "\"post\"",
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "KeepAlive",
        Alias:        "keepalive",
        Doc:          "Turn this on to enable HTTP keepalive support. Default value is true",
        Required:     false,
        Type:         "bool",
        DefaultValue: "true",
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "PoolMax",
        Alias:        "pool_max",
        Doc:          "Max number of concurrent connections. Default value is 1",
        Required:     false,
        Type:         "int",
        DefaultValue: "1",
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "ConnectTimeout",
        Alias:        "connect_timeout",
        Doc:          "Timeout (in seconds) to wait for a connection to be established. Default value is 10",
        Required:     false,
        Type:         "uint",
        DefaultValue: "5",
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "RequestTimeout",
        Alias:        "request_timeout",
        Doc:          "Timeout (in seconds) for the entire request. Default value is 60",
        Required:     false,
        Type:         "uint",
        DefaultValue: "30",
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Format",
        Alias:        "format",
        Doc:          "Set the format of the http body. Now supports only \"json_lines\"",
        Required:     false,
        Type:         "string",
        DefaultValue: "\"json_lines\"",
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "RetryableCodes",
        Alias:        "retryable_codes",
        Doc:          "If encountered as response codes this plugin will retry these requests",
        Required:     false,
        Type:         "array",
        DefaultValue: "[429, 500, 502, 503, 504]",
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "IgnorableCodes",
        Alias:        "ignorable_codes",
        Doc:          "If you would like to consider some non-2xx codes to be successes\nenumerate them here. Responses returning these codes will be considered successes",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "BatchInterval",
        Alias:        "batch_interval",
        Doc:          "",
        Required:     false,
        Type:         "uint",
        DefaultValue: "5",
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "BatchSize",
        Alias:        "batch_size",
        Doc:          "",
        Required:     false,
        Type:         "uint",
        DefaultValue: "100",
        ExampleLS:    "",
      },
    },
  },
  Ports: []*doc.ProcessorPort{},
}
}