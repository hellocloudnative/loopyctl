# `loopyctl` 命令行工具使用说明
`loopyctl` 是loopy命令行工具。它提供了 `setup` 和 `reset` 两个主要命令，用于配置和重置环境。以下是如何使用这些命令的详细说明。
## `loopyctl setup` 命令
`loopyctl setup` 命令用于设置环境。可以选择在单节点或多节点模式下运行。
### 使用方法
```sh
loopyctl setup [选项]
```
### 选项
- `--single`: 在单节点模式下运行设置。
- `--multi`: 在多节点模式下运行设置（即将支持）。
### 示例
#### 在单节点模式下设置环境
```sh
loopyctl setup --single
```
这将执行单节点模式的设置过程。
#### 在多节点模式下设置环境（即将支持）
```sh
loopyctl setup --multi
```
这将执行多节点模式的设置过程（即将支持）。
#### 显示帮助信息
```sh
loopyctl setup --help
```
这将显示 `setup` 命令的帮助信息。
## `loopyctl reset` 命令
`loopyctl reset` 命令用于重置环境。可以选择在单节点或多节点模式下运行。
### 使用方法
```sh
loopyctl reset [选项]
```
### 选项
- `--single`: 在单节点模式下运行重置。
- `--multi`: 在多节点模式下运行重置（即将支持）。
### 示例
#### 在单节点模式下重置环境
```sh
loopyctl reset --single
```
这将执行单节点模式的重置过程。
#### 在多节点模式下重置环境（即将支持）
```sh
loopyctl reset --multi
```
这将执行多节点模式的重置过程（即将支持）。
#### 显示帮助信息
```sh
loopyctl reset --help
```
这将显示 `reset` 命令的帮助信息。
以上是 `loopyctl` 命令行工具的基本使用说明，包括如何设置和重置环境。使用这些命令可以帮助您有效地管理环境配置。未来的更新将包括对多节点安装和重置的支持。
