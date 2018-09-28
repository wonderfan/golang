org1.example.com
├── ca                    # 根CA的私钥和证书
│   ├── 0e729224e8b3f31784c8a93c5b8ef6f4c1c91d9e6e577c45c33163609fe40011_sk                                     
│   └── ca.org1.example.com-cert.pem           
├── msp                   # MSP配置目录
│   ├── admincerts                              # MSP管理员证书
│   │   └── Admin@org1.example.com-cert.pem
│   ├── cacerts                                 # MSP根CA证书（与ca目录下的证书一致）
│   │   └── ca.org1.example.com-cert.pem
│   └── tlscacerts                              # TLS根CA证书
│       └── tlsca.org1.example.com-cert.pem
├── peers                  # peer节点的MSP配置私钥
│   ├── peer0.org1.example.com    # peer0节点配置        
│   │   ├── msp
│   │   │   ├── admincerts                      # 管理员证书
│   │   │   │   └── Admin@org1.example.com-cert.pem
│   │   │   ├── cacerts                         # 根CA证书
│   │   │   │   └── ca.org1.example.com-cert.pem
│   │   │   ├── keystore                        # 节点的私钥
│   │   │   │   └── 27db82c96b1482480baa1c75f80e5cce249beaab27b70c741bb0e2554355957e_sk
│   │   │   ├── signcerts                       # 节点的证书
│   │   │   │   └── peer0.org1.example.com-cert.pem
│   │   │   └── tlscacerts                      # TLS根CA证书
│   │   │       └── tlsca.org1.example.com-cert.pem
│   │   └── tls
│   │       ├── ca.crt                           # TLS根CA证书
│   │       ├── server.crt                       # 节点用的TLS证书
│   │       └── server.key                       # 节点的TLS私钥
│   └── peer1.org1.example.com
│       ├── msp
│       │   ├── admincerts
│       │   │   └── Admin@org1.example.com-cert.pem
│       │   ├── cacerts
│       │   │   └── ca.org1.example.com-cert.pem
│       │   ├── keystore
│       │   │   └── fdee12a3510fde3155c37128cfec26090ae249bfbca28f884e60c21338493edd_sk
│       │   ├── signcerts
│       │   │   └── peer1.org1.example.com-cert.pem
│       │   └── tlscacerts
│       │       └── tlsca.org1.example.com-cert.pem
│       └── tls
│           ├── ca.crt
│           ├── server.crt
│           └── server.key
├── tlsca       # TLS根CA证书和私钥
│   ├── 945092d936f5838c5a6f6484db974d857933706737d00d04bf65f74e3976f9f8_sk
│   └── tlsca.org1.example.com-cert.pem
└── users       # 默认生成的用户配置，一般会包含1个管理员和1个普通成员
    ├── Admin@org1.example.com                   # MSP管理员
    │   ├── msp                                 
    │   │   ├── admincerts                       # 管理员证书
    │   │   │   └── Admin@org1.example.com-cert.pem 
    │   │   ├── cacerts                          # CA根证书
    │   │   │   └── ca.org1.example.com-cert.pem
    │   │   ├── keystore                         # 管理员私钥
    │   │   │   └── 5890f0061619c06fb29dea8cb304edecc020fe63f41a6db109f1e227cc1cb2a8_sk
    │   │   ├── signcerts                        # 管理员证书
    │   │   │   └── Admin@org1.example.com-cert.pem
    │   │   └── tlscacerts                       # TLS根证书
    │   │       └── tlsca.org1.example.com-cert.pem
    │   └── tls
    │       ├── ca.crt                           # TLS根证书
    │       ├── server.crt                       # 管理员的TLS证书
    │       └── server.key                       # 管理员的TLS私钥
    └── User1@org1.example.com
        ├── msp
        │   ├── admincerts
        │   │   └── User1@org1.example.com-cert.pem
        │   ├── cacerts
        │   │   └── ca.org1.example.com-cert.pem
        │   ├── keystore
        │   │   └── 73cdc0072c7203f1ec512232c780fc84acc9752ef30ebc16be1f4666c02b614b_sk
        │   ├── signcerts
        │   │   └── User1@org1.example.com-cert.pem
        │   └── tlscacerts
        │       └── tlsca.org1.example.com-cert.pem
        └── tls
            ├── ca.crt
            ├── server.crt
            └── server.key

---------------------
