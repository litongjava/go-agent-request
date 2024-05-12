# go-agent-request

```shell
vi /etc/systemd/system/go-agent-request.service
```

```
[Unit]
Description=go-agent-request
After=network.target

[Service]
Type=simple
User=root
Restart=on-failure
RestartSec=5s
WorkingDirectory = /opt/go-agent-request
ExecStart=/opt/go-agent-request/go-agent-request

[Install]
WantedBy=multi-user.target
```
启动
```
sudo systemctl daemon-reload
sudo systemctl enable go-agent-request
sudo systemctl start go-agent-request
sudo systemctl status go-agent-request
sudo systemctl stop go-agent-request
sudo systemctl restart go-agent-request
```