package hooks

func initHooks(h *Hooks) {
	h.registerBeforeRequestHook(&MeshDefaultsHook{})
}
