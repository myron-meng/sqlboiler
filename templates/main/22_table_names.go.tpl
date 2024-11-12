{{- if .Table.IsView -}}
{{- else -}}
{{- $alias := .Aliases.Table .Table.Name -}}

{{"\n"}}// TableName returns the physical table name in database.
func (o *{{$alias.UpSingular}}) TableName() string {
    return "{{.Table.Name}}"
}

// SingularName returns the human readable singular name of the model.
func (o *{{$alias.UpSingular}}) SingularName() string {
    return "{{$alias.UpSingular}}"
}

// PluralName returns the human readable plural name of the model.
func (o *{{$alias.UpSingular}}) PluralName() string {
    return "{{$alias.UpPlural}}"
}

// StructName returns the name of the struct type.
func (o *{{$alias.UpSingular}}) StructName() string {
    return "{{$alias.UpSingular}}"
}

// SliceTypeName returns the name of the slice type for this struct.
func (o *{{$alias.UpSingular}}) SliceTypeName() string {
    return "{{$alias.UpSingular}}Slice"
}

// IsNil returns true if the object is nil.
func (o *{{$alias.UpSingular}}) IsNil() bool {
    return o == nil
}

{{end}}