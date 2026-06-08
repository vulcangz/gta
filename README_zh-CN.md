# 基于 Gemma4 的翻译助手（开发中）

[English Version](README.md)

GTA 是一个免费开源的桌面翻译助手工具。使用 [Kronk](https://github.com/ardanlabs/kronk) 加载并运行 Gemma 4 模型服务器，以服务器-客户端模式运行，支持局域网用户访问使用。
用 Go 语言编写，用户界面采用 [Fyne](https://github.com/fyne-io/fyne)。支持跨平台使用，并享受完整的硬件加速支持（通过 Kronk）。完全在您的本地设备上运行（LLM 模型下载完成后，无需再连接互联网）。

## 为什么？

- 专业性。很久以前，谷歌翻译就是最常使用的翻译工具。Gemma 4 发布之后，第一时间测试过它的翻译功能，果然还是一如既往的专业！与往时不同的是，现在还能看到思考过程。
- 本地部署灵活使用。Gemma 4 对硬件配置是有一定要求的。那对于 SOHO 而言，可以使用一台配置高的电脑、安装 26B（MoE，激活4B）或 31B（稠密）等更强的模型作为服务器，其他用户以客户端模式访问使用。
当然你也可以在自己的电脑上自行部署自己使用。
- 注重隐私。不使用 Ollama，本地部署大模型，不使用浏览器。减少中间环节，提高处理效率，保护隐私和信息安全。
- 可扩展性。后续可实现提取图片文本、撰写文章或总结等 Gemma 4 支持的其他功能。


## 入门指南

### 硬件要求

可参考: [unsloth 网站上的 Gemma 4 硬件要求](https://unsloth.ai/docs/models/gemma-4#hardware-requirements)

表：Gemma 4 推理 GGUF 推荐硬件要求 （单位 = 总内存：RAM + VRAM，或统一内存）。你可以在 MacOS、NVIDIA RTX GPU 等设备上使用 Gemma 4。

| Gemma 4 版本  |          4 位       |         8 位      |      BF16 / FP16   |            
|---------------|--------------------|--------------------|--------------------|
| E2B           |               4 GB |             5–8 GB |             10 GB  |
| E4B           |           5.5–6 GB |            9–12 GB |             16 GB  |
| 26B A4B       |           16–18 GB |           28–30 GB |             52 GB  |
| 31B           |           17–20 GB |           34–38 GB |             62 GB  |

- 缺省配置的模型是 gemma-4-E2B-it-Q4_K_M.gguf。
- 8GB 显存 (GPU) 或 16GB 内存 (CPU)


## 从源代码构建

### 先决条件

使用 Fyne 工具包构建跨平台应用程序。

#### Windows
请按照 Fyne 文档中的《入门指南》[此处](https://docs.fyne.io/started/) 设置 MSYS2，并在 MingW-w64 窗口中进行编译。

#### MacOS
打开终端窗口，输入以下内容以设置 Xcode 命令行工具：

`xcode-select --install`

#### Linux
请在 Fyne 文档中查找您所用发行版的依赖项列表：[此处](https://docs.fyne.io/started/)

### 构建

To build the project run the following command:
```bash
git clone https://github.com/vulcangz/gta.git
cd gta
go build
```

## 用法

以 Windows 系统为例。

客户端服务器在同一台电脑上运行时：

1. 在一终端窗口命令行运行 GTA 服务：gta 或者 gta serve 
2. 在另一终端窗口命令行运行桌面客户端：gta gui
3. 即可在打开的窗口中操作进行翻译

客户端服务器不在同一台电脑上运行时：

1. 在服务器终端窗口命令行运行 GTA 服务：gta 或者 gta serve 
2. 在客户端电脑上，先设置 gRPC 服务相关的环境变量，再在一终端窗口命令行运行桌面客户端：gta gui
```bash
set GTA_GRPC_HOST=你的服务器 IP 地址
set GTA_GRPC_PORT=你的服务器服务端口
gta gui
```
3. 即可在打开的窗口中操作进行翻译

完整的命令行参数选项和环境变量如下：
```bash
gta -h
Usage: gta [options...] [arguments...]

OPTIONS
      --app-name                 <string>              (default: GTA)
      --grpc-host                <string>              (default: 0.0.0.0)
      --grpc-port                <int>                 (default: 9000)
  -h, --help                                                                                                                                              display this help message
      --log-level                <string>
      --model-max-tokens         <int>                 (default: 2048)
      --model-temperature        <float>               (default: 0.7)
      --model-top-k              <int>                 (default: 40)
      --model-top-p              <float>               (default: 0.9)
      --model-url                <string>              (default: https://hf-mirror.com/unsloth/gemma-4-E2B-it-GGUF/blob/main/gemma-4-E2B-it-Q4_K_M.gguf)
      --translation-input-delay  <string>              (default: 300)
      --translation-languages    <string>,[string...]  (default: English;中文;Français;Italiano;日本語;한국어;Deutsch;繁體中文)
      --translation-source       <string>              (default: English)
      --translation-target       <string>              (default: 中文)
  -v, --version                                                                                                                                           display version

ENVIRONMENT
  GTA_APP_NAME                 <string>              (default: GTA)
  GTA_GRPC_HOST                <string>              (default: 0.0.0.0)
  GTA_GRPC_PORT                <int>                 (default: 9000)
  GTA_LOG_LEVEL                <string>
  GTA_MODEL_MAX_TOKENS         <int>                 (default: 2048)
  GTA_MODEL_TEMPERATURE        <float>               (default: 0.7)
  GTA_MODEL_TOP_K              <int>                 (default: 40)
  GTA_MODEL_TOP_P              <float>               (default: 0.9)
  GTA_MODEL_URL                <string>              (default: https://hf-mirror.com/unsloth/gemma-4-E2B-it-GGUF/blob/main/gemma-4-E2B-it-Q4_K_M.gguf)
  GTA_TRANSLATION_INPUT_DELAY  <string>              (default: 300)
  GTA_TRANSLATION_LANGUAGES    <string>,[string...]  (default: English;中文;Français;Italiano;日本語;한국어;Deutsch;繁體中文)
  GTA_TRANSLATION_SOURCE       <string>              (default: English)
  GTA_TRANSLATION_TARGET       <string>              (default: 中文)
```

> [!注意]
> 首次运行该项目时，系统会下载所需的模型，此过程可能需要几分钟乃至几十分钟，视网络情况而定。


## 屏幕截图
原文出处: [Luoyang, the peony city that stole a British girl's heart](https://www.ecns.cn/travel/2026-06-02/detail-ihffcfqs3229531.shtml)

![Screenshot](screenshot.png)

## 致谢
- [Kronk](https://github.com/ardanlabs/kronk)
- [Fyne](https://github.com/fyne-io/fyne)
- [Ctrl+Revise](https://github.com/bahelit/ctrl_plus_revise)
- [franslate](https://github.com/sercanarga/franslate)

## 许可
本项目采用 GPL-3.0 许可证发布——详情请参阅 [LICENSE](LICENSE) 文件。
