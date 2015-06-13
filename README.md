# Faroo-push
#####A push-service for Faroo.com. Based on [GCM's topic-messaging](https://developers.google.com/cloud-messaging/topic-messaging).
#####Each push-entry is the first item in feeds.
#####Each topic has [3 language](http://www.faroo.com/hp/api/api.html#description) that Faroo.com supports.

# About fork
- Check out project.
- Change application-id(dev, live) in app.yaml which is generated in Google-Dev-Console.
- Change API_KEY(dev, live) in config.go which is generated in Google-Dev-Console.
- Change FAROO_KEY(dev, live) in config.go which is generated in by [fraoo.com](http://www.faroo.com/hp/api/api.html#key).
 
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
| /topics/basketball-en  | [basketball](http://www.faroo.com/api?q=basketball&start=1&length=10&l=en&src=news&f=json).  | English |
| /topics/basketball-de  | [Basketball](http://www.faroo.com/api?q=Basketball&start=1&length=10&l=de&src=news&f=json).  | German  |
| /topics/basketball-zh  | [篮球](http://www.faroo.com/api?q=篮球&start=1&length=10&l=zh&src=news&f=json).  |  Chinese |
| /topics/IT-en  | [IT](http://www.faroo.com/api?q=Internet&start=1&length=10&l=en&src=news&f=json).  | English |
| /topics/IT-de  | [IT](http://www.faroo.com/api?q=Internet&start=1&length=10&l=de&src=news&f=json).  | German  |
| /topics/IT-zh  | [IT](http://www.faroo.com/api?q=Internet&start=1&length=10&l=zh&src=news&f=json).  |  Chinese |
| /topics/Android-en  | [Android](http://www.faroo.com/api?q=Android&start=1&length=10&l=en&src=news&f=json).  | English |
| /topics/Android-de  | [Android](http://www.faroo.com/api?q=Android&start=1&length=10&l=de&src=news&f=json).  | German  |
| /topics/Android-zh  | [安卓](http://www.faroo.com/api?q=Android&start=1&length=10&l=zh&src=news&f=json).  |  Chinese |
| /topics/Google-en  | [Google](http://www.faroo.com/api?q=Google&start=1&length=10&l=en&src=news&f=json).  | English |
| /topics/Google-de  | [Google](http://www.faroo.com/api?q=Google&start=1&length=10&l=de&src=news&f=json).  | German  |
| /topics/Google-zh  | [谷歌](http://www.faroo.com/api?q=Google&start=1&length=10&l=zh&src=news&f=json).  |  Chinese |
| /topics/iPhone-en  | [iPhone](http://www.faroo.com/api?q=iPhone&start=1&length=10&l=en&src=news&f=json).  | English |
| /topics/iPhone-de  | [iPhone](http://www.faroo.com/api?q=iPhone&start=1&length=10&l=de&src=news&f=json).  | German  |
| /topics/iPhone-zh  | [iPhone](http://www.faroo.com/api?q=iPhone&start=1&length=10&l=zh&src=news&f=json).  |  Chinese |
| /topics/Apple-en  | [Apple](http://www.faroo.com/api?q=Apple&start=1&length=10&l=en&src=news&f=json).  | English |
| /topics/Apple-de  | [Apple](http://www.faroo.com/api?q=Apple&start=1&length=10&l=de&src=news&f=json).  | German  |
| /topics/Apple-zh  | [苹果](http://www.faroo.com/api?q=Apple&start=1&length=10&l=zh&src=news&f=json).  |  Chinese |
 
 
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
