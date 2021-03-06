package faroo

import (
	"appengine"
	"appengine/urlfetch"
 
 	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"  
	"strings"
)
 

type Related struct {
	Title	string `json:"title"`
	Url	string `json:"url"`
	Domain	string `json:"domain"`
}

type Entry struct {
	Title	string `json:"title"`
	Kwic	string `json:"kwic"`
	Content string `json:"content"`
	Url	string `json:"url"`
	Iurl	string `json:"iurl"`
	Domain	string `json:"domain"`
	Authoor string `json:"author"` 
	Date	int64 `json:"date"` 
}


type EntryList struct {
	Results []Entry `json:"results"`
	Query	string `json:"query"`
	Suggestions []string `json:"suggestions"`
	Count int `json:"count"`
	Start int `json:"start"`
	Length int `json:"length"`
	Time string `json:"time"`
}


type PushToken struct {
	CreatedAt string `json:"createdAt"`
	CustomizedTopic1 string `json:"mCustomizedTopic1"`
	CustomizedTopic2 string `json:"mCustomizedTopic2"`
	CustomizedTopic3 string `json:"mCustomizedTopic3"`
	CustomizedTopic4 string `json:"mCustomizedTopic4"`
	CustomizedTopic5 string `json:"mCustomizedTopic5"`
	Language string  `json:"mLanguage"`
	DeviceId string `json:"mDeviceId"`
	GoogleId string `json:"mGoogleId"`
	PushToken string `json:"mPushToken"`
	ObjectId string `json:"objectId"`
	UpdatedAt string `json:"updatedAt"`
}

type PushTokenList struct {
	Results []PushToken  `json:"results"`
}

type Error string

func (e Error) Error() string {
	return string(e)
}

func init() { 
		http.HandleFunc("/pushEn", pushHandleEnglish) 
		http.HandleFunc("/pushDe", pushHandleGerman) 
		http.HandleFunc("/pushZh", pushHandleChinese) 
		http.HandleFunc("/pushCus", pushHandleCus) 
		http.HandleFunc("/apiListEn", pushHandleApiListEnglish) 
		http.HandleFunc("/apiListDe", pushHandleApiListGerman) 
		http.HandleFunc("/apiListZh", pushHandleApiListChinese) 
}

func pushHandleEnglish(w http.ResponseWriter, r *http.Request ) {
	pushHandle(w, r, TOPIC_LIST_EN)
}

func pushHandleGerman(w http.ResponseWriter, r *http.Request ) {
	pushHandle(w, r, TOPIC_LIST_DE)
}

func pushHandleChinese(w http.ResponseWriter, r *http.Request ) {
	pushHandle(w, r, TOPIC_LIST_ZH)
}


func pushHandleApiListEnglish(w http.ResponseWriter, r *http.Request ) {
	apiListHandle(w, r, TOPIC_LIST_EN)
}

func pushHandleApiListGerman(w http.ResponseWriter, r *http.Request ) {
	apiListHandle(w, r, TOPIC_LIST_DE)
}

func pushHandleApiListChinese(w http.ResponseWriter, r *http.Request ) {
	apiListHandle(w, r, TOPIC_LIST_ZH)
}

//The gernal api to getting json result of topics.
func apiListHandle(w http.ResponseWriter, r *http.Request , topicList map[string]Topic) {
	 defer func() {
		if err := recover(); err != nil {
			s := fmt.Sprintf(`{"status":%d }`, 300)
			w.Header().Set("Content-Type", API_RESTYPE)
			fmt.Fprintf(w, s)
		}
	}()
	
	bytes, _ := json.Marshal(topicList) 
	output := fmt.Sprintf(`{"status":%d, "api:":[%s] }`, 200, bytes)
	w.Header().Set("Content-Type", API_RESTYPE)
	fmt.Fprintf(w, output)	
}

