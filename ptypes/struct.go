// Go support for Protocol Buffers - Google's data interchange format
//
// Copyright 2016 The Go Authors.  All rights reserved.
// https://github.com/golang/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package ptypes

// This file implements conversions between google.protobuf.Struct
// and map[string]interface{}.

import (
	structpb "github.com/golang/protobuf/ptypes/struct"
)

func getNativeValue(v *structpb.Value) (val interface{}) {
	switch v.Kind.(type) {
	case *structpb.Value_NullValue:
		return nil
	case *structpb.Value_NumberValue:
		return v.GetNumberValue()
	case *structpb.Value_StringValue:
		return v.GetStringValue()
	case *structpb.Value_BoolValue:
		return v.GetBoolValue()
	case *structpb.Value_StructValue:
		// recurse
		m, _ := GetMap(v.GetStructValue())
		return m
	case *structpb.Value_ListValue:
		var list []interface{}
		for _, lv := range v.GetListValue().Values {
			list = append(list, getNativeValue(lv))
		}
		return list
	}
	return false
}

func GetMap(s *structpb.Struct) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	for k, v := range s.Fields {
		m[k] = getNativeValue(v)
	}
	return m, nil
}

// func GetStructPb(m map[string]interface{}) (*structpb.Struct, error) {
// }
