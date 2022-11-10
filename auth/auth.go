package auth

type Client struct {
	Url   string
	Token string
}

func CreateClient(token string, url string) (Client, error) {
	c := Client{
		Url:   url,
		Token: token,
	}

	return c, nil
}
