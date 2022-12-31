package main

import (
	"encoding/json"
	"fmt"
)

func DecodeBranchListInJson(body string) map[string]string {
	var jsonBlob = []byte(body)

	type Branch struct {
		Name   string
		Commit map[string]string
	}
	branch_data := make(map[string]string, 0)
	var branches []Branch
	err := json.Unmarshal(jsonBlob, &branches)
	if err != nil {
		fmt.Println("error:", err)
	}
	for _, v := range branches {
		branch_data[v.Name] = v.Commit["sha"]
	}
	return branch_data

}

func DecodeDefaultBranch(body string) string {
	type Repo struct {
		DefaultBranch string `json:"default_branch"`
	}
	var repo Repo
	err := json.Unmarshal([]byte(body), &repo)
	if err != nil {
		panic(err)
	}
	return repo.DefaultBranch
}

func DecodeLastCommitTime(body string) string {

	type Data struct {
		Commit struct {
			Committer struct {
				Date string `json:"date"`
			} `json:"committer"`
		} `json:"commit"`
	}

	var data Data
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		fmt.Println(err)
	}

	return data.Commit.Committer.Date
}
