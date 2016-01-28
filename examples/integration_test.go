package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"

	"golang.org/x/net/context"

	gw "github.com/gengo/grpc-gateway/examples/examplepb"
	server "github.com/gengo/grpc-gateway/examples/server"
	sub "github.com/gengo/grpc-gateway/examples/sub"
	"github.com/gengo/grpc-gateway/runtime"
	"github.com/golang/protobuf/proto"
)

func TestIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
		return
	}

	go func() {
		if err := server.Run(); err != nil {
			t.Errorf("server.Run() failed with %v; want success", err)
			return
		}
	}()
	go func() {
		if err := Run(":8080"); err != nil {
			t.Errorf("gw.Run() failed with %v; want success", err)
			return
		}
	}()

	time.Sleep(100 * time.Millisecond)
	testEcho(t, 8080, "application/json")
	testEchoBody(t)
	testABECreate(t)
	testABECreateBody(t)
	testABEBulkCreate(t)
	testABELookup(t)
	testABEList(t)
	testAdditionalBindings(t)
	testABEProto2Echo(t)
	testABEProto2EchoBody(t)
	testABEProto2BulkEcho(t)

	go func() {
		if err := Run(
			":8081",
			runtime.WithForwardResponseOption(
				func(_ context.Context, w http.ResponseWriter, _ proto.Message) error {
					w.Header().Set("Content-Type", "application/vnd.docker.plugins.v1.1+json")
					return nil
				},
			),
		); err != nil {
			t.Errorf("gw.Run() failed with %v; want success", err)
			return
		}
	}()

	testEcho(t, 8081, "application/vnd.docker.plugins.v1.1+json")
}

func testEcho(t *testing.T, port int, contentType string) {
	url := fmt.Sprintf("http://localhost:%d/v1/example/echo/myid", port)
	resp, err := http.Post(url, "application/json", strings.NewReader("{}"))
	if err != nil {
		t.Errorf("http.Post(%q) failed with %v; want success", url, err)
		return
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("iotuil.ReadAll(resp.Body) failed with %v; want success", err)
		return
	}

	if got, want := resp.StatusCode, http.StatusOK; got != want {
		t.Errorf("resp.StatusCode = %d; want %d", got, want)
		t.Logf("%s", buf)
	}

	var msg gw.SimpleMessage
	if err := json.Unmarshal(buf, &msg); err != nil {
		t.Errorf("json.Unmarshal(%s, &msg) failed with %v; want success", buf, err)
		return
	}
	if got, want := msg.Id, "myid"; got != want {
		t.Errorf("msg.Id = %q; want %q", got, want)
	}

	if value := resp.Header.Get("Content-Type"); value != contentType {
		t.Errorf("Content-Type was %s, wanted %s", value, contentType)
	}
}

func testEchoBody(t *testing.T) {
	sent := gw.SimpleMessage{Id: "example"}
	buf, err := json.Marshal(sent)
	if err != nil {
		t.Fatalf("json.Marshal(%#v) failed with %v; want success", sent, err)
	}

	url := "http://localhost:8080/v1/example/echo_body"
	resp, err := http.Post(url, "", bytes.NewReader(buf))
	if err != nil {
		t.Errorf("http.Post(%q) failed with %v; want success", url, err)
		return
	}
	defer resp.Body.Close()
	buf, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("iotuil.ReadAll(resp.Body) failed with %v; want success", err)
		return
	}

	if got, want := resp.StatusCode, http.StatusOK; got != want {
		t.Errorf("resp.StatusCode = %d; want %d", got, want)
		t.Logf("%s", buf)
	}

	var received gw.SimpleMessage
	if err := json.Unmarshal(buf, &received); err != nil {
		t.Errorf("json.Unmarshal(%s, &msg) failed with %v; want success", buf, err)
		return
	}
	if got, want := received, sent; !reflect.DeepEqual(got, want) {
		t.Errorf("msg.Id = %q; want %q", got, want)
	}
}

