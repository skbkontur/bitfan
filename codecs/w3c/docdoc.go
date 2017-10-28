// Code generated by "bitfanDoc -codec w3c"; DO NOT EDIT
package w3ccodec

import "github.com/vjeantet/bitfan/processors/doc"

func Doc() *doc.Codec {
	return &doc.Codec{
  Name:       "w3c",
  PkgName:    "w3ccodec",
  ImportPath: "github.com/vjeantet/bitfan/codecs/w3c",
  Doc:        "Parses comma-separated value data into individual fields",
  DocShort:   "",
  Decoder:    &doc.Decoder{
    Doc:     "",
    Options: &doc.CodecOptions{
      Doc:     "Parses comma-separated value data into individual fields",
      Options: []*doc.CodecOption{
        &doc.CodecOption{
          Name:           "Separator",
          Alias:          "separator",
          Doc:            "Define the column separator value. If this is not specified, the default is a whitespace. Optional",
          Required:       false,
          Type:           "string",
          DefaultValue:   "\" \"",
          PossibleValues: []string{},
          ExampleLS:      "",
        },
        &doc.CodecOption{
          Name:           "AutogenerateColumnNames",
          Alias:          "autogenerate_column_names",
          Doc:            "Define whether column names should autogenerated or not. Defaults to true.\nIf set to false, columns not having a header specified will not be parsed.",
          Required:       false,
          Type:           "bool",
          DefaultValue:   "true",
          PossibleValues: []string{},
          ExampleLS:      "",
        },
        &doc.CodecOption{
          Name:           "QuoteChar",
          Alias:          "quote_char",
          Doc:            "Define the character used to quote fields. If this is not specified the default is a double quote \". Optional.",
          Required:       false,
          Type:           "string",
          DefaultValue:   "\"\\\"\"",
          PossibleValues: []string{},
          ExampleLS:      "",
        },
        &doc.CodecOption{
          Name:           "Columns",
          Alias:          "columns",
          Doc:            "Define a list of column names\n\nIf columns is not configured, or there are not enough columns specified,\nthe default column names are \"column1\", \"column2\", etc.\n\nIn the case that there are more columns in the data than specified in this column\nlist, extra columns will be auto-numbered:\n(e.g. \"user_defined_1\", \"user_defined_2\", \"column3\", \"column4\", etc.)",
          Required:       false,
          Type:           "array",
          DefaultValue:   nil,
          PossibleValues: []string{},
          ExampleLS:      "",
        },
        &doc.CodecOption{
          Name:           "Comment",
          Alias:          "comment",
          Doc:            "Define the comment character.\nLines beginning with the Comment character without preceding whitespace are ignored.",
          Required:       false,
          Type:           "string",
          DefaultValue:   "\"#\"",
          PossibleValues: []string{},
          ExampleLS:      "",
        },
      },
    },
  },
  Encoder: (*doc.Encoder)(nil),
}
}