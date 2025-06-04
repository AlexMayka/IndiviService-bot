# Gogram

Gogram is a minimal, extensible Telegram Bot SDK for Go.  
It provides a clean routing system, middleware support, and an in-memory FSM for building conversational bots with ease — no external dependencies, just Go.

---

## 🚀 Features

- 🧠 **FSM support** — per-user state and key-value session memory
- 🔌 **Middleware** — structured and reusable middlewares (like Gin)
- ⚙️ **Routing DSL** — `Group`, `SetState`, `Command`, `Regex`, `Any`, and more
- 📦 **Modular architecture** — split by `core`, `infra`, `runtime`, `internal`
- 🪵 **Structured logging** — powered by `slog.Logger`
- 🔄 **Long polling out of the box**
- 🧪 **Ready-to-use examples** — Echo Bot, Register Flow

---

## 📦 Installation

```bash
go get github.com/AlexMayka/gogram
```

---
## ✍️ Quick Example
```go
r := router.New()

r.Command("/start", func(ctx core.Context) {
	ctx.Send(&commands.SendMessageRequest{Text: "Hello!"})
})

r.Any(func(ctx core.Context) {
	ctx.Send(&commands.SendMessageRequest{Text: "You said: " + ctx.Text()})
})

b := runtime.New(os.Getenv("TELEGRAM_TOKEN"), runtime.WithRouter(r))
b.Run(context.Background())
```
More examples available in examples/

---
## 🧠 FSM Example
```go
r := router.New()
r.Command("/register", func(ctx core.Context) {
	ctx.Send(&commands.SendMessageRequest{Text: "Enter your name:"})
	ctx.FMS().Set(ctx.ChatId(), "awaiting_name")
})

name := r.Group("/name").SetState("awaiting_name")
name.Any(func(ctx core.Context) {
	ctx.FMS().SetParam(ctx.ChatId(), "name", ctx.Text())
	ctx.Send(&commands.SendMessageRequest{Text: "Got it!"})
	ctx.FMS().Set(ctx.ChatId(), "")
})
```
---
## 🛠️ Project Structure
```
core/         → Interfaces and core abstractions (Router, Context, FSM)
infra/        → Implementations: FSM, Polling, Logging, HTTP transport
runtime/      → Main bot runtime & dispatcher
examples/     → Example bots (Echo, Register Flow)
types/        → Telegram API types (Message, User, etc.)
```

---

## 🧪 Run Examples
```bash
TELEGRAM_TOKEN=your_token go run ./examples/echo-bot/main.go
```