//Push customized topics.
func pushHandleCus(w http.ResponseWriter, r *http.Request ) {
	defer func() {
		if err := recover(); err != nil {
			s := fmt.Sprintf(`{"status":%d }`, 300)
			w.Header().Set("Content-Type", API_RESTYPE)
			fmt.Fprintf(w, s)
		}
	}()
	
	ch := make(chan string)
	go doCus(r, ch)
	pushedResult := <-ch 
	if pushedResult != "" {
		pushedResult = pushedResult[:len(pushedResult)-1]
	}
	output := fmt.Sprintf(`{"status":%d, "pushed:":[%s] }`, 200, pushedResult)
	w.Header().Set("Content-Type", API_RESTYPE)
	fmt.Fprintf(w, output)
}

//A gernal push handler on all languages which will be used by calling http://your-app.appspot.com/push
func pushHandle(w http.ResponseWriter, r *http.Request, topicList map[string]Topic ) {
	defer func() {
		if err := recover(); err != nil {
			s := fmt.Sprintf(`{"status":%d }`, 300)
			w.Header().Set("Content-Type", API_RESTYPE)
			fmt.Fprintf(w, s)
		}
	}()
	
	topicCount := len(topicList) 
	ch := make(chan string, topicCount)  
    for k, v := range topicList {
        go doPush(r,k,v, ch) 	
    }  
	var pushedResult string = ""
	for i:=0; i < topicCount; i++ {
		pushedResult += (<-ch) 
	}	
	if pushedResult != "" {
		pushedResult = pushedResult[:len(pushedResult)-1]
	}
	output := fmt.Sprintf(`{"status":%d, "pushed:":[%s] }`, 200, pushedResult)
	w.Header().Set("Content-Type", API_RESTYPE)
	fmt.Fprintf(w, output)
}

//Push topics.
func doPush(r *http.Request, api string, topic Topic, ch chan string) {
	var pushedContent string = ""   
	pRes := getEntryList(r, topic.LocalName, topic.Language) 
	topicApi := (TOPICS + api)
	if pRes != nil && pRes.Results != nil &&  len(pRes.Results) > 0 {
		entry := pRes.Results[0]
		
		entry.Title = strings.Replace(entry.Title, "\"", "'", -1)
		entry.Title = strings.Replace(entry.Title, "%", "％", -1)
		entry.Kwic = strings.Replace(entry.Kwic, "\"", "'", -1)
		entry.Kwic = strings.Replace(entry.Kwic, "%", "％", -1)
		
		resBytes, _ := json.Marshal(entry)
		data := string(resBytes) 
		pushed := push(r, topicApi, data, true)
		pushedContent += pushed 
	} else {
		pushedContent += (fmt.Sprintf(`"failed in %s"`, topicApi))
	}  
	pushedContent += ","  
	ch <- pushedContent
}

//Push customized things.
func doCus(r *http.Request, ch chan string) {
	pushTokenList := getPushTokens(r)
	if pushTokenList == nil {
		ch <- `"Error(nil) in getting push-token""`
		return
	}
	var allPushedContent string = "" 
	if pushTokenList.Results != nil && len(pushTokenList.Results) > 0 { 
		var sz int = len(pushTokenList.Results)
		out := make(chan string, sz)
		for _, v := range pushTokenList.Results {
    		go func(pushToken PushToken) {
				pRess := [CUSTOMIZED_TOTAL]*EntryList{}
				if pushToken.CustomizedTopic1 != "" {
					pRess[0] = getEntryList(r, pushToken.CustomizedTopic1, pushToken.Language) 
				} else {
					pRess[0] = nil
				}
				if pushToken.CustomizedTopic2 != "" {
					pRess[1] = getEntryList(r, pushToken.CustomizedTopic2, pushToken.Language) 
				} else {
					pRess[1] = nil
				}
				if pushToken.CustomizedTopic3 != "" {
					pRess[2] = getEntryList(r, pushToken.CustomizedTopic3, pushToken.Language) 
				} else {
					pRess[2] = nil
				}
				if pushToken.CustomizedTopic4 != "" {
					pRess[3] = getEntryList(r, pushToken.CustomizedTopic4, pushToken.Language) 
				} else {
					pRess[3] = nil
				}
				if pushToken.CustomizedTopic5 != "" {
					pRess[4] = getEntryList(r, pushToken.CustomizedTopic5, pushToken.Language) 
				} else {
					pRess[4] = nil
				}
				topicApi := pushToken.PushToken
				var pushedContent string = "" 
				for _, pRes := range pRess {
					if pRes != nil && pRes.Results != nil &&  len(pRes.Results) > 0 {
						entry := pRes.Results[0]
						
						entry.Title = strings.Replace(entry.Title, "\"", "'", -1)
						entry.Title = strings.Replace(entry.Title, "%", "％", -1)
						entry.Kwic = strings.Replace(entry.Kwic, "\"", "'", -1)
						entry.Kwic = strings.Replace(entry.Kwic, "%", "％", -1)
						
						resBytes, _ := json.Marshal(entry)
						data := string(resBytes) 
						pushed := push(r, topicApi, data, true)
						pushedContent += pushed 
					} else {
						pushedContent += (fmt.Sprintf(`"failed in %s"`, topicApi))
					}
					pushedContent += ","  
				}
				out <- pushedContent
    		}(v)	
		}
		for i := 0; i < sz; i++ {
			allPushedContent += <-out
		}
	} 
	ch <- allPushedContent
}

