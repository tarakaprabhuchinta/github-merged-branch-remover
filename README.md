[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=tarakaprabhuchinta_github-merged-branch-remover&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=tarakaprabhuchinta_github-merged-branch-remover)

# github-merged-branch-remover
Go code to check and remove merged branches in Gihub repositories in which merged branches has their commits more than 10 days ago.

### Github Token
A Github token is needed to run this code. The token needs to have read and write permissions to the repo.

```
 export GITHUB_TOKEN=<your-token>

 ```

## Run the code using the below command

```
go run main.go json-decode.go <owner/organization-name> <repo-name>

```
