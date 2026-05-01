# Email Dispatcher

A small Go program that reads recipients from a CSV file, renders a per-recipient message from an HTML/text template, and sends mail concurrently over SMTP using a fixed pool of worker goroutines.

## Requirements

- [Go](https://go.dev/dl/) **1.24.5** or compatible (see `go.mod`).

## Quick start

From the project root:

```bash
go run .
```

The program expects `emails.csv` and `email.tmpl` in the **current working directory** (paths are relative to where you run the binary).

## SMTP

Outgoing mail is sent via SMTP. The server address is configured in `consumer.go` (defaults to `localhost:1025` with no auth), which pairs well with local mail catchers such as [MailHog](https://github.com/mailhog/MailHog) or [Mailpit](https://github.com/axllent/mailpit) for development.

Start your catcher on the same host/port, then run the dispatcher. The **From** address is also set in `consumer.go`; adjust it for your environment.

## Input: `emails.csv`

The first row must be a header. Each following row is one recipient:

| Column | Description   |
|--------|---------------|
| 1      | Display name |
| 2      | Email address |

Example:

```csv
Name,Email
Jane Doe,jane@example.com
```

## Template: `email.tmpl`

[`html/template`](https://pkg.go.dev/html/template) is used (HTML-aware escaping). Each recipient exposes:

- `.Name` — from the first CSV column  
- `.Email` — from the second CSV column  

The default template includes `To:` and `Subject:` lines suitable for simple SMTP payloads. Edit `email.tmpl` to change copy or structure.

## How it works

1. **Producer** — `loadRecipient` reads the CSV and sends `Recipient` values on a channel, then closes the channel.  
2. **Consumers** — Several `emailWorker` goroutines read from the channel, render `email.tmpl`, and call `smtp.SendMail` until the channel is drained.  
3. **Synchronization** — `sync.WaitGroup` waits for all workers to finish.

Worker count is set in `main.go` (`workerCount`).

## Project layout

| File           | Role                                      |
|----------------|-------------------------------------------|
| `main.go`      | Entry point, channel, workers, template execution helper |
| `producer.go`  | CSV loading into the recipient channel    |
| `consumer.go`  | SMTP workers                              |
| `email.tmpl`   | Message body template                     |
| `emails.csv`   | Recipient list                            |

## License

Add a license file if you plan to distribute this repository.
