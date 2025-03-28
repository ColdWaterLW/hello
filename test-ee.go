package hello

import "fmt"

func test()  {
	fmt.Println("testfffaa")
}

func (d *Oauth2ConfigurationUsecase) backendLogout(ctx context.Context, configuration *Oauth2Configuration, idToken string) error {
	logoutUrl, err := url.Parse(configuration.ServerLogoutUrl)
	if err != nil {
		return fmt.Errorf("parse logout url failed: %v", err)
	}
	query := logoutUrl.Query()
	for key := range query {
		val := query.Get(key)
		if val == userVariableIdToken {
			query.Set(key, idToken)
		} else if val == userVariableSqleUrl {
			query.Del(key)
		}
	}
	logoutUrl.RawQuery = query.Encode()
	logoutUrlStr := logoutUrl.String()
	d.log.Infof("backendLogout url: %s", logoutUrlStr)

	client := &http.Client{Timeout: time.Minute}
	resp, err := client.Get(logoutUrlStr)
	if err != nil {
		return fmt.Errorf("request logout url failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request logout url resp Status: %v", resp.Status)
	}
	return nil
}
