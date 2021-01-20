package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetComments -
func GetComments(session string, commenttableID int) (*ResponseCommentSerializer, error) {
	url := fmt.Sprintf("https://www.yuque.com/api/comments?commentable_id=%d&commentable_type=Doc&include_reactions=true&include_to_user=true", commenttableID)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("cookie", fmt.Sprintf("_yuque_session=%s", session))

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	comments := ResponseCommentSerializer{}
	err = json.Unmarshal(body, &comments)
	if err != nil {
		return nil, err
	}
	return &comments, nil
}

// curl 'https://www.yuque.com/api/comments' \
//   -H 'authority: www.yuque.com' \
//   -H 'sec-ch-ua: "Google Chrome";v="87", " Not;A Brand";v="99", "Chromium";v="87"' \
//   -H 'accept: application/json' \
//   -H 'x-csrf-token: ozxrN2Et79upjX24oQ3MTUkg' \
//   -H 'x-requested-with: XMLHttpRequest' \
//   -H 'sec-ch-ua-mobile: ?0' \
//   -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36' \
//   -H 'content-type: application/json' \
//   -H 'origin: https://www.yuque.com' \
//   -H 'sec-fetch-site: same-origin' \
//   -H 'sec-fetch-mode: cors' \
//   -H 'sec-fetch-dest: empty' \
//   -H 'referer: https://www.yuque.com/dovics/test/wgaghg' \
//   -H 'accept-language: zh,zh-CN;q=0.9' \
//   -H 'cookie: lang=zh-cn; _yuque_session=-hE0TwC7704DfI_AVesjAs4vcDpaZ_-NzCCPNavr07Qx6WWpyMO_TA4PRUNEcGP3_nuQ9gF7d_EV9VZVmrakWg==; UM_distinctid=1763d94fe405db-0bf2fb3d4fc58-930346c-144000-1763d94fe41693; yuque_ctoken=ozxrN2Et79upjX24oQ3MTUkg; _TRACERT_COOKIE__SESSION=dbac49b7-fd15-492f-804d-13c56fba6d9f; CNZZDATA1272061571=1419980719-1607348096-%7C1611112243; tree=a385%01ba380c24-0dcf-4126-b365-8f8a9675cd8e%017' \
//   --data-binary $'{"commentable_type":"Doc","commentable_id":30371501,"body_asl":"<\u0021doctype lake><meta name=\\"doc-version\\" content=\\"1\\" /><meta name=\\"viewport\\" content=\\"fixed\\" /><meta name=\\"typography\\" content=\\"traditional\\" /><p data-lake-id=\\"a20d224c568e48b9d67847a2c66a8c01_p_0\\">1</p><\u0021bd15a6bae5411959542913d61de0a2e6d1563e4abb7b11199f92a57e66523fc8>","format":"lake"}' \
//   --compressed

// AddComment -
func AddComment() {

}

// curl 'https://www.yuque.com/api/comments/9796572' \
//   -X 'DELETE' \
//   -H 'authority: www.yuque.com' \
//   -H 'sec-ch-ua: "Google Chrome";v="87", " Not;A Brand";v="99", "Chromium";v="87"' \
//   -H 'accept: application/json' \
//   -H 'x-csrf-token: ozxrN2Et79upjX24oQ3MTUkg' \
//   -H 'x-requested-with: XMLHttpRequest' \
//   -H 'sec-ch-ua-mobile: ?0' \
//   -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36' \
//   -H 'content-type: application/json' \
//   -H 'origin: https://www.yuque.com' \
//   -H 'sec-fetch-site: same-origin' \
//   -H 'sec-fetch-mode: cors' \
//   -H 'sec-fetch-dest: empty' \
//   -H 'referer: https://www.yuque.com/dovics/test/wgaghg' \
//   -H 'accept-language: zh,zh-CN;q=0.9' \
//   -H 'cookie: lang=zh-cn; _yuque_session=-hE0TwC7704DfI_AVesjAs4vcDpaZ_-NzCCPNavr07Qx6WWpyMO_TA4PRUNEcGP3_nuQ9gF7d_EV9VZVmrakWg==; UM_distinctid=1763d94fe405db-0bf2fb3d4fc58-930346c-144000-1763d94fe41693; yuque_ctoken=ozxrN2Et79upjX24oQ3MTUkg; _TRACERT_COOKIE__SESSION=dbac49b7-fd15-492f-804d-13c56fba6d9f; CNZZDATA1272061571=1419980719-1607348096-%7C1611112243; tree=a385%01ba380c24-0dcf-4126-b365-8f8a9675cd8e%018' \
//   --compressed

// DeleteComment -
func DeleteComment() {

}
