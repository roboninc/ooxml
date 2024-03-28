// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml_test

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/roboninc/ooxml/drawing/vml"
	"github.com/stretchr/testify/require"
)

func TestShadowType(t *testing.T) {
	type Entity struct {
		Attribute vml.ShadowType `xml:"attribute,attr,omitempty"`
	}

	list := map[vml.ShadowType]string{
		vml.ShadowType(0):         "",
		vml.ShadowTypeDouble:      vml.ShadowTypeDouble.String(),
		vml.ShadowTypeEmboss:      vml.ShadowTypeEmboss.String(),
		vml.ShadowTypePerspective: vml.ShadowTypePerspective.String(),
		vml.ShadowTypeSingle:      vml.ShadowTypeSingle.String(),
	}

	for k, v := range list {
		t.Run(v, func(tt *testing.T) {
			entity := Entity{Attribute: k}
			encoded, err := xml.Marshal(&entity)

			require.Empty(tt, err)
			if k == 0 {
				require.Equal(tt, `<Entity></Entity>`, string(encoded))
			} else {
				require.Equal(tt, fmt.Sprintf(`<Entity attribute="%s"></Entity>`, v), string(encoded))
			}

			var decoded Entity
			err = xml.Unmarshal(encoded, &decoded)
			require.Empty(tt, err)

			require.Equal(tt, entity, decoded)
			require.Equal(tt, v, decoded.Attribute.String())
		})
	}
}
