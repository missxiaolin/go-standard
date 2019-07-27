package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main()  {
	for {
		go func() {
			request,err := http.NewRequest(http.MethodGet, "https://www.vtoall.com/inquirysheet/url/extraterritorial/queryProject?projectId=20190725zpxw&companyId=70037&lng=0&lat=0", nil);
			request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1");
			request.Header.Add("Cookie","gr_user_id=ab621b8b-cedc-45c2-bdf2-9b97acdf67b5; grwng_uid=478d8ac4-df93-4221-ab59-7db67d224577; UM_distinctid=16b4f7d48fd2c-0c41c73ef1b73a-37677e05-13c680-16b4f7d48fe49d; user=ha3yag; random=692607; token=a9f35df9de62dd0738ac2a29132603e5; Hm_lvt_fcc498b3767cd0ad6bb21d78e7b275b6=1563760482,1563801345,1564109723,1564229066; a1ca3f5cc714d6a1_gr_session_id=7c290021-6317-4565-8d89-9b7e383fbdba; CNZZDATA1254450087=1088316045-1560401820-https%253A%252F%252Fwww.baidu.com%252F%7C1564224395; a1ca3f5cc714d6a1_gr_session_id_7c290021-6317-4565-8d89-9b7e383fbdba=true; JSESSIONID=F515AACF14C0C1791E2C9FBCA817682E; a1ca3f5cc714d6a1_gr_last_sent_sid_with_cs1=7c290021-6317-4565-8d89-9b7e383fbdba; a1ca3f5cc714d6a1_gr_last_sent_cs1=70037; a1ca3f5cc714d6a1_gr_cs1=70037; Hm_lpvt_fcc498b3767cd0ad6bb21d78e7b275b6=1564229648");
			if err != nil {
				panic(err)
			}
			resp, err := http.DefaultClient.Do(request)
			if err != nil {
				panic(err)
			}

			defer resp.Body.Close()

			s, err := httputil.DumpResponse(resp, true)
			if err != nil {
				panic(err)
			}

			fmt.Println(s);

			fmt.Printf("%s\n",s)
		}()
	}


}