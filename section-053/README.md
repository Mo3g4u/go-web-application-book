# プロジェクトの初期化

## 作業ログ

- gh install
  - [github-cli](https://cli.github.com/)
```sh
 $ curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | sudo dd of=/usr/share/keyrings/githubcli-archive-keyring.gpg
 $ sudo chmod go+r /usr/share/keyrings/githubcli-archive-keyring.gpg
 $ echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | sudo tee /etc/apt/sources.list.d/github-cli.list > /dev/null
 $ sudo apt update
 $ sudo apt install gh
```

- gh auth login

```sh
$ gh auth login
? What account do you want to log into? GitHub.com
? What is your preferred protocol for Git operations? SSH
? Upload your SSH public key to your GitHub account? /home/****/.ssh/id_rsa.pub
? Title for your SSH key: GitHub CLI
? How would you like to authenticate GitHub CLI? Paste an authentication token
Tip: you can generate a Personal Access Token here https://github.com/settings/tokens
The minimum required scopes are 'repo', 'read:org', 'admin:public_key'.
? Paste your authentication token: ****************************************
- gh config set -h github.com git_protocol ssh
✓ Configured git protocol
HTTP 422: Validation Failed (https://api.github.com/user/keys)
key is already in use
```

- リポジトリ作成
```sh
$ gh repo create
? What would you like to do? Create a new repository on GitHub from scratch
? Repository name go_todo_app
? Description TODO Web Application with AUTH by Go.
? Visibility Public
? Would you like to add a .gitignore? Yes
? Choose a .gitignore template Go
? Would you like to add a license? Yes
? Choose a license MIT License
? This will create "go_todo_app" as a public repository on GitHub. Continue? Yes
✓ Created repository Mo3g4u/go_todo_app on GitHub
? Clone the new repository locally? Yes
Cloning into 'go_todo_app'...
remote: Enumerating objects: 4, done.
remote: Counting objects: 100% (4/4), done.
remote: Compressing objects: 100% (4/4), done.
remote: Total 4 (delta 0), reused 0 (delta 0), pack-reused 0
Receiving objects: 100% (4/4), done.
```

- ここからは別リポジトリで作業
  - [go_todo_app](https://github.com/Mo3g4u/go_todo_app)