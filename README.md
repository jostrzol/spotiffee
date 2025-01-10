# Spotiffee

☕ for 🎵 Spotify

## Desciption

Blocks GNOME from automatically suspending the system when spotify is playing.

## Install

```sh
go install github.com/jostrzol/spotiffee/cmd/spotiffee@latest
```

*Note*: ensure `$HOME/go/bin` is added `PATH`

For autostart:

1. Copy the contents of [Spotiffee.service](./service/Spotiffee.service) to `$HOME/.config/systemd/user/Spotiffee.service`
2. `systemctl --user daemon-reload`
3. `systemctl --user enable Spotiffee.service`
4. `systemctl --user start Spotiffee.service`

## Build

```sh
go build ./cmd/spotiffee
```
