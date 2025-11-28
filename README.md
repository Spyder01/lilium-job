# **Lilium Job Runner Module**

A lightweight and production-ready **background job runner** for the **Lilium Go framework**.  
Run scheduled or one-time tasks with retries, backoff, timeouts, and concurrency limits — all configurable via `lilium.yaml`.

> 🚀 Plug-and-play background job execution for Lilium apps.

---

## ✨ Features

- ⏱️ Interval-based scheduling (`5s`, `1m`, etc.)
- 🔁 One-time or repeating jobs
- 🧵 Max concurrency per job
- ⛑️ Automatic retries with exponential backoff
- ⌛ Per-task timeout support
- 🎛️ Fully configurable via `lilium.yaml`
- 🧩 Strongly typed API & task injection in code
- 🧹 Graceful shutdown with context cancellation

---

## 📦 Installation

```bash
go get github.com/spyder01/lilium-job
````

---

## ⚙️ Configuration Example (`lilium.yaml`)

```yaml
name: "LiliumTestApp"

server:
  port: 8080

logger:
  toStdout: true
  debugEnabled: true

jobs:
  - name: "test-job"
    interval: 5s
    repeat: true
    description: "Testing background job"
    retries: 2
    enabled: true
    max_concurrency: 1
    timeout: 10s
```

---

## 📌 Job Configuration Explained

Each job under `jobs:` defines a background task.

| Key                 | Type     |  Required  | Default | Description                                                        |
| ------------------- | -------- | :--------: | ------- | ------------------------------------------------------------------ |
| **name**            | string   |      ✅     | —       | Unique identifier for the job — must match `RegisterTask()` name   |
| **enabled**         | bool     |      ❌     | `true`  | Toggle job ON/OFF without code changes                             |
| **interval**        | duration | ⚠️ Usually | none    | Time between runs (`5s`, `1m`, etc.). Ignored when `repeat: false` |
| **repeat**          | bool     |      ❌     | `true`  | Whether job runs multiple times. `false` = run once then stop      |
| **description**     | string   |      ❌     | —       | Logging/monitoring helper                                          |
| **retries**         | int      |      ❌     | `0`     | Retry after failure                                                |
| **initial_backoff** | duration |      ❌     | `1s`    | First delay before retry                                           |
| **max_backoff**     | duration |      ❌     | `30s`   | Upper limit for exponential retry delay                            |
| **timeout**         | duration |      ❌     | none    | Force stop a run if it exceeds duration                            |
| **max_concurrency** | int      |      ❌     | `1`     | Max parallel instances of same job                                 |

---

### 🔁 Job Types

| Intent                     | Suggested Config            | Behavior                    |
| -------------------------- | --------------------------- | --------------------------- |
| Run only once at startup   | `repeat: false`             | Executes once → stops       |
| Scheduled recurring task   | `repeat: true` + `interval` | Executes indefinitely       |
| Flaky external call        | `retries > 0`               | Keeps retrying with backoff |
| High-throughput async work | `max_concurrency > 1`       | Parallel execution          |

---

## 🧩 Usage Example

```go
package main

import (
	"context"

	lilium "github.com/spyder01/lilium-go"
	liliumjob "github.com/spyder01/lilium-job"
)

func main() {
	// Load base Lilium config
	cfg := lilium.LoadConfig("lilium.yaml")
	app := lilium.New(cfg, context.Background())
	router := lilium.NewRouter(app.Context)

	// Load job configurations
	jobCfg, err := liliumjob.LoadLiliumJobsConfig("lilium.yaml")
	if err != nil {
		panic(err)
	}

	// Initialize job module
	module := liliumjob.New(jobCfg)

	// Register job function
	module.RegisterTask("test-job", func(ctx *lilium.AppContext) error {
		ctx.GetLogger().Info("Hello from job")
		return nil
	})

	// Attach module
	app.UseModule(module)

	// Start server + jobs
	app.Start(router)
}
```

---

## 🧪 Runtime Logging

Initialization:

```
Initializing background jobs module...
Job registered: name=test-job interval=5s repeat=true
```

Execution:

```
Job run begin: test-job (attempt 1)
Job success: test-job (120ms)
```

---

## 📞 Support

Issues and feature requests are welcome!
👉 If you like it, please ⭐ the repo!

---

## 📜 License

MIT License © [Spyder01](https://github.com/Spyder01)
