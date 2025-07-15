package setConfigGameservices

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"

	packageglobalsutils "github.com/autoika/package-globals/src/utils"
)

type SetGameConfigOptionQueryParams struct {
	Action         string
	Function       string
	Category       string
	Value          string
	BackgroundView string
	TemplateView   string
	ActionRequest  string
	Ajax           string
}

func SetGameConfigOption(
	proxy *packageglobalsutils.ParsedProxy,
	worldBaseUrl string,
	cookiesStr string,
	userAgentHeader string,
	secChUaHeader string,
	query SetGameConfigOptionQueryParams,
) (string, error) {
	u, err := url.Parse(worldBaseUrl + "/index.php")
	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Set("action", query.Action)
	params.Set("function", query.Function)
	params.Set("category", query.Category)
	params.Set("value", query.Value)
	params.Set("backgroundView", query.BackgroundView)
	params.Set("templateView", query.TemplateView)
	params.Set("actionRequest", query.ActionRequest)
	params.Set("ajax", query.Ajax)
	u.RawQuery = params.Encode()

	transport := &http.Transport{
		DialContext: (&net.Dialer{Timeout: 5 * time.Second}).DialContext,
	}
	if proxy != nil && proxy.Host != "" && proxy.Port != "" {
		proxyURL := &url.URL{
			Scheme: "http",
			Host:   net.JoinHostPort(proxy.Host, proxy.Port),
		}
		if proxy.User != "" {
			proxyURL.User = url.UserPassword(proxy.User, proxy.Password)
		}
		transport.Proxy = http.ProxyURL(proxyURL)
	}

	client := &http.Client{Transport: transport, Timeout: 10 * time.Second}

	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "es-419,es;q=0.9,en;q=0.8")
	req.Header.Set("User-Agent", userAgentHeader)
	req.Header.Set("Sec-CH-UA", secChUaHeader)
	req.Header.Set("Sec-CH-UA-Mobile", "?0")
	req.Header.Set("Sec-CH-UA-Platform", "\"macOS\"")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Referer", worldBaseUrl+"/index.php?view=options")
	req.Header.Set("Referrer-Policy", "strict-origin-when-cross-origin")
	req.Header.Set("Cookie", cookiesStr)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status %d: %s", resp.StatusCode, string(body))
	}

	return string(body), nil
}
