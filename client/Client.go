//go:generate goversioninfo -icon=client.ico
package main

import (
	"dongpo_proxy/client/view"
)

func main() {
	manager := view.UIManager{}
	manager.CreatePanel()
}
