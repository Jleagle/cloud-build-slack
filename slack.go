package slack

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/nlopes/slack"
)

func HelloPubSub(ctx context.Context, m PubSubMessage) (err error) {

	build := CloudBuildMessage{}

	err = json.Unmarshal(m.Data, &build)
	if err != nil {
		return err
	}

	return slack.PostWebhook(os.Getenv("ENDPOINT"), &slack.WebhookMessage{
		Username: "Cloud Build",
		IconURL:  "https://avatars0.githubusercontent.com/u/38220399",
		Text: fmt.Sprintf("Project *%s*, Repo: *%s*, Status: *%s*",
			build.SourceProvenance.ResolvedRepoSource.ProjectID,
			build.SourceProvenance.ResolvedRepoSource.RepoName,
			build.Status,
		),
	})
}

type PubSubMessage struct {
	Data []byte `json:"data"`
}

type CloudBuildMessage struct {
	ID        string `json:"id"`
	ProjectID string `json:"projectId"`
	Status    string `json:"status"`
	Source    struct {
		RepoSource struct {
			ProjectID  string `json:"projectId"`
			RepoName   string `json:"repoName"`
			BranchName string `json:"branchName"`
		} `json:"repoSource"`
	} `json:"source"`
	Steps []struct {
		Name   string   `json:"name"`
		Args   []string `json:"args"`
		Timing struct {
			StartTime time.Time `json:"startTime"`
			EndTime   time.Time `json:"endTime"`
		} `json:"timing"`
		PullTiming struct {
			StartTime time.Time `json:"startTime"`
			EndTime   time.Time `json:"endTime"`
		} `json:"pullTiming"`
		Status string `json:"status"`
	} `json:"steps"`
	Results struct {
		Images []struct {
			Name       string `json:"name"`
			Digest     string `json:"digest"`
			PushTiming struct {
				StartTime time.Time `json:"startTime"`
				EndTime   time.Time `json:"endTime"`
			} `json:"pushTiming"`
		} `json:"images"`
		BuildStepImages []string `json:"buildStepImages"`
	} `json:"results"`
	CreateTime time.Time `json:"createTime"`
	StartTime  time.Time `json:"startTime"`
	FinishTime time.Time `json:"finishTime"`
	Timeout    string    `json:"timeout"`
	Images     []string  `json:"images"`
	Artifacts  struct {
		Images []string `json:"images"`
	} `json:"artifacts"`
	LogsBucket       string `json:"logsBucket"`
	SourceProvenance struct {
		ResolvedRepoSource struct {
			ProjectID string `json:"projectId"`
			RepoName  string `json:"repoName"`
			CommitSha string `json:"commitSha"`
		} `json:"resolvedRepoSource"`
	} `json:"sourceProvenance"`
	BuildTriggerID string `json:"buildTriggerId"`
	Options        struct {
		SubstitutionOption string `json:"substitutionOption"`
		Logging            string `json:"logging"`
	} `json:"options"`
	LogURL string   `json:"logUrl"`
	Tags   []string `json:"tags"`
	Timing struct {
		BUILD struct {
			StartTime time.Time `json:"startTime"`
			EndTime   time.Time `json:"endTime"`
		} `json:"BUILD"`
		FETCHSOURCE struct {
			StartTime time.Time `json:"startTime"`
			EndTime   time.Time `json:"endTime"`
		} `json:"FETCHSOURCE"`
		PUSH struct {
			StartTime time.Time `json:"startTime"`
			EndTime   time.Time `json:"endTime"`
		} `json:"PUSH"`
	} `json:"timing"`
}
