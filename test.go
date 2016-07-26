package main

import ("github.com/ChimeraCoder/anaconda"
 "fmt"
 "net/url"
)




func main() {
tweet := TwitterApi_Client()
fmt.Println(tweet)
}

func TwitterApi_Client() (anaconda.Tweet){
anaconda.SetConsumerKey("2mNTyLU4k8j2VkIBWzd4dnWIl")
anaconda.SetConsumerSecret("NXNFxyLOvmP04bCwy3XSyZz6uFzM1kpOmcl2vR3viCEjenDj0x")
api := anaconda.NewTwitterApi("755434618774355968-akRoeduqPcv1uYZ0uHXJi837KuJ5Pl3", "WnTZdqHTJkU7FRIlDvJSz1IYEgWsEun3eerhwL9pPVd1Q")
myurl := url.Values{}
myurl.Set("text" , "https//twitter.com/HaikuSlip")
tweet, err := api.PostTweet("snapping several\ntree limbs, see how fast that wind\nis blowing in the rain", myurl)
if err != nil {
 	panic(err)
}
return tweet
}


//func (a TwitterApi) PostTweet("hello", "https://twitter.com/HaikuSlip") (tweet Tweet, err error)