func testABECreate(t *testing.T) {
	want := gw.ABitOfEverything{
		FloatValue:    1.5,
		DoubleValue:   2.5,
		Int64Value:    4294967296,
		Uint64Value:   9223372036854775807,
		Int32Value:    -2147483648,
		Fixed64Value:  9223372036854775807,
		Fixed32Value:  4294967295,
		BoolValue:     true,
		StringValue:   "strprefix/foo",
		Uint32Value:   4294967295,
		BytesValue:    []byte{0, 1, 2, 3, 4, 127, 128, 129, 251, 252, 253, 254, 255},
		Sfixed32Value: 2147483647,
		Sfixed64Value: -4611686018427387904,
		Sint32Value:   2147483647,
		Sint64Value:   4611686018427387903,
	}
	bytes := url.QueryEscape(string(want.BytesValue))
	url := fmt.Sprintf("http://localhost:8080/v1/example/a_bit_of_everything/%f/%f/%d/separator/%d/%d/%d/%d/%v/%s/%d/%s/%d/%d/%d/%d", want.FloatValue, want.DoubleValue, want.Int64Value, want.Uint64Value, want.Int32Value, want.Fixed64Value, want.Fixed32Value, want.BoolValue, want.StringValue, want.Uint32Value, bytes, want.Sfixed32Value, want.Sfixed64Value, want.Sint32Value, want.Sint64Value)

	resp, err := http.Post(url, "application/json", strings.NewReader("{}"))
	if err != nil {
		t.Errorf("http.Post(%q) failed with %v; want success", url, err)
		return
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("iotuil.ReadAll(resp.Body) failed with %v; want success", err)
		return
	}

	if got, want := resp.StatusCode, http.StatusOK; got != want {
		t.Errorf("resp.StatusCode = %d; want %d", got, want)
		t.Logf("%s", buf)
	}

	var msg gw.ABitOfEverything
	if err := json.Unmarshal(buf, &msg); err != nil {
		t.Errorf("json.Unmarshal(%s, &msg) failed with %v; want success", buf, err)
		return
	}
	if msg.Uuid == "" {
		t.Error("msg.Uuid is empty; want not empty")
	}
	msg.Uuid = ""
	if got := msg; !reflect.DeepEqual(got, want) {
		t.Errorf("msg= %v; want %v", &got, &want)
	}
}

func testABECreateBody(t *testing.T) {
	want := gw.ABitOfEverything{
		FloatValue:    1.5,
		DoubleValue:   2.5,
		Int64Value:    4294967296,
		Uint64Value:   9223372036854775807,
		Int32Value:    -2147483648,
		Fixed64Value:  9223372036854775807,
		Fixed32Value:  4294967295,
		BoolValue:     true,
		StringValue:   "strprefix/foo",
		Uint32Value:   4294967295,
		BytesValue:    []byte{0, 1, 2, 3, 4, 127, 128, 129, 251, 252, 253, 254, 255},
		Sfixed32Value: 2147483647,
		Sfixed64Value: -4611686018427387904,
		Sint32Value:   2147483647,
		Sint64Value:   4611686018427387903,

		Nested: []*gw.ABitOfEverything_Nested{
			{
				Name:   "bar",
				Amount: 10,
			},
			{
				Name:   "baz",
				Amount: 20,
			},
		},
	}
	url := "http://localhost:8080/v1/example/a_bit_of_everything"
	buf, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("json.Marshal(%#v) failed with %v; want success", want, err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(buf))
	if err != nil {
		t.Errorf("http.Post(%q) failed with %v; want success", url, err)
		return
	}
	defer resp.Body.Close()
	buf, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("iotuil.ReadAll(resp.Body) failed with %v; want success", err)
		return
	}

	if got, want := resp.StatusCode, http.StatusOK; got != want {
		t.Errorf("resp.StatusCode = %d; want %d", got, want)
		t.Logf("%s", buf)
	}

	var msg gw.ABitOfEverything
	if err := json.Unmarshal(buf, &msg); err != nil {
		t.Errorf("json.Unmarshal(%s, &msg) failed with %v; want success", buf, err)
		return
	}
	if msg.Uuid == "" {
		t.Error("msg.Uuid is empty; want not empty")
	}
	msg.Uuid = ""
	if got := msg; !reflect.DeepEqual(got, want) {
		t.Errorf("msg= %v; want %v", &got, &want)
	}
}

