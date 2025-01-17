package jpushclient

const (
	TAG     = "tag"
	TAG_AND = "tag_and"
	ALIAS   = "alias"
	ID      = "registration_id"
)

type Audience struct {
	Object   interface{}
	audience map[string][]string
}

func (audience *Audience) All() {
	audience.Object = "all"
}

func (audience *Audience) SetID(ids []string) {
	audience.set(ID, ids)
}

func (audience *Audience) SetTag(tags []string) {
	audience.set(TAG, tags)
}

func (audience *Audience) SetTagAnd(tags []string) {
	audience.set(TAG_AND, tags)
}

func (audience *Audience) SetAlias(alias []string) {
	audience.set(ALIAS, alias)
}

func (audience *Audience) set(key string, v []string) {
	if audience.audience == nil {
		audience.audience = make(map[string][]string)
		audience.Object = audience.audience
	}

	//v, ok = audience.audience[key]
	//if ok {
	//	return
	//}
	audience.audience[key] = v
}
