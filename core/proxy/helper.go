package proxy

import (
	"bufio"
	"errors"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var proxyRegex = regexp.MustCompile(`(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}):(\d+)`)

func LoadFromFile(protocol ProxyProtocol, path string, manager *ProxyManager) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		match := proxyRegex.FindStringSubmatch(line)
		if len(match) != 3 {
			return err
		}

		ip := match[1]
		port := match[2]

		proxy := &Proxy{
			Ip:       ip,
			Port:     port,
			Protocol: protocol,
		}

		manager.Add(proxy)
	}

	return scanner.Err()
}

func LoadFromURL(protocol ProxyProtocol, url string, manager *ProxyManager) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to fetch proxies from URL")
	}

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		match := proxyRegex.FindStringSubmatch(line)
		if len(match) != 3 {
			return errors.New("invalid proxy format")
		}

		ip := match[1]
		port := match[2]

		proxy := &Proxy{
			Ip:       ip,
			Port:     port,
			Protocol: protocol,
		}

		manager.Add(proxy)
	}

	return scanner.Err()
}