func testABEBulkCreate(t *testing.T) {
	r, w := io.Pipe()
	go func(w io.WriteCloser) {
		defer func() {
			if cerr := w.Close(); cerr != nil {
				t.Errorf("w.Close() failed with %v; want success", cerr)
			}
		}()
		for _, val := range []string{
			"foo", "bar", "baz", "qux", "quux",
		} {
			want := gw.ABitOfEverything{
				FloatValue:    1.5,
				DoubleValue:   2.5,
				Int64Value:    4294967296,
				Uint64Value:   9223372036854775807,
				Int32Value:    -2147483648,
				Fixed64Value:  9223372036854775807,
				Fixed32Value:  4294967295,
				BoolValue:     true,
				StringValue:   fmt.Sprintf("strprefix/%s", val),
				Uint32Value:   4294967295,
				BytesValue:    []byte{0, 1, 2, 3, 4, 127, 128, 129, 251, 252, 253, 254, 255},
				Sfixed32Value: 2147483647,
				Sfixed64Value: -4611686018427387904,
				Sint32Value:   2147483647,
				Sint64Value:   4611686018427387903,

				Nested: []*gw.ABitOfEverything_Nested{
					{
						Name:   "hoge",
						Amount: 10,
					},
					{
						Name:   "fuga",
						Amount: 20,
					},
				},
			}
			buf, err := json.Marshal(want)
			if err != nil {
				t.Fatalf("json.Marshal(%#v) failed with %v; want success", want, err)
			}
			if _, err := w.Write(buf); err != nil {
				t.Errorf("w.Write(%s) failed with %v; want success", buf, err)
				return
			}
			if _, err := io.WriteString(w, "\n"); err != nil {
				t.Errorf("w.Write(%s) failed with %v; want success", buf, err)
				return
			}
		}
	}(w)
	url := "http://localhost:8080/v1/example/a_bit_of_everything/bulk"
	resp, err := http.Post(url, "application/json", r)
	if err != nil {
		t.Errorf("http.Post(%q) failed with %v; want success", url, err)
		return
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("iotuil.ReadAll(resp.Body) failed with %v; want success", err)
		return
	}

	if got, want := resp.StatusCode, http.StatusOK; got != want {
		t.Errorf("resp.StatusCode = %d; want %d", got, want)
		t.Logf("%s", buf)
	}

	var msg gw.EmptyMessage
	if err := json.Unmarshal(buf, &msg); err != nil {
		t.Errorf("json.Unmarshal(%s, &msg) failed with %v; want success", buf, err)
		return
	}
}

func testABELookup(t *testing.T) {
	url := "http://localhost:8080/v1/example/a_bit_of_everything"
	cresp, err := http.Post(url, "application/json", strings.NewReader(`
		{"bool_value": true, "string_value": "strprefix/example"}
	`))
	if err != nil {
		t.Errorf("http.Post(%q) failed with %v; want success", url, err)
		return
	}
	defer cresp.Body.Close()
	buf, err := ioutil.ReadAll(cresp.Body)
	if err != nil {
		t.Errorf("iotuil.ReadAll(cresp.Body) failed with %v; want success", err)
		return
	}
	if got, want := cresp.StatusCode, http.StatusOK; got != want {
		t.Errorf("resp.StatusCode = %d; want %d", got, want)
		t.Logf("%s", buf)
		return
	}

	var want gw.ABitOfEverything
	if err := json.Unmarshal(buf, &want); err != nil {
		t.Errorf("json.Unmarshal(%s, &want) failed with %v; want success", buf, err)
		return
	}

	url = fmt.Sprintf("%s/%s", url, want.Uuid)
	resp, err := http.Get(url)
	if err != nil {
		t.Errorf("http.Get(%q) failed with %v; want success", url, err)
		return
	}
	defer resp.Body.Close()

	buf, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("ioutil.ReadAll(resp.Body) failed with %v; want success", err)
		return
	}

	var msg gw.ABitOfEverything
	if err := json.Unmarshal(buf, &msg); err != nil {
		t.Errorf("json.Unmarshal(%s, &msg) failed with %v; want success", buf, err)
		return
	}
	if got := msg; !reflect.DeepEqual(got, want) {
		t.Errorf("msg= %v; want %v", &got, &want)
	}
}

func testABEList(t *testing.T) {
	url := "http://localhost:8080/v1/example/a_bit_of_everything"
	resp, err := http.Get(url)
	if err != nil {
		t.Errorf("http.Get(%q) failed with %v; want success", url, err)
		return
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	var i int
	for i = 0; ; i++ {
		var msg gw.ABitOfEverything
		err := dec.Decode(&msg)
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Errorf("dec.Decode(&msg) failed with %v; want success; i = %d", err, i)
		}
	}
	if i <= 0 {
		t.Errorf("i == %d; want > 0", i)
	}
}

