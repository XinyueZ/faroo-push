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




type Error string

func (e Error) Error() string {
	return string(e)
}

func init() { 
		http.HandleFunc("/pushEn", pushHandleEnglish) 
		http.HandleFunc("/pushDe", pushHandleGerman) 
		http.HandleFunc("/pushZh", pushHandleChinese) 
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

//A push handler on all languages which will be used by calling http://your-app.appspot.com/push
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
				cxt.Infof("json: %s", string(bytes))
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

 