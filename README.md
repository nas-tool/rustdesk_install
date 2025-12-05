# RustDesk 一键脚本

面向 RustDesk 自建服务器（`hbbs`/`hbbr`）的 自动化安装工具，支持跨平台，提供交互式部署，无需手动配置任何参数即可自动完成安装。

- 交互式部署：支持公网 IP、域名（DDNS）与内网静态 IP 选择
- 安装检测：支持全新安装，覆盖安装，卸载功能
- 服务管理：存在 systemd 时生成并启用服务；否则以后台守护方式运行
- 架构支持：`linux/amd64`、`linux/arm64`、`linux/armv7`

作者：wanghaoyu.com.cn



## 使用要求
- 以 `root` 权限运行本仓库脚本
- Linux 服务器 支持 Linux AMD64 ARM64 ARMv7
- Windows 服务器 仅在专业版中支持

## 专业版
- 自定义端口：支持自定义 `hbbs/hbbr` 及相关端口（21114/21115/21116/21117/21118/21119，含 TCP/UDP）
- 支持用户级别的安全防护，保证隐私安全
- Web 控制：通过 Web 控制台启动/停止服务、查看与监测服务状态
- 功能覆盖：包含 RustDesk 专业版所有功能，配套管理与监控，支持自定义端口，lucky反向代理，NAT机器部署，内网穿透部署
- Web 控制台：提供可视化操作界面，支持多用户


## 旗舰版

在专业版的功能之上增加以下功能；

- WEB控制台支持多租户，多节点功能，支持用户到期时间限制，自助付费，IP黑白名单，地域限制，第三方平台接入等功能，更多功能请联系作者


## 快速开始
1. 从Release下载对应架构的可执行文件
2. 上传至目标服务器并赋予执行权限：
   ```bash
   chmod +x ./rustdesk_install
   sudo ./rustdesk_install
   ```
3. 按提示完成部署：
   - 选择部署方式（公网 IP / 域名 / 内网 IP）
   - 已安装时选择覆盖安装或卸载

完成后将输出：
- `ID服务器地址`
- `Key`（安装目录中 `.pub` 文件内容）
- `日志目录`（默认 `/var/log/rustdesk`）

## 架构与模块
- 入口：`cmd/rustdesk_install/main.go`
- 安装流程：`internal/installer/installer.go`
- 交互：`internal/prompt/prompt.go`
- 下载与进度：`internal/downloader/downloader.go`
- 解压：`internal/archive/unzip.go`
- 文件移动：`internal/fsops/binmove.go`
- 网络解析：`internal/netutil/netutil.go`
- systemd 管理：`internal/systemd/systemd.go`
- 通用工具：`internal/util/util.go`


开源不等于安全；请确保服务器及网络环境安全防护到位。

---
如需添加无人值守参数（命令行参数）、端口占用检测或自动证书配置，欢迎继续提出，我将按你的使用场景扩展实现。


后续开发计划：
- 支持 Web SSH 一键安装
- 支持 Docker部署
- 支持OpenWrt套件部署

