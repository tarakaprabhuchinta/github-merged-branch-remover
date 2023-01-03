package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	token, ok := os.LookupEnv("GITHUB_TOKEN")
	if !ok {
		fmt.Printf("Github token not set")
	}
	base_url := "https://api.github.com/repos/"
	// pass github owner and github repo as command line arguments
	if len(os.Args) != 3 {
		panic("Enter Github owner and repo. " +
			"Make sure not to have any spaces in between...")
	}
	github_owner := os.Args[1]
	gibhub_repo := os.Args[2]
	url := base_url + github_owner + "/" + gibhub_repo
	unrefined_branch_list := ListBranches(url, token)
	RefineBranchList(unrefined_branch_list, url, token, github_owner)
}

func SendRequest(url string, req_method string, token string) string {
	client := &http.Client{}
	req, err := http.NewRequest(req_method, url, nil)
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer "+token)
	// Change the api version here
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	resp, err := client.Do(req)
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if req_method == "GET" {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			// handle error
			fmt.Println(err)
		}
		return string(body)
	} else if req_method == "DELETE" {
		return fmt.Sprint(resp.StatusCode)
	}
	return ""
}

func ListBranches(url string, token string) map[string]string {
	// set protected query as false to get unprotected branches
	body := SendRequest(url+"/branches?protected=false", "GET", token)
	return DecodeBranchListInJson(body)

}

func GetDefaultBranch(url string, token string) string {
	body := SendRequest(url, "GET", token)
	df_branch := DecodeDefaultBranch(body)
	return df_branch
}

func CheckIfBranchHasOpenPullRequests(url string, owner string, branch string, token string) bool {
	// This function is limited by the Github API. Github API lets you filter
	// requests by head user or head organization when listing pull requests.
	// This function may fail, if you are trying to list PR's created by
	// another user
	body := SendRequest(url+"/pulls?head="+owner+":"+branch, "GET", token)
	// api returns '[]' string if branch has no open pr's.
	// so string length is checked for 2 characters and it returns 2
	// if the length is not equal to 2
	return len(body) != 2
}

func LastCommitTimeOfBranchInDays(url string, branch string, token string) int {
	body := SendRequest(url+"/commits/heads/"+branch, "GET", token)
	t, err := time.Parse(time.RFC3339, DecodeLastCommitTime(body))
	if err != nil {
		fmt.Println(err)
	}
	// print(branch + "has last commit in " + strconv.Itoa(int(time.Since(t).Hours()/24)) + " days")
	// Calculate the difference between the current date and the specific date in days
	return int(time.Since(t).Hours() / 24)

}
func RefineBranchList(bl map[string]string, url string, token string, owner string) {
	// This function checks the branch list and filter out default branch,
	// branches which has open pull requests, branches which has been merged
	// but has the last commit done in less than 10 days

	// Refined list contains branches as keys, sha commits as values
	rl := make(map[string]string)
	default_branch := GetDefaultBranch(url, token)

	for k, v := range bl {
		// not assigning to default branch if it is a default branch
		if k == default_branch {
			continue
		}
		// Branch is only entered into refined list only if it doesn't have pull
		// request and last commit is more than 10 days back
		if !CheckIfBranchHasOpenPullRequests(url, owner, k, token) && LastCommitTimeOfBranchInDays(url, k, token) > 10 {
			rl[k] = v
		}
	}

	// Only display branches if refined list is not empty
	if len(rl) > 0 {
		fmt.Println("List of branches which are about to be deleted are as follows...")
		for k := range rl {
			fmt.Println(k)
		}
	}
	for k, v := range rl {
		fmt.Println(url + "git/refs/heads/" + k)
		status := SendRequest(url+"/git/refs/heads/"+k, "DELETE", token)
		if status == "204" {
			fmt.Println("Branch deletion with branch name " + k +
				" with sha code of " + v + " is successful")
		} else {
			fmt.Println("Branch deletion of " + k + " is unsuccessful with status " + status)
		}
	}
}
