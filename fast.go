package main

import (
	"bytes"
	"fmt"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
	"io"
	"os"
	"strings"
)

// suppress unused package warning
var (
	EOL = []byte("\n")
)

//easyjson:json
type User struct {
	Browsers []string `json:"browsers,[]string"`
	Name     string   `json:"name,string"`
	Email    string   `json:"email,string"`
}

// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	user := User{}
	nullSlice := make([]byte, 2048)
	buffer := make([]byte, 1024)
	data := make([]byte, 2048)
	hasAndroid := false
	hasMSIE := false
	num := 0
	index := 0
	position := 0

	seenBrowsers := make(map[string]bool)
	browser := ""

	fmt.Fprintln(out, "found users:")
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return
		}

		copy(data[position:position+n], buffer[:n])

		position = position + n
		for index = bytes.Index(data, EOL); index > -1; index = bytes.Index(data, EOL) {

			hasAndroid = false
			hasMSIE = false

			if err := user.UnmarshalJSON(data[:index]); err != nil {
				break
			}

			copy(data, data[index+1:])
			position = position - index - 1
			copy(data[position:], nullSlice[position:])

			for _, browser = range user.Browsers {
				if strings.Contains(browser, "MSIE") {
					hasMSIE = true
					seenBrowsers[browser] = true
					continue
				}

				if strings.Contains(browser, "Android") {
					hasAndroid = true
					seenBrowsers[browser] = true
					continue
				}
			}

			if hasAndroid && hasMSIE {
				fmt.Fprintln(out, fmt.Sprintf("[%d] %s <%s>", num, user.Name, strings.Replace(user.Email, "@", " [at] ", 1)))
			}

			num++
		}
	}

	fmt.Fprintln(out, "\nTotal unique browsers", len(seenBrowsers))
}

func easyjson3486653aDecodeCourseraHomeworkBench(in *jlexer.Lexer, out *User) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "browsers":
			if in.IsNull() {
				in.Skip()
				out.Browsers = nil
			} else {
				in.Delim('[')
				if out.Browsers == nil {
					if !in.IsDelim(']') {
						out.Browsers = make([]string, 0, 4)
					} else {
						out.Browsers = []string{}
					}
				} else {
					out.Browsers = (out.Browsers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Browsers = append(out.Browsers, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "name":
			out.Name = string(in.String())
		case "email":
			out.Email = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3486653aEncodeCourseraHomeworkBench(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"browsers\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Browsers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Browsers {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Email))
	}
	out.RawByte('}')
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3486653aDecodeCourseraHomeworkBench(&r, v)
	return r.Error()
}
