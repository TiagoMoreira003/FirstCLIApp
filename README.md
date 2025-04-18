# FirstCLIApp

This repository was created to help me understand how to build a CLI application using **Golang** and the **Cobra** framework.

Although it's an educational project, I’ve tried to make it clear and simple for anyone who wants to understand how it works.

## 🔧 Project Overview

The project started as a simple **password generator**.

Later, I decided to expand it a bit: instead of just printing passwords, the app now **saves passwords to files**. This allows for basic storage without implementing user accounts, logins, or complex authentication/isolation mechanisms.

## 📦 Commands

This CLI app currently supports two main commands:

- `generate` – Generates and saves a password
- `list` – Lists saved passwords from file(s)

---

## 🔐 `generate` Command

Generates a password and saves it to a file.

### Available flags:

- `-u` — Include uppercase letters
- `-s` — Include special characters
- `-l <length>` — Length of the password
- `-f <filename>` — File to save the generated password

### Example:

```bash
go run main.go generate -u -s -l 16 -f mypasswords.txt
```
