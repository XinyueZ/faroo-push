# Faroo-push
#####A push-service for Faroo.com. Based on [GCM's topic-messaging](https://developers.google.com/cloud-messaging/topic-messaging).
#####Each push-entry is the first item in feeds.
#####Each topic has [3 language](http://www.faroo.com/hp/api/api.html#description) that Faroo.com supports.
 
# Supported topics:
######The list will be grown in the furture and the client-apps must keep pace with updated list to subscribe new pushs.

| Topic  | Content | language |
| ------------- | ------------- | ------------- |
| /topics/global-en  | [Trending News](http://www.faroo.com/api?q=&start=1&length=10&l=en&src=news&f=json).  | English |
| /topics/global-de  | [Neueste Nachricht](http://www.faroo.com/api?q=&start=1&length=10&l=de&src=news&f=json).  | German  |
| /topics/global-zh  | [最新闻](http://www.faroo.com/api?q=&start=1&length=10&l=zh&src=news&f=json).  |  Chinese |
| /topics/football-en  | [Soccer](http://www.faroo.com/api?q=soccer&start=1&length=10&l=en&src=news&f=json).  | English |
| /topics/football-de  | [Fußball](http://www.faroo.com/api?q=fußball&start=1&length=10&l=de&src=news&f=json).  | German  |
| /topics/football-zh  | [足球](http://www.faroo.com/api?q=足球&start=1&length=10&l=zh&src=news&f=json).  |  Chinese |
| /topics/basketball-en  | [basketball](http://www.faroo.com/api?q=soccer&start=1&length=10&l=en&src=news&f=json).  | English |
| /topics/basketball-de  | [Basketball](http://www.faroo.com/api?q=fußball&start=1&length=10&l=de&src=news&f=json).  | German  |
| /topics/basketball-zh  | [篮球](http://www.faroo.com/api?q=篮球&start=1&length=10&l=zh&src=news&f=json).  |  Chinese |
| /topics/it-en  | [IT](http://www.faroo.com/api?q=internet&start=1&length=10&l=en&src=news&f=json).  | English |
| /topics/it-de  | [IT](http://www.faroo.com/api?q=internet&start=1&length=10&l=de&src=news&f=json).  | German  |
| /topics/it-zh  | [IT](http://www.faroo.com/api?q=internet&start=1&length=10&l=zh&src=news&f=json).  |  Chinese |
| /topics/android-en  | [Android](http://www.faroo.com/api?q=android&start=1&length=10&l=en&src=news&f=json).  | English |
| /topics/android-de  | [Android](http://www.faroo.com/api?q=android&start=1&length=10&l=de&src=news&f=json).  | German  |
| /topics/android-zh  | [安卓](http://www.faroo.com/api?q=android&start=1&length=10&l=zh&src=news&f=json).  |  Chinese |
| /topics/google-en  | [Google](http://www.faroo.com/api?q=google&start=1&length=10&l=en&src=news&f=json).  | English |
| /topics/google-de  | [Google](http://www.faroo.com/api?q=google&start=1&length=10&l=de&src=news&f=json).  | German  |
| /topics/google-zh  | [谷歌](http://www.faroo.com/api?q=google&start=1&length=10&l=zh&src=news&f=json).  |  Chinese |
| /topics/iphone-en  | [iPhone](http://www.faroo.com/api?q=iphone&start=1&length=10&l=en&src=news&f=json).  | English |
| /topics/iphone-de  | [iPhone](http://www.faroo.com/api?q=iphone&start=1&length=10&l=de&src=news&f=json).  | German  |
| /topics/iphone-zh  | [iPhone](http://www.faroo.com/api?q=iphone&start=1&length=10&l=zh&src=news&f=json).  |  Chinese |
| /topics/apple-en  | [Apple](http://www.faroo.com/api?q=apple&start=1&length=10&l=en&src=news&f=json).  | English |
| /topics/apple-de  | [Apple](http://www.faroo.com/api?q=apple&start=1&length=10&l=de&src=news&f=json).  | German  |
| /topics/apple-zh  | [苹果](http://www.faroo.com/api?q=apple&start=1&length=10&l=zh&src=news&f=json).  |  Chinese |
 
