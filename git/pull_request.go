package git

import (
	"github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
	"context"
)

type PullRequest struct {
	Client      *github.Client
	Log         *logrus.Entry
	Repo        *github.Repository
	PullRequest *github.PullRequest
}

type AffectedFile struct {
	Name   string
	Status string
}

func (pr *PullRequest) GetRepoLanguages() ([]string, error) {
	return pr.getLanguages(*pr.Repo.LanguagesURL)
}

func (pr *PullRequest) getLanguages(langUrl string) ([]string, error) {
	req, err := pr.Client.NewRequest("GET", langUrl, nil)
	if err != nil {
		return nil, err
	}

	langsStat := make(map[string]interface{})
	_, err = pr.Client.Do(context.Background(), req, &langsStat)
	if err != nil {
		return nil, err
	}

	languages := make([]string, 0, len(langsStat))
	for lang := range langsStat {
		languages = append(languages, lang)
	}

	return languages, nil
}

func (pr *PullRequest) GetAffectedFiles() ([]AffectedFile, error) {
	files, _, err := pr.Client.PullRequests.
		ListFiles(context.Background(), *pr.Repo.Owner.Login, *pr.Repo.Name, *pr.PullRequest.Number, nil)

	if err != nil {
		return nil, err
	}

	fileNames := make([]AffectedFile, len(files))
	for _, file := range files {
		fileNames = append(fileNames, AffectedFile{*file.Filename, *file.Status})
	}
	return fileNames, nil
}

func (pr *PullRequest) SetStatus(status, reason string) {
	if _, _, err := pr.Client.Repositories.
		CreateStatus(context.Background(), *pr.Repo.Owner.Login, *pr.Repo.Name, *pr.PullRequest.Head.SHA,
		&github.RepoStatus{
			State:       &status,
			Context:     github.String("alien-ike/prow-spike"),
			Description: &reason,
		}); err != nil {
		pr.Log.Info("Error handling event.", err)
	}
}

func (pr *PullRequest) GetAffectedFiles() (string, error) {

}