func testAdditionalBindings(t *testing.T) {
	for i, f := range []func() *http.Response{
		func() *http.Response {
			url := "http://localhost:8080/v1/example/a_bit_of_everything/echo/hello"
			resp, err := http.Get(url)
			if err != nil {
				t.Errorf("http.Get(%q) failed with %v; want success", url, err)
				return nil
			}
			return resp
		},
		func() *http.Response {
			url := "http://localhost:8080/v2/example/echo"
			resp, err := http.Post(url, "application/json", strings.NewReader(`"hello"`))
			if err != nil {
				t.Errorf("http.Post(%q, %q, %q) failed with %v; want success", url, "application/json", `"hello"`, err)
				return nil
			}
			return resp
		},
		func() *http.Response {
			url := "http://localhost:8080/v2/example/echo?value=hello"
			resp, err := http.Get(url)
			if err != nil {
				t.Errorf("http.Get(%q) failed with %v; want success", url, err)
				return nil
			}
			return resp
		},
	} {
		resp := f()
		if resp == nil {
			continue
		}

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("iotuil.ReadAll(resp.Body) failed with %v; want success; i=%d", err, i)
			return
		}
		if got, want := resp.StatusCode, http.StatusOK; got != want {
			t.Errorf("resp.StatusCode = %d; want %d; i=%d", got, want, i)
			t.Logf("%s", buf)
		}

		var msg sub.StringMessage
		if err := json.Unmarshal(buf, &msg); err != nil {
			t.Errorf("json.Unmarshal(%s, &msg) failed with %v; want success; %i", buf, err, i)
			return
		}
		if got, want := msg.GetValue(), "hello"; got != want {
			t.Errorf("msg.GetValue() = %q; want %q", got, want)
		}
	}
}

