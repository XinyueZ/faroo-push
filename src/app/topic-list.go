package faroo

 
type Topic struct {
	Language string
	Name string
	LocalName string
}

 
var TOPIC_LIST_EN = make(map[string]Topic)
var TOPIC_LIST_DE = make(map[string]Topic)
var TOPIC_LIST_ZH = make(map[string]Topic)

func init() {
	TOPIC_LIST_EN["global-en"] = Topic{"en", "", ""}
	TOPIC_LIST_DE["global-de"] =Topic {"de", "", ""}
	TOPIC_LIST_ZH["global-zh"] = Topic{"zh", "", ""}
	TOPIC_LIST_EN["football-en"] =Topic {"en", "football", "soccer"}
	TOPIC_LIST_DE["football-de"] = Topic{"de", "football", "fußball"}
	TOPIC_LIST_ZH["football-zh"] = Topic{"zh", "football", "足球"}
	TOPIC_LIST_EN["basketball-en"] = Topic{"en", "basketball", "basketball"}
	TOPIC_LIST_DE["basketball-de"] = Topic{"de", "basketball", "Basketball"}
	TOPIC_LIST_ZH["basketball-zh"] =Topic {"zh", "basketball", "篮球"}
	TOPIC_LIST_EN["IT-en"] = Topic{"en", "IT", "Internet"}
	TOPIC_LIST_DE["IT-de"] =Topic {"de", "IT", "Internet"}
	TOPIC_LIST_ZH["IT-zh"] = Topic{"zh", "IT", "网络"}
	TOPIC_LIST_EN["Android-en"] = Topic{"en", "Android", "Android"}
	TOPIC_LIST_DE["Android-de"] =Topic {"de", "Android", "Android"}
	TOPIC_LIST_ZH["Android-zh"] = Topic{"zh", "Android", "安卓"}
	TOPIC_LIST_EN["Google-en"] = Topic{"en", "Google", "Google"}
	TOPIC_LIST_DE["Google-de"] = Topic{"de", "Google", "Google"}
	TOPIC_LIST_ZH["Google-zh"] = Topic{"zh", "Google", "谷歌"}
	TOPIC_LIST_EN["iPhone-en"] = Topic{"en", "iPhone", "iPhone"}
	TOPIC_LIST_DE["iPhone-de"] = Topic{"de", "iPhone", "iPhone"}
	TOPIC_LIST_ZH["iPhone-zh"] = Topic{"zh", "iPhone", "iPhone"}
	TOPIC_LIST_EN["Apple-en"] = Topic{"en", "Apple", "Apple"}
	TOPIC_LIST_DE["Apple-de"] = Topic{"de", "Apple", "Apple"}
	TOPIC_LIST_ZH["Apple-zh"] = Topic{"zh", "Apple", "苹果公司"}
}



 