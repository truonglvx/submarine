package sight

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/ungerik/go3d/float64/vec2"
	"lib/navmesh"
	"strings"
	"testing"
)

func TestSight(t *testing.T) {
	mesh, _ := navmesh.LoadMeshFromJSONFile("../fixtures/mesh.json")
	navMesh := navmesh.New(mesh)
	lightMap := GenerateLightMap(navMesh, 1, 3)

	Convey("Sight", t, func() {
		s := New(lightMap)

		Convey("#Clear", func() {
			Convey("should clear cells and lights", func() {
				s.PutLight(&vec2.T{0, 1})
				s.Clear()
				isCellLit := false
				for _, cellsY := range s.cells {
					for _, isLit := range cellsY {
						isCellLit = isCellLit || isLit
					}
				}
				So(isCellLit, ShouldBeFalse)
				So(s.putLights, ShouldBeEmpty)
			})
		})

		Convey("#PutLight", func() {
			Convey("should light points of the specified light", func() {
				s.PutLight(&vec2.T{0, 1})
				So(s.DebugString(), ShouldEqual, strings.Join([]string{
					"_____________________",
					"_____________________",
					"_____________________",
					"__________X__________",
					"________XXXXX________",
					"________XXXXX________",
					"_______XXXXXXX_______",
					"__________XXX________",
					"__________XXX________",
					"__________X__________",
					"_____________________",
					"_____________________",
					"_____________________",
					"_____________________",
					"_____________________",
					"_____________________",
					"_____________________",
					"_____________________",
					"_____________________",
					"_____________________",
					"_____________________",
					"_____________________",
				}, "\n"))

				for y, innerCells := range s.cells {
					for x := range innerCells {
						p := s.lightMap.Helper.navMeshPointByCellPoint(&cellPoint{x, y})
						s.PutLight(p)
					}
				}
				So(s.DebugString(), ShouldEqual, strings.Join([]string{
					"XXXXXXXXXXXXXXXXXXX__",
					"XXXXXXXXXXXXXXXXXXX__",
					"XXXXXXXXXXXXXXXXXXX__",
					"XXXXXXXXXXXXXXXXXXX__",
					"XXXXXXXXXXXXXXXXXXX__",
					"XXXXXXXXXXXXXXXXXXX__",
					"XXXXXXXXXXXXXXXXXXX__",
					"XXXXXXXXXXXXXXXXXXX__",
					"__________XXXXXXXXXX_",
					"__________XXXXXXXXXX_",
					"__________XXXXXXXXXX_",
					"__________XXXXXXXXXX_",
					"__________XXXXXXXXXX_",
					"__________XXXXXXXXXX_",
					"__________XXXXXXXXXX_",
					"__________XXXXXXXXXX_",
					"__________XXXXXXXXXXX",
					"_________XXXXXXXXXXXX",
					"________XXXXXXXXXXXXX",
					"_______XXXXXXXXXXXXXX",
					"______XXXXXXXXXXXXXXX",
					"_____XXXXXXXXXXXXXXXX",
				}, "\n"))
			})
		})

		Convey("#IsLitPoint", func() {
			s.PutLight(&vec2.T{0, 1})

			Convey("with a lit point", func() {
				Convey("should return true", func() {
					So(s.IsLitPoint(&vec2.T{0, 1}), ShouldBeTrue)
				})
			})

			Convey("with a non-lit point", func() {
				Convey("should return false", func() {
					So(s.IsLitPoint(&vec2.T{5, 5}), ShouldBeFalse)
				})
			})

			Convey("with a invalid point", func() {
				Convey("should return false", func() {
					So(s.IsLitPoint(&vec2.T{9, 8}), ShouldBeFalse)
				})
			})
		})
	})
}
