// Code generated by "bitfanDoc "; DO NOT EDIT
package inputsql

import "github.com/vjeantet/bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
  Name:       "inputsql",
  ImportPath: "/Users/sodadi/go/src/github.com/vjeantet/bitfan/processors/input-sql",
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
        Name:         "Tags",
        Alias:        "",
        Doc:          "If this filter is successful, add arbitrary tags to the event. Tags can be dynamic\nand include parts of the event using the %{field} syntax.",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Type",
        Alias:        "",
        Doc:          "Add a type field to all events handled by this input",
        Required:     false,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Codec",
        Alias:        "",
        Doc:          "The codec used for input data. Input codecs are a convenient method for decoding\nyour data before it enters the input, without needing a separate filter in your bitfan pipeline",
        Required:     false,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Driver",
        Alias:        "driver",
        Doc:          "GOLANG driver class to load, for example, \"mysql\".",
        Required:     true,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "driver => \"mysql\"",
      },
      &doc.ProcessorOption{
        Name:         "EventBy",
        Alias:        "event_by",
        Doc:          "Send an event row by row or one event with all results\npossible values \"row\", \"result\"",
        Required:     false,
        Type:         "string",
        DefaultValue: "\"row\"",
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Statement",
        Alias:        "statement",
        Doc:          "SQL Statement\nWhen there is more than 1 statement, only data from the last one will generate events.",
        Required:     true,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "statement => \"SELECT * FROM mytable\"",
      },
      &doc.ProcessorOption{
        Name:         "Interval",
        Alias:        "interval",
        Doc:          "Set an interval when this processor is used as a input",
        Required:     true,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "interval => \"10\"",
      },
      &doc.ProcessorOption{
        Name:         "ConnectionString",
        Alias:        "connection_string",
        Doc:          "",
        Required:     true,
        Type:         "string",
        DefaultValue: nil,
        ExampleLS:    "connection_string => \"username:password@tcp(192.168.1.2:3306)/mydatabase?charset=utf8\"",
      },
      &doc.ProcessorOption{
        Name:         "Var",
        Alias:        "var",
        Doc:          "You can set variable to be used in Statements by using ${var}.\neach reference will be replaced by the value of the variable found in Statement's content\nThe replacement is case-sensitive.",
        Required:     false,
        Type:         "hash",
        DefaultValue: nil,
        ExampleLS:    "var => {\"hostname\"=>\"myhost\",\"varname\"=>\"varvalue\"}",
      },
      &doc.ProcessorOption{
        Name:         "Target",
        Alias:        "target",
        Doc:          "Define the target field for placing the retrieved data. If this setting is omitted,\nthe data will be stored in the \"data\" field\nSet the value to \".\" to store value to the root (top level) of the event",
        Required:     false,
        Type:         "string",
        DefaultValue: "\"data\"",
        ExampleLS:    "target => \"data\"",
      },
    },
  },
  Ports: []*doc.ProcessorPort{},
}
}