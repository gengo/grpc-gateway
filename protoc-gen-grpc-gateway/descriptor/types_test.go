package descriptor

import (
	"testing"

	"github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
)

func TestGoPackageStandard(t *testing.T) {
	for _, spec := range []struct {
		pkg  GoPackage
		want bool
	}{
		{
			pkg:  GoPackage{Path: "fmt", Name: "fmt"},
			want: true,
		},
		{
			pkg:  GoPackage{Path: "encoding/json", Name: "json"},
			want: true,
		},
		{
			pkg:  GoPackage{Path: "github.com/golang/protobuf/jsonpb", Name: "jsonpb"},
			want: false,
		},
		{
			pkg:  GoPackage{Path: "golang.org/x/net/context", Name: "context"},
			want: false,
		},
		{
			pkg:  GoPackage{Path: "github.com/grpc-ecosystem/grpc-gateway", Name: "main"},
			want: false,
		},
		{
			pkg:  GoPackage{Path: "github.com/google/googleapis/google/api/http.pb", Name: "http_pb", Alias: "htpb"},
			want: false,
		},
	} {
		if got, want := spec.pkg.Standard(), spec.want; got != want {
			t.Errorf("%#v.Standard() = %v; want %v", spec.pkg, got, want)
		}
	}
}

func TestGoPackageString(t *testing.T) {
	for _, spec := range []struct {
		pkg  GoPackage
		want string
	}{
		{
			pkg:  GoPackage{Path: "fmt", Name: "fmt"},
			want: `"fmt"`,
		},
		{
			pkg:  GoPackage{Path: "encoding/json", Name: "json"},
			want: `"encoding/json"`,
		},
		{
			pkg:  GoPackage{Path: "github.com/golang/protobuf/jsonpb", Name: "jsonpb"},
			want: `"github.com/golang/protobuf/jsonpb"`,
		},
		{
			pkg:  GoPackage{Path: "golang.org/x/net/context", Name: "context"},
			want: `"golang.org/x/net/context"`,
		},
		{
			pkg:  GoPackage{Path: "github.com/grpc-ecosystem/grpc-gateway", Name: "main"},
			want: `"github.com/grpc-ecosystem/grpc-gateway"`,
		},
		{
			pkg:  GoPackage{Path: "github.com/google/googleapis/google/api/http.pb", Name: "http_pb", Alias: "htpb"},
			want: `htpb "github.com/google/googleapis/google/api/http.pb"`,
		},
	} {
		if got, want := spec.pkg.String(), spec.want; got != want {
			t.Errorf("%#v.String() = %q; want %q", spec.pkg, got, want)
		}
	}
}

func TestIsOneof(t *testing.T) {
	const src = `
		name: 'M'
		field <
			name: 'oneof_option'
			label: LABEL_OPTIONAL
			type: TYPE_STRING
			number: 1
			oneof_index: 0
		>
		oneof_decl <
			name: 'oneof_field'
		>
	`
	var d descriptor.DescriptorProto
	if err := proto.UnmarshalText(src, &d); err != nil {
		t.Fatalf("proto.UnmarshalText(%s, &d) failed with %v; want success", src, err)
	}
	m := &Message{DescriptorProto: &d}
	f := &Field{
		Message:              m,
		FieldDescriptorProto: d.Field[0],
	}

	if got, want := f.IsOneof(), true; got != want {
		t.Errorf("f.IsOneof() = %v; want %v", got, want)
	}

	f.OneofIndex = nil
	if got, want := f.IsOneof(), false; got != want {
		t.Errorf("f.IsOneof() = %v; want %v", got, want)
	}
}

func TestGoOneofName(t *testing.T) {
	const src = `
		name: 'M'
		field <
			name: 'oneof_option'
			label: LABEL_OPTIONAL
			type: TYPE_STRING
			number: 1
			oneof_index: 0
		>
		oneof_decl <
			name: 'oneof_field'
		>
	`
	var d descriptor.DescriptorProto
	if err := proto.UnmarshalText(src, &d); err != nil {
		t.Fatalf("proto.UnmarshalText(%s, &d) failed with %v; want success", src, err)
	}
	m := &Message{DescriptorProto: &d}
	f := &Field{
		Message:              m,
		FieldDescriptorProto: d.Field[0],
	}

	name, err := f.GoOneofName()
	if err != nil {
		t.Errorf("f.GoOneofName() failed with %v; want success", err)
	}
	if got, want := name, "OneofField"; got != want {
		t.Errorf("name = %q; want %q", got, want)
	}

	f.OneofIndex = nil
	if name, err := f.GoOneofName(); err == nil {
		t.Errorf("f.GoOneofName() = %q; want failure", name)
	}
}

