package unfurl

import (
	"fmt"
	"testing"
)

var testData = map[string]map[string]string{
	"bitly": {
		"http://j.mp/Y4seGv":    "http://www.nytimes.com/2013/03/11/world/asia/karzai-accuses-us-and-taliban-of-colluding-in-afghanistan.html?ref=global-home&_r=0",
		"http://bit.ly/1T1Hul6": "http://www.polygon.com/2015/11/18/9757114/fallout-4-vault-tec-phone-call",
	},
	"t.co": {
		"http://t.co/bxPFQgZ1AV":  "http://www.nytimes.com/2013/03/14/crosswords/bridge/bridge-spring-north-american-championships.html?partner=rss&emc=rss&_r=0",
		"https://t.co/SlINea0uWD": "https://github.com/lyst/MakingLyst/blob/master/mobile/ios/coding-standards/style-guide.md",
	},
	"tinyurl": {
		"http://tinyurl.com/pj94dvk": "https://github.com/blog/2085-a-new-look-for-repositories",
		"http://tinyurl.com/2lekkm":  "https://github.com/",
	},
	"unfurled urls": {
		"https://twitter.com": "https://twitter.com",
		"https://github.com":  "https://github.com",
	},
}

func TestUnfurlSucceeds(t *testing.T) {
	for k, v := range testData {
		fmt.Printf("Subject: %s\n", k)
		for input, expectedOutput := range v {
			fmt.Printf("    %s ----(Redirects)----> %s = ", input, expectedOutput)
			c := NewClient()
			output, err := c.Process(input)
			if err != nil {
				t.Error(err)
				return
			}
			if output != expectedOutput {
				fmt.Println("FAILED")
				t.Errorf("Expecting %s as result, got %s", expectedOutput, output)
			} else {
				fmt.Println("OK")
			}
		}
	}
}

func TestMaxHops(t *testing.T) {
	c := NewClientWithOptions(Options{MaxHops: 0})
	_, err := c.Process("http://j.mp/Y4seGv")
	if err == nil {
		t.Error("Expecting an TooManyRedirects erro")
	}
	if err != ErrTooManyRedirects {
		t.Error("Expecting an TooManyRedirects erro")
	}
}

func TestGoErrors(t *testing.T) {
	c := NewClientWithOptions(Options{MaxHops: 0})
	_, err := c.Process("http://thisdomaindoesnotexist.nope")
	if err == nil {
		t.Error("Expecting an TooManyRedirects erro")
	}
}
