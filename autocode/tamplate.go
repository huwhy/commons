package autocode

var modelTemp = "package {{.Package}}\n\ntype {{camel .Table true}} struct {\n    {{range .Columns}}{{camel .ColumnName true}} {{typeName .DataType .Length}} `json:\"{{camel .ColumnName false}}\"`\n    {{end}}\n}\n\nfunc (m *{{camel .Table true}}) TableName() string {\n\treturn \"{{.Table}}\"\n}\n"
