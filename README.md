# Faroo-push
- A push-service for Faroo.com. Based on [GCM's topic-messaging](https://developers.google.com/cloud-messaging/topic-messaging).
- Each push-entry is the first item in feeds.
- Each topic has [3 language](http://www.faroo.com/hp/api/api.html#description) that Faroo.com supports.
- Support push topics automaitically, see cron.yaml
- Support push customized topics from clients automaitically.
- Example Android App

[![https://play.google.com/store/apps/details?id=com.cusnews](https://dl.dropbox.com/s/phrg0387osr3riz/images.jpeg)](https://play.google.com/store/apps/details?id=com.cusnews)


# About fork
- Check out project.
- Change application-id(dev, live) in app.yaml which is generated in Google-Dev-Console.
- Change API_KEY(dev, live) in config.go which is generated in Google-Dev-Console.
- Change FAROO_KEY(dev, live) in config.go which is generated in by [fraoo.com](http://www.faroo.com/hp/api/api.html#key).
- Default the app is backend by [bmob](http://www.bmob.cn/) to save all push-tokens and different customized topics that users generated and defined on client-app with a Google-Id(after G+ Login). A 3rd part backend is of cos available like [parse.com](http://www.parse.com), [Firebase](http://www.firebase.com) etc.

HOW TO:

1. Add DB_HEADER_APP_ID, the header-name of application-id to config.go, i.e X-Bmob-Application-Id for [bmob](http://www.bmob.cn/).
2. Add DB_HEADER_API_KEY, the header-name of application-API-key to config.go, i.e X-Bmob-REST-API-Key for [bmob](http://www.bmob.cn/).
3. Add DB_APP_ID, the application-id to config.go.
4. Add DB_API_KEY, the application-API-key to config.go
5. Add DB_PATH, the absolute path of your database on internet, i.e https://api.bmob.cn/1/classes/ for [bmob](http://www.bmob.cn/).
6. Add DB_PUSH_TOKEN_TAB, the table to store all push-tokens with a Google-Id(after G+ Login). 
7. Add CUSTOMIZED_TOTAL, the total of customized topics you allow for client to push, default is  ```CUSTOMIZED_TOTAL = 5```.

The scheme of push-token(DB_PUSH_TOKEN_TAB):
 ```json
	{
		"createdAt": "2015-06-24 17:10:59",
		"mCustomizedTopic1": "USA",
		"mCustomizedTopic2": "Shanghai",
		"mCustomizedTopic3": "deutsche",
		"mDeviceId": "f069ce9f6asdfasdf65900b3facce96656b6ad9b",
		"mGoogleId": "1079asdfasdfasdfasdf38088430373916013",
		"mPushToken": "bGjgqC4ucT_bBG7-QkGkgT5m_Nc9y2WFaOf1XVgojCzqIaZtFQ-_SI1knT1GwVe91WX8kp8nD1e_AK2khfJ0euQVOeV54Mge",
		"objectId": "fad1573d78",
		"updatedAt": "2015-06-24 17:11:21"
	}
 ```
 
 Example of backen-db config:
 ```go
	//The backend-database key as backend-database.
	DB_PATH="https://api.bmob.cn/1/classes/"
	DB_PUSH_TOKEN_TAB = "PushToken" 
	DB_HEADER_APP_ID = "X-Bmob-Application-Id"
	DB_HEADER_API_KEY = "X-Bmob-REST-API-Key"
	//Live
	//DB_APP_ID="345w45tsfgsdfgsdfg"
	//DB_API_KEY="asdfasdfq345q345"
	//Dev
	DB_APP_ID="2asdfart34t"
	DB_API_KEY="asdfasd345345q345q345"
	
	//How many customized topics are supported.
	CUSTOMIZED_TOTAL = 5
 ```
 
# Supported topics:
######The list will be grown in the furture and the client-apps must keep pace with updated list to subscribe new pushs.
######Topic is case-sensitive.

| Topic  | Content | language |
| ------------- | ------------- | ------------- |
| /topics/global-en  | [Trending News](http://www.faroo.com/api?q=&start=1&length=10&l=en&src=news&f=json).  | English |
| /topics/global-de  | [Neueste Nachricht](http://www.faroo.com/api?q=&start=1&length=10&l=de&src=news&f=json).  | German  |
| /topics/global-zh  | [最新闻](http://www.faroo.com/api?q=&start=1&length=10&l=zh&src=news&f=json).  |  Chinese |
| /topics/football-en  | [Soccer](http://www.faroo.com/api?q=soccer&start=1&length=10&l=en&src=news&f=json).  | English |
| /topics/football-de  | [Fußball](http://www.faroo.com/api?q=fußball&start=1&length=10&l=de&src=news&f=json).  | German  |
| /topics/football-zh  | [足球](http://www.faroo.com/api?q=足球&start=1&length=10&l=zh&src=news&f=json).  |  Chinese |
| /topics/IT-en  | [IT](http://www.faroo.com/api?q=Internet&start=1&length=10&l=en&src=news&f=json).  | English |
| /topics/IT-de  | [IT](http://www.faroo.com/api?q=Internet&start=1&length=10&l=de&src=news&f=json).  | German  |
| /topics/IT-zh  | [IT](http://www.faroo.com/api?q=网络&start=1&length=10&l=zh&src=news&f=json).  |  Chinese |
| /topics/Google-en  | [Google](http://www.faroo.com/api?q=Google&start=1&length=10&l=en&src=news&f=json).  | English |
| /topics/Google-de  | [Google](http://www.faroo.com/api?q=Google&start=1&length=10&l=de&src=news&f=json).  | German  |
| /topics/Google-zh  | [谷歌](http://www.faroo.com/api?q=谷歌&start=1&length=10&l=zh&src=news&f=json).  |  Chinese |
| /topics/Apple-en  | [Apple](http://www.faroo.com/api?q=Apple&start=1&length=10&l=en&src=news&f=json).  | English |
| /topics/Apple-de  | [Apple](http://www.faroo.com/api?q=Apple&start=1&length=10&l=de&src=news&f=json).  | German  |
| /topics/Apple-zh  | [苹果公司](http://www.faroo.com/api?q=苹果公司&start=1&length=10&l=zh&src=news&f=json).  |  Chinese |
 
 
#License

 ```java
 									The MIT License (MIT)

							Copyright (c) 2015 Chris Xinyue Zhao
			
			Permission is hereby granted, free of charge, to any person obtaining a copy
			of this software and associated documentation files (the "Software"), to deal
			in the Software without restriction, including without limitation the rights
			to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
			copies of the Software, and to permit persons to whom the Software is
			furnished to do so, subject to the following conditions:
			
			The above copyright notice and this permission notice shall be included in all
			copies or substantial portions of the Software.
			
			THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
			IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
			FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
			AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
			LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
			OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
			SOFTWARE.
 ```