//Get all push-tokens of clients inc. the customized topics user subscribed.
func getPushTokens(r *http.Request)(pRes *PushTokenList) {
	cxt := appengine.NewContext(r)
	url := DB_PATH + DB_PUSH_TOKEN_TAB
	if req, err := http.NewRequest("GET", url, nil); err == nil {
		req.Header.Add(DB_HEADER_APP_ID, DB_APP_ID)
		req.Header.Add(DB_HEADER_API_KEY, DB_API_KEY)
		httpClient := urlfetch.Client(cxt)
		r, err := httpClient.Do(req)
		if r != nil {
			defer r.Body.Close()
		}
		if err == nil {
			if bytes, err := ioutil.ReadAll(r.Body); err == nil { 
				pRes = new(PushTokenList)
				json.Unmarshal(bytes, pRes) 
				cxt.Infof("json: %s", string(bytes))
			} else {
				cxt.Errorf("getPushTokens unmarshal: %v", err)
				pRes = nil
			}
		} else {
			cxt.Errorf("getPushTokens doing: %v", err)
			pRes = nil
		}
	} else {
		cxt.Errorf("getPushTokens: %v", err)
		pRes = nil
	}
	return
}


//Get a news-entry and channel it.
func getEntryList(r *http.Request, query string, lang string)(pRes *EntryList) {
	cxt := appengine.NewContext(r)
	if strings.Contains(query, "global") {
		query = ""
	}
	url := fmt.Sprintf(FAROO_API, query, lang) 
	if req, err := http.NewRequest("GET", url + "&rlength=0", nil); err == nil {
		httpClient := urlfetch.Client(cxt)
		r, err := httpClient.Do(req)
		if r != nil {
			defer r.Body.Close()
		}
		if err == nil {
			if bytes, err := ioutil.ReadAll(r.Body); err == nil { 
				pRes = new(EntryList)
				json.Unmarshal(bytes, pRes)  
			} else {
				cxt.Errorf("getEntryList unmarshal: %v", err)
				pRes = nil
			}
		} else {
			cxt.Errorf("getEntryList doing: %v", err)
			pRes = nil
		}
	} else {
		cxt.Errorf("getEntryList: %v", err)
		pRes = nil
	}
	return
}

//Make push
func push(r *http.Request, topicApi string, data string, scheduledTask bool) (push string) {
	body := fmt.Sprintf(`{"to" : "%s","data" : %s}`, topicApi, data)
	bodyBytes := bytes.NewBufferString(body)		
	cxt := appengine.NewContext(r)
	if req, err := http.NewRequest("POST", PUSH_SENDER, bodyBytes); err == nil {
		req.Header.Add("Authorization", PUSH_KEY)
		req.Header.Add("Content-Type", API_RESTYPE)
		if scheduledTask {
			req.Header.Add("X-AppEngine-Cron", "true")
		}
		client := urlfetch.Client(cxt)
		res, _ := client.Do(req)
		if res != nil {
			defer res.Body.Close()
		}
		if err != nil {
			cxt.Errorf("Push summary doing: %v", err)
			push =  ""
		} else {
			cxt.Infof(body)
			push = body
		}
	} else {
		cxt.Errorf("Push summary: %v", err)
		push = ""
	} 
	return
}

 