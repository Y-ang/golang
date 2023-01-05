golintx三个规则的优化

1. feat: 优化wg.add后紧跟gostmt的误报情况
2. feat: 优化wg的定义在loop内的误报
3. feat: 优化hanging goroutine召回率
 Code difference 工具的使用
二、尝试了几个不同的工具有：go-diff等，还是git命令方便简洁
git diff v1.23.0 v1.24.0 test/e2e/apps/deployment.go
k8s旧版本出现的问题在新版本中仍然出现
