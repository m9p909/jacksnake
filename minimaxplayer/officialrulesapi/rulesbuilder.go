package officialrulesapi

import (
	"jacksnake/minimaxplayer/coreplayer"
)

func GetOfficialRules() coreplayer.Simulator {
	res := OfficialRulesAdapterImpl{}
	res.init(&OfficialRulesImpl{})
	return &res
}
