# gh_authkey_checker
This is an sshd helper intended to be used in conjunction with the `AuthorizedKeysCommand` configuration option to check a user's presented key against their GitHub account.  This is definitely *not* the most secure thing to do, but can be used for transient instances like Amazon and other cloud instances for temporary deployments so that you, as the admin, don't have to manage SSH keys :-) Just add users and go. *Note: Local usernames must match the user's GitHub account name*

## Building 
Simply clone the repository and use `go build -o gh_authkey_checker main.go` to build a binary. This repository is also mirrored on [GitHub](https://github.com/jrdemasi/gh_authkey_checker), because, afterall, it's Go. If you trust me, the most up-to-date copy of the tool can also be found [here](https://jrdemasi.com/files/binaries/gh_authkey_checker "gh_authkey_checker binary download")
