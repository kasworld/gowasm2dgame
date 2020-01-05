// Code generated by "genenum -typename=ActType -packagename=acttype -basedir=enums -statstype=int"

package acttype

import "fmt"

type ActType uint8

const (
	Nothing       ActType = iota //
	Shield                       //
	SuperShield                  //
	HommingShield                //
	Bullet                       //
	SuperBullet                  //
	HommingBullet                //
	Accel                        //

	ActType_Count int = iota
)

var _ActType2string = [ActType_Count]string{
	Nothing:       "Nothing",
	Shield:        "Shield",
	SuperShield:   "SuperShield",
	HommingShield: "HommingShield",
	Bullet:        "Bullet",
	SuperBullet:   "SuperBullet",
	HommingBullet: "HommingBullet",
	Accel:         "Accel",
}

func (e ActType) String() string {
	if e >= 0 && e < ActType(ActType_Count) {
		return _ActType2string[e]
	}
	return fmt.Sprintf("ActType%d", uint8(e))
}

var _string2ActType = map[string]ActType{
	"Nothing":       Nothing,
	"Shield":        Shield,
	"SuperShield":   SuperShield,
	"HommingShield": HommingShield,
	"Bullet":        Bullet,
	"SuperBullet":   SuperBullet,
	"HommingBullet": HommingBullet,
	"Accel":         Accel,
}

func String2ActType(s string) (ActType, bool) {
	v, b := _string2ActType[s]
	return v, b
}