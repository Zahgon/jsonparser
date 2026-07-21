//go:build (!appengine && !appenginevm && ignore) || !tinygo
// +build !appengine,!appenginevm,ignore !tinygo

package jsonparser

func equalStr(b *[]byte, s string) bool { _ = "STUB: not implemented"; return false }

func parseFloat(b *[]byte) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func bytesToString(b *[]byte) string { _ = "STUB: not implemented"; return "" }

func StringToBytes(s string) []byte { _ = "STUB: not implemented"; return nil }
