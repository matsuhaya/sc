# sc

A simple CLI tool to remember your VS Code keyboard shortcuts.

## Installation

```bash
go install github.com/matsuhaya/sc@latest
```

## Usage

### Add a shortcut

```bash
sc new "Cmd+Shift+P" "Command Palette"
```

### List shortcuts

```bash
sc
```

Output:

```
Cmd+Shift+P  Command Palette
Cmd+P        Quick Open
Cmd+Shift+F  Search All Files
```

### Edit or delete shortcuts

```bash
sc edit
```

Opens the config file in your editor (`$EDITOR` or `vi`).

## Commands

| Command | Description |
|---------|-------------|
| `sc` | List all shortcuts |
| `sc list` | List all shortcuts |
| `sc new <key> <description>` | Add a new shortcut |
| `sc edit` | Open config file in editor |
| `sc help` | Show help |
| `sc version` | Show version |

## Config

`~/.config/sc/shortcuts.yaml`

```yaml
shortcuts:
  - key: "Cmd+Shift+P"
    description: "Command Palette"
  - key: "Cmd+P"
    description: "Quick Open"
```

## License

MIT
