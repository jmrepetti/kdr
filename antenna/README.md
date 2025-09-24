# Antenna

Pick up os.Signals

USAGE:

```go
func main() {
    ant := antenna.NewAntenna(syscall.SIGINT, syscall.SIGTERM)
    // ....
    // Your code
    // ....
    <-ant.Wait() // Wait for syscall.SIGINT or syscall.SIGTERM.
    // Shutdown/Clean up code
}
```