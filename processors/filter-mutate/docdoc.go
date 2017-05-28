// Code generated by "bitfanDoc "; DO NOT EDIT
package mutate

import "github.com/vjeantet/bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
  Name:       "mutate",
  ImportPath: "/Users/sodadi/go/src/github.com/vjeantet/bitfan/processors/filter-mutate",
  Doc:        "mutate filter allows to perform general mutations on fields. You can rename, remove, replace, and modify fields in your event.",
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
        Name:         "Add_tag",
        Alias:        "",
        Doc:          "If this filter is successful, add arbitrary tags to the event.\nTags can be dynamic and include parts of the event using the %{field} syntax.",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Convert",
        Alias:        "",
        Doc:          "Convert a field’s value to a different type, like turning a string to an integer.\nIf the field value is an array, all members will be converted. If the field is a hash,\nno action will be taken.\nIf the conversion type is boolean, the acceptable values are:\nTrue: true, t, yes, y, and 1\nFalse: false, f, no, n, and 0\nIf a value other than these is provided, it will pass straight through and log a warning message.\nValid conversion targets are: integer, float, string, and boolean.",
        Required:     false,
        Type:         "hash",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Gsub",
        Alias:        "",
        Doc:          "Convert a string field by applying a regular expression and a replacement. If the field is not a string, no action will be taken.\nThis configuration takes an array consisting of 3 elements per field/substitution.\nBe aware of escaping any backslash in the config file.",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Join",
        Alias:        "",
        Doc:          "Join an array with a separator character. Does nothing on non-array fields",
        Required:     false,
        Type:         "hash",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Lowercase",
        Alias:        "",
        Doc:          "Convert a value to its lowercase equivalent",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Merge",
        Alias:        "",
        Doc:          "Merge two fields of arrays or hashes. String fields will be automatically be converted into an array",
        Required:     false,
        Type:         "hash",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Remove_field",
        Alias:        "",
        Doc:          "If this filter is successful, remove arbitrary fields from this event.",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Remove_tag",
        Alias:        "",
        Doc:          "If this filter is successful, remove arbitrary tags from the event.\nTags can be dynamic and include parts of the event using the %{field} syntax",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Rename",
        Alias:        "",
        Doc:          "Rename key on one or more fields",
        Required:     false,
        Type:         "hash",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Replace",
        Alias:        "",
        Doc:          "Replace a field with a new value. The new value can include %{foo} strings to\nhelp you build a new value from other parts of the event",
        Required:     false,
        Type:         "hash",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Split",
        Alias:        "",
        Doc:          "Split a field to an array using a separator character. Only works on string fields",
        Required:     false,
        Type:         "hash",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Strip",
        Alias:        "",
        Doc:          "Strip whitespace from processors. NOTE: this only works on leading and trailing whitespace",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Update",
        Alias:        "",
        Doc:          "Update an existing field with a new value. If the field does not exist, then no action will be taken",
        Required:     false,
        Type:         "hash",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Uppercase",
        Alias:        "",
        Doc:          "Convert a value to its uppercase equivalent",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
      &doc.ProcessorOption{
        Name:         "Remove_all_but",
        Alias:        "",
        Doc:          "remove all fields, except theses fields (work only with first level fields)",
        Required:     false,
        Type:         "array",
        DefaultValue: nil,
        ExampleLS:    "",
      },
    },
  },
  Ports: []*doc.ProcessorPort{
    &doc.ProcessorPort{
      Default: true,
      Name:    "PORT_SUCCESS",
      Number:  0,
      Doc:     "",
    },
  },
}
}