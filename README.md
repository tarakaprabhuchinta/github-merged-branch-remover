[![License](https://img.shields.io/github/license/tarakaprabhuchinta/github-merged-branch-remover)](LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/tarakaprabhuchinta/github-merged-branch-remover)](https://goreportcard.com/report/github.com/tarakaprabhuchinta/github-merged-branch-remover)
# Github-Merged-Branch-Remover
Github Action in Go to check and remove merged branches in Gihub repositories in which merged branches has their commits more than 10 days ago.

### Prerequsite

A Github token is needed with read and write permissons to run this action
on a repo.

### To run as action

```
name: Github Merged Branch Remover
        uses: tarakaprabhuchinta/github-merged-branch-remover@v2.0.0
        with:
          organization: <organization-name>
          repository: <repo-name>
          github-token: <github-token>

```