func testABEProto2Echo(t *testing.T) {
	want := gw.ABitOfEverythingProto2{
		FloatValue:    proto.Float32(1.5),
		DoubleValue:   proto.Float64(2.5),
		Int64Value:    proto.Int64(4294967296),
		Uint64Value:   proto.Uint64(9223372036854775807),
		Int32Value:    proto.Int32(-2147483648),
		Fixed64Value:  proto.Uint64(9223372036854775807),
		Fixed32Value:  proto.Uint32(4294967295),
		BoolValue:     proto.Bool(true),
		StringValue:   proto.String("strprefix/foo"),
		Uint32Value:   proto.Uint32(4294967295),
		BytesValue:    []byte{0, 1, 2, 3, 4, 127, 128, 129, 251, 252, 253, 254, 255},
		Sfixed32Value: proto.Int32(2147483647),
		Sfixed64Value: proto.Int64(-4611686018427387904),
		Sint32Value:   proto.Int32(2147483647),
		Sint64Value:   proto.Int64(4611686018427387903),
	}
	bytes := url.QueryEscape(string(want.BytesValue))
	url := fmt.Sprintf("http://localhost:8080/v1/example/a_bit_of_everything/%f/%f/%d/separator/%d/%d/%d/%d/%v/%s/%d/%s/%d/%d/%d/%d:proto2_echo", *want.FloatValue, *want.DoubleValue, *want.Int64Value, *want.Uint64Value, *want.Int32Value, *want.Fixed64Value, *want.Fixed32Value, *want.BoolValue, *want.StringValue, *want.Uint32Value, bytes, *want.Sfixed32Value, *want.Sfixed64Value, *want.Sint32Value, *want.Sint64Value)

	resp, err := http.Post(url, "application/json", strings.NewReader("{}"))
	if err != nil {
		t.Errorf("http.Post(%q) failed with %v; want success", url, err)
		return
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("iotuil.ReadAll(resp.Body) failed with %v; want success", err)
		return
	}

	if got, want := resp.StatusCode, http.StatusOK; got != want {
		t.Errorf("resp.StatusCode = %d; want %d", got, want)
		t.Logf("%s", buf)
	}

	var got gw.ABitOfEverythingProto2
	if err := json.Unmarshal(buf, &got); err != nil {
		t.Errorf("json.Unmarshal(%s, &got) failed with %v; want success", buf, err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("msg = %v; want %v", got, want)
	}
}

func testABEProto2EchoBody(t *testing.T) {
	want := gw.ABitOfEverythingProto2{
		FloatValue:    proto.Float32(1.5),
		DoubleValue:   proto.Float64(2.5),
		Int64Value:    proto.Int64(4294967296),
		Uint64Value:   proto.Uint64(9223372036854775807),
		Int32Value:    proto.Int32(-2147483648),
		Fixed64Value:  proto.Uint64(9223372036854775807),
		Fixed32Value:  proto.Uint32(4294967295),
		BoolValue:     proto.Bool(true),
		StringValue:   proto.String("strprefix/foo"),
		Uint32Value:   proto.Uint32(4294967295),
		BytesValue:    []byte{0, 1, 2, 3, 4, 127, 128, 129, 251, 252, 253, 254, 255},
		Sfixed32Value: proto.Int32(2147483647),
		Sfixed64Value: proto.Int64(-4611686018427387904),
		Sint32Value:   proto.Int32(2147483647),
		Sint64Value:   proto.Int64(4611686018427387903),

		Nested: []*gw.ABitOfEverythingProto2_Nested{
			{
				Name:   proto.String("bar"),
				Amount: proto.Uint32(10),
			},
			{
				Name:   proto.String("baz"),
				Amount: proto.Uint32(20),
			},
		},
	}
	url := "http://localhost:8080/v1/example/a_bit_of_everything:proto2_echo"
	buf, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("json.Marshal(%#v) failed with %v; want success", want, err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(buf))
	if err != nil {
		t.Errorf("http.Post(%q) failed with %v; want success", url, err)
		return
	}
	defer resp.Body.Close()
	buf, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("iotuil.ReadAll(resp.Body) failed with %v; want success", err)
		return
	}

	if got, want := resp.StatusCode, http.StatusOK; got != want {
		t.Errorf("resp.StatusCode = %d; want %d", got, want)
		t.Logf("%s", buf)
	}

	var got gw.ABitOfEverythingProto2
	if err := json.Unmarshal(buf, &got); err != nil {
		t.Errorf("json.Unmarshal(%s, &got) failed with %v; want success", buf, err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("msg = %v; want %v", got, want)
	}
}

func testABEProto2BulkEcho(t *testing.T) {
	var sent, recv []*gw.ABitOfEverythingProto2
	r, w := io.Pipe()
	done := make(chan struct{})
	go func(w io.WriteCloser) {
		defer close(done)
		defer func() {
			if cerr := w.Close(); cerr != nil {
				t.Errorf("w.Close() failed with %v; want success", cerr)
			}
		}()
		for _, val := range []string{
			"foo", "bar", "baz", "qux", "quux",
		} {
			m := gw.ABitOfEverythingProto2{
				FloatValue:    proto.Float32(1.5),
				DoubleValue:   proto.Float64(2.5),
				Int64Value:    proto.Int64(4294967296),
				Uint64Value:   proto.Uint64(9223372036854775807),
				Int32Value:    proto.Int32(-2147483648),
				Fixed64Value:  proto.Uint64(9223372036854775807),
				Fixed32Value:  proto.Uint32(4294967295),
				BoolValue:     proto.Bool(true),
				StringValue:   proto.String(fmt.Sprintf("strprefix/%s", val)),
				Uint32Value:   proto.Uint32(4294967295),
				BytesValue:    []byte{0, 1, 2, 3, 4, 127, 128, 129, 251, 252, 253, 254, 255},
				Sfixed32Value: proto.Int32(2147483647),
				Sfixed64Value: proto.Int64(-4611686018427387904),
				Sint32Value:   proto.Int32(2147483647),
				Sint64Value:   proto.Int64(4611686018427387903),
			}
			sent = append(sent, &m)
			buf, err := json.Marshal(m)
			if err != nil {
				t.Fatalf("json.Marshal(%#v) failed with %v; want success", m, err)
			}
			if _, err := w.Write(buf); err != nil {
				t.Errorf("w.Write(%s) failed with %v; want success", buf, err)
				return
			}
			if _, err := io.WriteString(w, "\n"); err != nil {
				t.Errorf("w.Write(%s) failed with %v; want success", buf, err)
				return
			}
		}
	}(w)
	url := "http://localhost:8080/v1/example/a_bit_of_everything/bulk:proto2_echo"
	resp, err := http.Post(url, "application/json", r)
	if err != nil {
		t.Errorf("http.Post(%q) failed with %v; want success", url, err)
		return
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("iotuil.ReadAll(resp.Body) failed with %v; want success", err)
		return
	}

	if got, want := resp.StatusCode, http.StatusOK; got != want {
		t.Errorf("resp.StatusCode = %d; want %d", got, want)
		t.Logf("%s", buf)
	}

	dec := json.NewDecoder(bytes.NewReader(buf))
	for {
		var m struct {
			Result *gw.ABitOfEverythingProto2 `json:"result,omitempty"`
		}
		err := dec.Decode(&m)
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Errorf("dec.Decode(&m) failed with %v; want success; buf=%s", err, buf)
		}
		recv = append(recv, m.Result)
	}

	<-done
	if got, want := recv, sent; !reflect.DeepEqual(got, want) {
		t.Errorf("got = %v; want %v", got, want)
	}
}