// nestedExamples returns a pair of message descriptors for testing
// field paths to nested fields with them.
func nestedExamples(t *testing.T) (proto3, proto2 *Message) {
	var fds []*descriptor.FileDescriptorProto
	for _, src := range []string{
		`
		name: 'example.proto'
		package: 'example'
		message_type <
			name: 'Nest'
			field <
				name: 'nest2_field'
				label: LABEL_OPTIONAL
				type: TYPE_MESSAGE
				type_name: 'Nest2'
				number: 1
			>
			field <
				name: 'terminal_field'
				label: LABEL_OPTIONAL
				type: TYPE_STRING
				number: 2
			>
			field <
				name: 'oneof_opt'
				label: LABEL_OPTIONAL
				type: TYPE_INT32
				number: 3
				oneof_index: 0
			>
			oneof_decl <
				name: 'oneof_field'
			>
		>
		syntax: "proto3"
		`, `
		name: 'another.proto'
		package: 'example'
		message_type <
			name: 'Nest2'
			field <
				name: 'nest_field'
				label: LABEL_OPTIONAL
				type: TYPE_MESSAGE
				type_name: 'Nest'
				number: 1
			>
			field <
				name: 'terminal_field'
				label: LABEL_OPTIONAL
				type: TYPE_STRING
				number: 2
			>
			# NOTE: oneof field is not supported in proto2 syntax
		>
		syntax: "proto2"
		`,
	} {
		var fd descriptor.FileDescriptorProto
		if err := proto.UnmarshalText(src, &fd); err != nil {
			t.Fatalf("proto.UnmarshalText(%s, &fd) failed with %v; want success", src, err)
		}
		fds = append(fds, &fd)
	}
	proto3 = &Message{
		DescriptorProto: fds[0].MessageType[0],
		Fields: []*Field{
			{FieldDescriptorProto: fds[0].MessageType[0].Field[0]},
			{FieldDescriptorProto: fds[0].MessageType[0].Field[1]},
			{FieldDescriptorProto: fds[0].MessageType[0].Field[2]},
		},
	}
	proto2 = &Message{
		DescriptorProto: fds[1].MessageType[0],
		Fields: []*Field{
			{FieldDescriptorProto: fds[1].MessageType[0].Field[0]},
			{FieldDescriptorProto: fds[1].MessageType[0].Field[1]},
		},
	}
	file1 := &File{
		FileDescriptorProto: fds[0],
		GoPkg:               GoPackage{Path: "example", Name: "example"},
		Messages:            []*Message{proto3},
	}
	file2 := &File{
		FileDescriptorProto: fds[1],
		GoPkg:               GoPackage{Path: "example", Name: "example"},
		Messages:            []*Message{proto2},
	}
	crossLinkFixture(file1)
	crossLinkFixture(file2)

	return proto3, proto2
}

func TestFieldPath(t *testing.T) {
	proto3, proto2 := nestedExamples(t)

	c1 := FieldPathComponent{
		Name:   "nest_field",
		Target: proto2.Fields[0],
	}
	if got, want := c1.LHS(), "GetNestField()"; got != want {
		t.Errorf("c1.LHS() = %q; want %q", got, want)
	}
	if got, want := c1.RHS(), "NestField"; got != want {
		t.Errorf("c1.RHS() = %q; want %q", got, want)
	}

	c2 := FieldPathComponent{
		Name:   "nest2_field",
		Target: proto3.Fields[0],
	}
	if got, want := c2.LHS(), "Nest2Field"; got != want {
		t.Errorf("c2.LHS() = %q; want %q", got, want)
	}
	if got, want := c2.LHS(), "Nest2Field"; got != want {
		t.Errorf("c2.LHS() = %q; want %q", got, want)
	}

	fp := FieldPath{
		c1, c2, c1, FieldPathComponent{
			Name:   "terminal_field",
			Target: proto3.Fields[1],
		},
	}
	if got, want := fp.RHS("resp"), "resp.GetNestField().Nest2Field.GetNestField().TerminalField"; got != want {
		t.Errorf("fp.RHS(%q) = %q; want %q", "resp", got, want)
	}

	fp2 := FieldPath{
		c2, c1, c2, FieldPathComponent{
			Name:   "terminal_field",
			Target: proto2.Fields[1],
		},
	}
	if got, want := fp2.RHS("resp"), "resp.Nest2Field.GetNestField().Nest2Field.TerminalField"; got != want {
		t.Errorf("fp2.RHS(%q) = %q; want %q", "resp", got, want)
	}

	var fpEmpty FieldPath
	if got, want := fpEmpty.RHS("resp"), "resp"; got != want {
		t.Errorf("fpEmpty.RHS(%q) = %q; want %q", "resp", got, want)
	}
}

func TestOneofFieldPath(t *testing.T) {
	proto3, proto2 := nestedExamples(t)

	c1 := FieldPathComponent{
		Name:   "nest_field",
		Target: proto2.Fields[0],
	}
	c2 := FieldPathComponent{
		Name:   "nest2_field",
		Target: proto3.Fields[0],
	}

	fp := FieldPath{
		c1, c2, c1, FieldPathComponent{
			Name:   "oneof_opt",
			Target: proto3.Fields[2],
		},
	}
	if got, want := fp.RHS("resp"), "resp.GetNestField().Nest2Field.GetNestField().alloc_OneofOpt().OneofOpt"; got != want {
		t.Errorf("fp.RHS(%q) = %q; want %q", "resp", got, want)
	}
}
