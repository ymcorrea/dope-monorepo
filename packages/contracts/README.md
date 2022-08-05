# Contracts Readme

## Getting started

### ⛔️ WARNING ⛔️

_The author (@faces-of-eth) is a solidity noob and trying to understand how to work with these files. These are one person's clumsy notes. This file would probably be better written by the person who wrote most of this code (@tarrencev)._

---

### 1. Install Foundry

[Foundry](https://book.getfoundry.sh/) is a faster Rust-based alternative to Hardhat

### 2. Install Dependencies

Basic stuff, run `yarn`

### 3. Working with git submodules

Foundry uses local files to specify dependencies. These are located usually in the `[folder]/src/lib` directory. I've tried to add missing ones I found using [the information in this stackoverflow post](https://stackoverflow.com/questions/12898278/issue-with-adding-common-code-as-git-submodule-already-exists-in-the-index).

```sh
# Load submodules
git submodule update --init --recursive

# Pull with submodules
git pull --recurse-submodules

# Add a submodule to existing repo
git submodule add [URL to git repo]
git submodule init
```

---


## Working with contracts in force

### Build contracts

```sh
forge build
```

### Run tests

```sh
forge test --vvv
```

## Adding a new contract
