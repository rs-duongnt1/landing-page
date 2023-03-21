## Setting https with mkcert
Install mkcert (homebrew)

Run command:

```sh
# If it's the firt install of mkcert, run
mkcert -install

# Generate certificate for domain "docker.localhost", "domain.local" and their sub-domains
mkcert -cert-file certs/local-cert.pem -key-file certs/local-key.pem "docker.localhost" "*.docker.localhost" "domain.local" "*.domain.local"

```

## Important
As you can see from the following line of your curl/wget output

```sh
Resolving localhost (localhost)... ::1, 127.0.0.1, 127.0.0.1
```

your system is first resolving localhost to ::1, the IPv6 lookback, which is not supported by Docker at the moment.

Possible solutions are:

1. Use 127.0.0.1: curl http://127.0.0.1:8000/members
2. Force IPv4: curl -4 http://localhost:8000/members
3. Disable IPv6: As root, open `/etc/sysctl.conf` and add the following:
```sh
net.ipv6.conf.all.disable_ipv6 = 1
net.ipv6.conf.default.disable_ipv6 = 1
```