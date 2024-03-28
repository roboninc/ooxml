// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml_test

import (
	"bytes"
	"encoding/xml"
	"strings"
	"testing"

	"github.com/roboninc/ooxml/drawing/vml"
	"github.com/roboninc/ooxml/drawing/vml/css"
	"github.com/stretchr/testify/require"
)

func TestStroke(t *testing.T) {
	data := strings.NewReplacer("\t", "", "\n", "", "\r", "").Replace(`
<xml xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:x="urn:schemas-microsoft-com:office:excel" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships">
	<v:shape path="m 1,1 l 1,200, 200,200, 200,1 x e" style="position:relative;top:1;left:1;width:400;height:400">
		<v:stroke color="red" linestyle="thickThin" weight="50px"></v:stroke>
	</v:shape>
</xml>
`)

	//decode
	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	entity := &vml.Excel{}
	err := decoder.DecodeElement(entity, nil)
	require.Nil(t, err)

	require.Equal(t, "stroke", entity.Shape[0].Stroke.XMLName.Local)
	require.Equal(t, css.NewNumber(50), entity.Shape[0].Stroke.Weight)
	require.Equal(t, "red", entity.Shape[0].Stroke.Color)
	require.Equal(t, vml.StrokeLineStyleThickThin, entity.Shape[0].Stroke.LineStyle)

	//encode
	encoded, err := xml.Marshal(&entity)
	require.Nil(t, err)

	require.Equal(t, data, string(encoded))
}
