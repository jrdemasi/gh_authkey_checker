# gh_authkey_checker
This is an sshd helper intended to be used in conjunction with the `AuthorizedKeysCommand` configuration option to check a user's presented key against their GitHub account.  This is definitely *not* the most secure thing to do, but can be used for transient instances like Amazon and other cloud instances for temporary deployments so that you, as the admin, don't have to manage SSH keys :-) Just add users and go. *Note: Local usernames must match the user's GitHub account name*

## Building 
Simply clone the repository and use `go build -o gh_authkey_checker main.go` to build a binary. This repository is also mirrored on [GitHub](https://github.com/jrdemasi/gh_authkey_checker), because, afterall, it's Go. If you trust me, the most up-to-date copy of the tool can also be found [here](https://files.jthan.io/binaries/gh_authkey_checker "gh_authkey_checker binary download")

## Installation
Per FHS, this most likely belongs in `/opt`, but I'm sticking it in `/usr/local/sbin` to minimize configuration (again, on testing and short deployment cloud instances).  Simply copy the binary you built or downloaded in `/usr/local/sbin`, then `chown root:root /usr/local/sbin/gh_authkey_checker`.  Lastly, make sure that only root can execute, `chmod 700 /usr/local/sbin/gh_authkey_checker`.

## Configuration
This utility is only tested on Arch and CentOS 7 at this point, but should work fine on Debian and Ubuntu as well.  Uncomment/add/modify the following lines in `/etc/ssh/sshd_config`:
```    
AuthorizedKeysCommand /usr/local/sbin/gh_authkey_checker
AuthorizedKeysCommandUser root
```
then restart sshd `systemctl restart sshd`.

### Contributions
Special thanks to [amayer](https://github.com/amayer5125) for always critiquing my Go and being a smartass :-)
