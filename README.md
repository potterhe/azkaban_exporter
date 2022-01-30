```sh
go mod init github.com/potterhe/azkaban_exporter
cobra init --pkg-name github.com/potterhe/azkaban_exporter
```

## Usage

```sh
./azkaban_exporter --server http://10.128.0.211:8081 --listen-address :8090
```