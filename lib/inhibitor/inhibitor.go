package inhibitor

import (
	"github.com/godbus/dbus/v5"
	"github.com/jostrzol/spotiffee/lib/consts"
)

const (
	SesMgrName      = "org.gnome.SessionManager"
	SesMgrPath      = "/org/gnome/SessionManager"
	SesMgrInhibit   = "org.gnome.SessionManager.Inhibit"
	SesMgrUninhibit = "org.gnome.SessionManager.Uninhibit"
)

type InhibitFlags uint32

const (
	// Inhibit logging out
	InhibitFlagLogOut InhibitFlags = 1 << iota
	// Inhibit user switching
	InhibitFlagSwitchUser
	// Inhibit suspending the session or computer
	InhibitFlagSuspend
	// Inhibit the session being marked as idle
	InhibitFlagMarkIdle
	// Inhibit auto-mounting removable media for the session
	InhibitFlagAutoMount
)

type Inhibitor struct {
	conn   *dbus.Conn
	sesmgr dbus.BusObject
	cookie uint32
}

func New() (*Inhibitor, error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}

	sesmgr := conn.Object(SesMgrName, SesMgrPath)
	inhibitor := &Inhibitor{conn: conn, sesmgr: sesmgr}
	return inhibitor, err
}

func (i *Inhibitor) Inhibit(reason string) error {
	if i.IsInhibited() {
		return nil
	}

	call := i.sesmgr.Call(
		SesMgrInhibit,
		0,
		consts.MyAppId,
		uint32(0),
		reason,
		InhibitFlagSuspend,
	)
	if call.Err != nil {
		return call.Err
	}

	i.cookie = (call.Body[0]).(uint32)
	return nil
}

func (i *Inhibitor) IsInhibited() bool {
	return i.cookie != 0
}

func (i *Inhibitor) Uninhibit() error {
	if !i.IsInhibited() {
		return nil
	}

	call := i.sesmgr.Call(
		SesMgrUninhibit,
		0,
		i.cookie,
	)
	if call.Err != nil {
		return call.Err
	}

	i.cookie = 0
	return nil
}

func (i *Inhibitor) Close() error {
	return i.conn.Close()
}
