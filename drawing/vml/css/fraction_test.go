// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package css_test

import (
	"encoding/xml"
	"testing"

	"github.com/roboninc/ooxml/drawing/vml/css"
	"github.com/stretchr/testify/require"
)

func TestFraction(t *testing.T) {
	type Entity struct {
		Opacity css.Fraction `xml:"opacity,attr,omitempty"`
	}

	//empty
	entity := Entity{Opacity: 0}
	encoded, err := xml.Marshal(&entity)
	require.Empty(t, err)
	require.Equal(t, `<Entity></Entity>`, string(encoded))

	entity = Entity{Opacity: -0.5}
	encoded, err = xml.Marshal(&entity)

	require.Empty(t, err)
	require.Equal(t, `<Entity opacity="-0.5"></Entity>`, string(encoded))

	//decode
	var decoded Entity
	err = xml.Unmarshal([]byte(`"<Entity opacity="-0.5"></Entity>"`), &decoded)
	require.Empty(t, err)
	require.Equal(t, entity, decoded)

	decoded = Entity{}
	err = xml.Unmarshal([]byte(`"<Entity opacity="-50%"></Entity>"`), &decoded)
	require.Empty(t, err)
	require.Equal(t, entity, decoded)
}
