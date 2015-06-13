package faroo

const(
	//The API-KEY that has been defined in Google-Console.
	//Live
	//API-KEY = ""
	//Dev
	API-KEY = ""
	
	PUSH_SENDER = "https://android.googleapis.com/gcm/send"
  	PUSH_KEY = "key=" + API_KEY
  	API_METHOD = "POST"
  	API_RESTYPE = "application/json"
	TOPICS = "/topics/"
	
	FAROO_API = "http://www.faroo.com/api?q=%s&start=1&length=1&l=%s&src=news&f=json"
)  