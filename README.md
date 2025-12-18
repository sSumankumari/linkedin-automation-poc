# LinkedIn Automation – Browser Stealth POC (Go + Rod)

> ⚠️ **Educational & Demonstration Only**
> This project is a **technical proof-of-concept** created solely for learning and evaluation purposes.
>
> **It does NOT automate LinkedIn actions in production** and must **never** be used on real accounts.
> Automating LinkedIn violates their Terms of Service.

---

## 📌 Project Overview

This repository demonstrates:

* Advanced **browser automation architecture** using **Golang + Rod**
* **Human-like interaction simulation** (mouse, typing, scrolling, timing)
* **Browser fingerprint masking** techniques
* Clean, modular, and maintainable Go code structure
* Safe **demo-only execution** without performing real LinkedIn actions

The tool launches a visible Chrome browser, navigates to LinkedIn login, and **simulates** realistic human behavior using dummy credentials.

---

## 🧠 Implemented Features

### ✅ Core Demonstration Features

| Feature            | Description                        |
| ------------------ | ---------------------------------- |
| Browser Automation | Chrome automation using Rod        |
| Visible Execution  | Non-headless mode for demo         |
| Human-like Mouse   | Smooth, randomized mouse movement  |
| Human-like Typing  | Variable keystroke delay           |
| Random Scrolling   | Natural scrolling behavior         |
| Random Delays      | Think time between actions         |
| Scheduler          | Time-window based execution        |
| Logging            | Structured CLI logs                |
| Config Handling    | `.env`-based configuration         |
| State-safe Demo    | No real LinkedIn actions performed |

---

### 🕵️ Stealth & Anti-Detection (Demo-Safe)

The following **non-abusive** stealth concepts are demonstrated:

* `navigator.webdriver` masking
* Random viewport size
* User-like mouse paths
* Typing rhythm variation
* Action cooldowns
* Randomized delays
* Execution window scheduling

> 🔒 These are **conceptual demonstrations**, not production bypasses.

---

## 📂 Project Structure

```
linkedin-automation-poc/
│
├── cmd/
│   └── main.go                # Application entry point
│
├── internal/
│   ├── auth/                  # Login demo logic
│   ├── browser/               # Browser & launcher setup
│   ├── config/                # Env config loader
│   ├── connect/               # Connection demo service
│   ├── messaging/             # Messaging demo service
│   ├── search/                # Profile search simulation
│   ├── scheduler/             # Time-window control
│   ├── stealth/               # Human behavior simulation
│   ├── storage/               # State placeholder
│   ├── types/                 # Domain models
│   └── logger/                # Structured logging
│
├── .env.example                # Environment template
├── go.mod
├── go.sum
└── README.md
```

---

## ⚙️ System Requirements

| Requirement | Version                   |
| ----------- | ------------------------- |
| Go          | **1.22+**                 |
| OS          | Windows / macOS / Linux   |
| Browser     | Google Chrome (installed) |

> ⚠️ On **Windows**, antivirus may block automation tools.
> This project disables Rod leakless mode safely.

---

## 🧩 Setup Instructions

### 1️⃣ Clone Repository

```bash
git clone <your-repo-url>
cd linkedin-automation-poc
```

---

### 2️⃣ Create Environment File

Copy the example file:

```bash
cp .env.example .env
```

Edit `.env`:

```env
LINKEDIN_EMAIL=dummy@example.com
LINKEDIN_PASSWORD=dummy_password
HEADLESS=false
START_HOUR=0
END_HOUR=23
```

> ℹ️ **Dummy credentials only** — no real account required.

---

### 3️⃣ Install Dependencies

```bash
go mod tidy
```

---

### 4️⃣ Build Project

```bash
go build ./...
```

---

## ▶️ Running the Demo

```bash
go run ./cmd
```

---

## 🧪 What Happens During Execution

When you run the project:

1. Scheduler validates allowed time window
2. Chrome launches visibly
3. Viewport size is randomized
4. Browser fingerprint is masked
5. Mouse moves naturally
6. Page scrolls randomly
7. LinkedIn login page opens
8. Dummy credentials are typed using human-like typing
9. Search, connect, and messaging services **simulate behavior**
10. All steps are logged in CLI

---

## 🖥 Sample CLI Output

```json
{"level":"info","message":"LinkedIn Automation POC started"}
[DEMO] Browser launched with stealth configuration
[DEMO] Fingerprint masking applied
[DEMO] Human-like mouse movement executed
[DEMO] Typing simulation in progress
[DEMO] Profile collection simulated
[DEMO] Connection request simulated
{"level":"info","message":"Execution completed"}
```

---

## 🧾 Conclusion

This project demonstrates a demo-only implementation of browser automation concepts using Golang and Rod, focusing on human-like behavior simulation, stealth awareness, and clean modular architecture.

It is intentionally designed for technical evaluation and learning, showcasing how such systems are structured without performing real-world automation or violating platform policies.
