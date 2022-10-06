# gopenvpn

## Installation

1. Clone the repository:
```
git clone https://github.com/Ribas160/gopenvpn.git
cd gopenvpn
```

2. Build the application:
```
go build -o gopenvpn ./cmd
```

3. Clone application to easy-rsa folder
```
sudo cp gopenvpn /etc/openvpn/easy-rsa
```


## Usage

Create new vpn config
```
sudo ./gopenvpn build {config_name}
```


Assemble config from already created files
```
sudo ./gopenvpn config {config_name}
```
