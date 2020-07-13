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
# Anycast address for DNSaaS
baseURL: http://10.0.0.1

# Timeout in seconds
timeout: 5

# Output type format (text, json)
output-type: text

# Dump HTTP query and response to STDERR
debug: false
```
