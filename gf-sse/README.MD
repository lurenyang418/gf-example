# goframe sse

## 后端

通过 `config.yaml` 修改端口

## 前端

VSCode 插件：Live Server 进行预览

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>SSE Demo</title>
</head>
<body>
    哈哈
</body>

<script>
    const eventSource = new EventSource("http://localhost:8000/sse/create")
    eventSource.addEventListener("connected", function (event) {
        const data = JSON.parse(event.data)
        console.log("接收数据", data)
    })
     eventSource.addEventListener("message", function (event) {
        const data = JSON.parse(event.data)
        console.log("接收数据", data)
    })
    eventSource.addEventListener("ai_response", function (event) {
        const data = JSON.parse(event.data)
        console.log("接收数据", data)
    })
    eventSource.onopen = function (event) {
        console.log("连接成功", event.data)
    }
    eventSource.onerror = function (event) {
        console.log("连接失败")
    }

</script>
</html>
```