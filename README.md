# ssl-tool

A tool to deal with SSL things.


## Usage

### `ssl-tool`

```plain
Usage: ssl-tool [--log-level LOG-LEVEL] <command> [<args>]

Options:
  --log-level LOG-LEVEL, -l LOG-LEVEL [default: warning]
  --help, -h             display this help and exit

Commands:
  cert
  client
```

- `--log-level` can be one of `debug`, `warning`, `error`

### `ssl-tool cert`

```plain
Usage: ssl-tool cert <command> [<args>]

Global options:
  --log-level LOG-LEVEL, -l LOG-LEVEL [default: warning]
  --help, -h             display this help and exit

Commands:
  parse
```

#### `ssl-tool cert parse`

Used to parse a certificate (chain).

```plain
$ ssl-tool cert parse cert.pem
Certificate
 NAME                     VALUE                                          
 Subject                  CN=*.google.com                                
 Issuer                   CN=GTS CA 1C3,O=Google Trust Services LLC,C=US 
 DNS Names                *.google.com                                   
                          *.appengine.google.com                         
                          *.bdn.dev                                      
                          *.origin-test.bdn.dev                          
                          ...                            
 Not Before               2023-03-06 09:16:21                            
 Not After                2023-05-29 10:16:20                            
 Public Key Algorithm     ECDSA                                          
 Public Key Size (bits)   256                                            

Certificate Chain
 Subject        Issuer             
 *.google.com   GTS CA 1C3         
 GTS CA 1C3     GTS Root R1        
 GTS Root R1    GlobalSign Root CA 
```

#### `ssl-tool client get-certificate URL`

Used to parse a certificate (chain).

```
$ ssl-tool client get-certificate https://google.com
Certificate
 NAME                     VALUE                                          
 Subject                  CN=*.google.com                                
 Issuer                   CN=GTS CA 1C3,O=Google Trust Services LLC,C=US 
 DNS Names                *.google.com                                   
                          *.appengine.google.com                         
                          *.bdn.dev                                      
                          *.origin-test.bdn.dev                          
                          ...                            
 Not Before               2023-03-06 09:16:21                            
 Not After                2023-05-29 10:16:20                            
 Public Key Algorithm     ECDSA                                          
 Public Key Size (bits)   256                                            

Certificate Chain
 Subject        Issuer             
 *.google.com   GTS CA 1C3         
 GTS CA 1C3     GTS Root R1        
 GTS Root R1    GlobalSign Root CA 
```