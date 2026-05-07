# Gopherlock

A secure CLI password manager built with Go to learn Golang fundamentals and practice security concepts.

## Objectives

- Learn Go and understand its module system
- Build a functional CLI application
- Practice cryptographic security concepts

## Features

- Master password with Argon2id key derivation
- AES-GCM encryption for stored passwords
- Secure password input (hidden from terminal)
- JSON-based vault storage

## Installation

```bash
# Clone the repository
git clone <repository-url>
cd gopherlock

# Build the application
go build -o gopherlock

# Run
./gopherlock --help
```

## Commands

### init

Initialize a new vault with a master password. This creates the `vault.json` file which stores all your credentials securely.

```bash
./gopherlock init
```

**What it does:**
- Prompts you to enter a master password
- Prompts you to confirm the master password (must match)
- Generates a random salt (16 bytes)
- Derives a master key using Argon2id (memory: 64MB, iterations: 3, parallelism: 1)
- Hashes the master key with SHA-256 to create a verification hash
- Creates `vault.json` with the salt and check hash

**Security:** The master password is never stored. Only a salt and a hash of the derived key are stored for verification.

---

### set

Add a new password entry to the vault.

```bash
./gopherlock set
```

**What it does:**
- Prompts for your master password (authenticates you)
- Prompts for the account name (e.g., "Gmail", "GitHub", "Wiki")
- Prompts for the username
- Prompts for the password (hidden input)
- Encrypts the password using AES-GCM with the master key
- Saves the entry to the vault

In case the tuple account-username already exists, the user can update the actual password or cancel the command

**Note:** The encryption uses a unique nonce for each password entry.

---

### get

Retrieve a stored password by account name.

```bash
./gopherlock get <account>
```

**Example:**
```bash
./gopherlock get Gmail
```

**What it does:**
- Takes the account name as a command-line argument
- Prompts for your master password (authenticates you)
- Searches for the account in the vault
- Decrypts and displays the username and password

**Error handling:** If the account doesn't exist, displays "There's no password linked to this account".

---

### list

List all stored account names.

```bash
./gopherlock list
```

**What it does:**
- Reads the vault file
- Displays all account names (one per line)

**Note:** Does not require authentication - only displays account names, not sensitive data.

---

## Security

### Key Derivation
- **Algorithm:** Argon2id
- **Parameters:**
  - Memory: 64MB (65536 KB)
  - Iterations: 3
  - Parallelism: 1
  - Salt: 16 bytes
  - Output: 32 bytes

### Password Encryption
- **Algorithm:** AES-256-GCM
- **Nonce:** 12 bytes (unique per password)

### Verification
- Master password verification uses SHA-256 hash of the derived key
- Failed verification causes immediate program exit

## Project Structure

```
gopherlock/
├── main.go              # Entry point
├── cmd/
│   ├── root.go          # Root command configuration
│   ├── init.go          # init command implementation
│   ├── set.go           # set command implementation
│   ├── get.go           # get command implementation
│   └── list.go          # list command implementation
└── internal/
    ├── crypto.go        # Encryption/decryption functions
    ├── login.go         # Master password authentication
    └── storage.go       # Data structures (Vault, Entry)
```

## Dependencies

- **cobra** - CLI framework
- **golang.org/x/crypto** - Argon2id implementation
- **golang.org/x/term** - Secure password input
