# Goroutine Management with Context in Go

This repository demonstrates how to use **`context`** in Go to start, cancel, and manage multiple goroutines at runtime.  
It’s a simple interactive program that helps you understand how *cancellation* works in Go contexts.

---

## ✨ Features

- Start a counter goroutine that increments every second  
- Dynamically add more counter goroutines  
- Cancel individual goroutines  
- Cancel (clear) all running goroutines at once  
- View the current value of a shared counter

---

## 📂 Project Structure

```
.
├── main.go     # Source code
└── README.md   # Documentation
```

---

## 🚀 How It Works

1. The program starts with a **root context** and launches an initial goroutine (`startCounting`).
2. Each goroutine runs an infinite loop:
   - Increments a shared `counter` every second  
   - Stops when its context is cancelled
3. A global `holder` slice keeps track of all `ContextData` instances (each holding `context.Context` + `CancelFunc`).
4. Through a simple CLI menu, you can:
   - Show the current counter value  
   - Add new goroutines (each has its own cancel function)  
   - Remove a specific goroutine by index  
   - Cancel all goroutines at once

---

## 📜 Menu Options

When you run the app, you’ll see:

| Option | Action |
|--------|---------|
| `1` | Show current counter value |
| `2` | Add a new goroutine |
| `3` | Remove a goroutine (cancel by index) |
| `4` | Clear all goroutines |

---

## ▶️ Usage

### 1️⃣ Run the program

```bash
go run main.go
```

### 2️⃣ Follow the prompts

Example session:

```
Enter number: 1
42
Enter number: 2
Len of holder =>  2
Enter number: 3
0) {context.Background cancelFunc}
1) {context.Background cancelFunc}
Enter number to see the list: 1
Cancelled...
```

---

## 🧠 What You’ll Learn

- How to create and use `context.WithCancel`  
- How to propagate cancellation signals to goroutines  
- How to manage multiple goroutines dynamically  
- How to safely stop goroutines instead of leaving them running forever

> This example is **educational** and focuses on understanding context cancellation, not on production-ready architecture.

---

## ⚠️ Notes & Improvements

- The `counter` is shared between goroutines without synchronization → use `sync.Mutex` or `atomic` if you need thread safety.  
- Goroutines started here run simple loops; in real apps, they’d handle I/O, API calls, or long-running tasks.  
- For large numbers of goroutines, consider using a `WaitGroup` to wait for graceful shutdown.

---

## 📚 References

- [Go official docs: Context package](https://pkg.go.dev/context)  
- [Go Concurrency Patterns: Context](https://blog.golang.org/context)

---

Would you like me to include a small diagram (e.g., ASCII or image) to visualize the relationships between `holder`, contexts, and goroutines?
