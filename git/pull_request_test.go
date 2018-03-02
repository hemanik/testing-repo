package git

import (
	"testing"
	"golang.org/x/oauth2"
	"github.com/google/go-github/github"
	"context"
	"github.com/sirupsen/logrus"
	"fmt"
	. "github.com/onsi/gomega"
)

func testShouldLoadTests(t *testing.T) {
	RegisterTestingT(t)

	token := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "71cc06cf7818d5a8d2d96ad135550735c16b0ff7"},
	)
	client := github.NewClient(oauth2.NewClient(context.Background(), token))
	pr := PullRequest{
		Client:      client,
		Log:         logrus.StandardLogger().WithField("ike-plugins", "test"),
		Repo:        nil,
		PullRequest: nil,
	}

	langs, err := pr.getLanguages("https://api.github.com/repos/arquillian/ike-prow-plugins/languages")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(langs)
}
