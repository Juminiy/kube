package docker_file

// Cmd
// CMD ["executable","param1","param2"] (exec form)
// CMD ["param1","param2"] (exec form, as default parameters to ENTRYPOINT)
// CMD command param1 param2 (shell form)
type Cmd struct{ Exe }
