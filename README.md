# dnscli
The dnscli utility is used to manage DNS zones on domains and domain aliases through CLI.

#### Install
```
go get github.com/mixanemca/dnscli
```

#### Build from source
```
go get -d github.com/mixanemca/dnscli
cd $GOPATH/src/github.com/mixanemca/dnscli
sudo make install
```

#### Config example
```yaml
---
# Address or IP for API server
baseURL: https://example.com

# TLS
# controls whether a client uses TLS transport
tls: false
# path to Certificate Authority certificate chain
cacert: /Users/john/certs/personal-certificate-chained.pem
# path to personal TLS key
key: /Users/john/private/johndoe.key
# path to personal TLS certificate
cert: /Users/john/private/johndoe.crt

# Timeout in seconds
timeout: 5

# Output type format (text, json)
output-type: text

# Dump HTTP query and response to STDERR
debug: false
```
