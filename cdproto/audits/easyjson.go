// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package audits

import (
	json "encoding/json"
	network "bitbucket.org/ShipwrightTibi/chromecrawlingnew/cdproto/network"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC5a4559bDecodeGithubComChromedpCdprotoAudits(in *jlexer.Lexer, out *GetEncodedResponseReturns) {
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
		case "body":
			out.Body = string(in.String())
		case "originalSize":
			out.OriginalSize = int64(in.Int64())
		case "encodedSize":
			out.EncodedSize = int64(in.Int64())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoAudits(out *jwriter.Writer, in GetEncodedResponseReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Body != "" {
		const prefix string = ",\"body\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Body))
	}
	if in.OriginalSize != 0 {
		const prefix string = ",\"originalSize\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.OriginalSize))
	}
	if in.EncodedSize != 0 {
		const prefix string = ",\"encodedSize\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.EncodedSize))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetEncodedResponseReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAudits(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetEncodedResponseReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAudits(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetEncodedResponseReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAudits(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetEncodedResponseReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAudits(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoAudits1(in *jlexer.Lexer, out *GetEncodedResponseParams) {
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
		case "requestId":
			out.RequestID = network.RequestID(in.String())
		case "encoding":
			(out.Encoding).UnmarshalEasyJSON(in)
		case "quality":
			out.Quality = float64(in.Float64())
		case "sizeOnly":
			out.SizeOnly = bool(in.Bool())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoAudits1(out *jwriter.Writer, in GetEncodedResponseParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"requestId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RequestID))
	}
	{
		const prefix string = ",\"encoding\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Encoding).MarshalEasyJSON(out)
	}
	if in.Quality != 0 {
		const prefix string = ",\"quality\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Quality))
	}
	if in.SizeOnly {
		const prefix string = ",\"sizeOnly\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.SizeOnly))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetEncodedResponseParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAudits1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetEncodedResponseParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoAudits1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetEncodedResponseParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAudits1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetEncodedResponseParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoAudits1(l, v)
}
