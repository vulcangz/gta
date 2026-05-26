# Gemma4-powered Translation Assistant (WIP)

[简体中文](README_zh-CN.md)

GTA is a free, open-source desktop translation assistant tool. It loads and runs the Gemma 4 model server via [Kronk](https://github.com/ardanlabs/kronk), operates in a server-client mode, and supports access and use by users on a local area network. The tool is written in Go, and its user interface utilizes the [Fyne framework](https://github.com/fyne-io/fyne). It supports cross-platform use and offers full hardware acceleration (implemented via Kronk). It runs entirely on your local device (no internet connection is required once the LLM model has been downloaded).

## Why?

1) Professionalism. A long time ago, Google Translate was the most commonly used translation tool. After Gemma 4 was released, I tested its translation capabilities right away, and sure enough, it remained as professional as ever! Unlike in the past, you can now even see the translation’s thought process.
2) Flexible local deployment. Gemma 4 has certain hardware requirements. For SOHO users, you can use a high-spec computer to host a more powerful model—such as the 26B (MoE, with 4B activated) or 31B (dense) version—as a server, while other devices access it via client mode.
Of course, you can also deploy it on your own computer for personal use.
3) Privacy-focused. By deploying large models locally without using Ollama or a browser, you eliminate intermediaries, improve processing efficiency, and protect privacy and data security.
4) Scalability. Future implementations can support other Gemma 4 features such as image text extraction, article writing, or summarization.

## License
This project is licensed under the GPL-3.0 License - see the [LICENSE](LICENSE) file for details.
