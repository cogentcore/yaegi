package stdlib

// Code generated by 'goexports os'. DO NOT EDIT.

import (
	"os"
	"reflect"
)

func init() {
	Value["os"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Args":              reflect.ValueOf(&os.Args).Elem(),
		"Chdir":             reflect.ValueOf(os.Chdir),
		"Chmod":             reflect.ValueOf(os.Chmod),
		"Chown":             reflect.ValueOf(os.Chown),
		"Chtimes":           reflect.ValueOf(os.Chtimes),
		"Clearenv":          reflect.ValueOf(os.Clearenv),
		"Create":            reflect.ValueOf(os.Create),
		"DevNull":           reflect.ValueOf(os.DevNull),
		"Environ":           reflect.ValueOf(os.Environ),
		"ErrClosed":         reflect.ValueOf(&os.ErrClosed).Elem(),
		"ErrExist":          reflect.ValueOf(&os.ErrExist).Elem(),
		"ErrInvalid":        reflect.ValueOf(&os.ErrInvalid).Elem(),
		"ErrNoDeadline":     reflect.ValueOf(&os.ErrNoDeadline).Elem(),
		"ErrNotExist":       reflect.ValueOf(&os.ErrNotExist).Elem(),
		"ErrPermission":     reflect.ValueOf(&os.ErrPermission).Elem(),
		"Executable":        reflect.ValueOf(os.Executable),
		"Exit":              reflect.ValueOf(os.Exit),
		"Expand":            reflect.ValueOf(os.Expand),
		"ExpandEnv":         reflect.ValueOf(os.ExpandEnv),
		"FindProcess":       reflect.ValueOf(os.FindProcess),
		"Getegid":           reflect.ValueOf(os.Getegid),
		"Getenv":            reflect.ValueOf(os.Getenv),
		"Geteuid":           reflect.ValueOf(os.Geteuid),
		"Getgid":            reflect.ValueOf(os.Getgid),
		"Getgroups":         reflect.ValueOf(os.Getgroups),
		"Getpagesize":       reflect.ValueOf(os.Getpagesize),
		"Getpid":            reflect.ValueOf(os.Getpid),
		"Getppid":           reflect.ValueOf(os.Getppid),
		"Getuid":            reflect.ValueOf(os.Getuid),
		"Getwd":             reflect.ValueOf(os.Getwd),
		"Hostname":          reflect.ValueOf(os.Hostname),
		"Interrupt":         reflect.ValueOf(&os.Interrupt).Elem(),
		"IsExist":           reflect.ValueOf(os.IsExist),
		"IsNotExist":        reflect.ValueOf(os.IsNotExist),
		"IsPathSeparator":   reflect.ValueOf(os.IsPathSeparator),
		"IsPermission":      reflect.ValueOf(os.IsPermission),
		"IsTimeout":         reflect.ValueOf(os.IsTimeout),
		"Kill":              reflect.ValueOf(&os.Kill).Elem(),
		"Lchown":            reflect.ValueOf(os.Lchown),
		"Link":              reflect.ValueOf(os.Link),
		"LookupEnv":         reflect.ValueOf(os.LookupEnv),
		"Lstat":             reflect.ValueOf(os.Lstat),
		"Mkdir":             reflect.ValueOf(os.Mkdir),
		"MkdirAll":          reflect.ValueOf(os.MkdirAll),
		"ModeAppend":        reflect.ValueOf(os.ModeAppend),
		"ModeCharDevice":    reflect.ValueOf(os.ModeCharDevice),
		"ModeDevice":        reflect.ValueOf(os.ModeDevice),
		"ModeDir":           reflect.ValueOf(uint32(os.ModeDir)),
		"ModeExclusive":     reflect.ValueOf(os.ModeExclusive),
		"ModeIrregular":     reflect.ValueOf(os.ModeIrregular),
		"ModeNamedPipe":     reflect.ValueOf(os.ModeNamedPipe),
		"ModePerm":          reflect.ValueOf(os.ModePerm),
		"ModeSetgid":        reflect.ValueOf(os.ModeSetgid),
		"ModeSetuid":        reflect.ValueOf(os.ModeSetuid),
		"ModeSocket":        reflect.ValueOf(os.ModeSocket),
		"ModeSticky":        reflect.ValueOf(os.ModeSticky),
		"ModeSymlink":       reflect.ValueOf(os.ModeSymlink),
		"ModeTemporary":     reflect.ValueOf(os.ModeTemporary),
		"ModeType":          reflect.ValueOf(uint32(os.ModeType)),
		"NewFile":           reflect.ValueOf(os.NewFile),
		"NewSyscallError":   reflect.ValueOf(os.NewSyscallError),
		"O_APPEND":          reflect.ValueOf(os.O_APPEND),
		"O_CREATE":          reflect.ValueOf(os.O_CREATE),
		"O_EXCL":            reflect.ValueOf(os.O_EXCL),
		"O_RDONLY":          reflect.ValueOf(os.O_RDONLY),
		"O_RDWR":            reflect.ValueOf(os.O_RDWR),
		"O_SYNC":            reflect.ValueOf(os.O_SYNC),
		"O_TRUNC":           reflect.ValueOf(os.O_TRUNC),
		"O_WRONLY":          reflect.ValueOf(os.O_WRONLY),
		"Open":              reflect.ValueOf(os.Open),
		"OpenFile":          reflect.ValueOf(os.OpenFile),
		"PathListSeparator": reflect.ValueOf(os.PathListSeparator),
		"PathSeparator":     reflect.ValueOf(os.PathSeparator),
		"Pipe":              reflect.ValueOf(os.Pipe),
		"Readlink":          reflect.ValueOf(os.Readlink),
		"Remove":            reflect.ValueOf(os.Remove),
		"RemoveAll":         reflect.ValueOf(os.RemoveAll),
		"Rename":            reflect.ValueOf(os.Rename),
		"SEEK_CUR":          reflect.ValueOf(os.SEEK_CUR),
		"SEEK_END":          reflect.ValueOf(os.SEEK_END),
		"SEEK_SET":          reflect.ValueOf(os.SEEK_SET),
		"SameFile":          reflect.ValueOf(os.SameFile),
		"Setenv":            reflect.ValueOf(os.Setenv),
		"StartProcess":      reflect.ValueOf(os.StartProcess),
		"Stat":              reflect.ValueOf(os.Stat),
		"Stderr":            reflect.ValueOf(&os.Stderr).Elem(),
		"Stdin":             reflect.ValueOf(&os.Stdin).Elem(),
		"Stdout":            reflect.ValueOf(&os.Stdout).Elem(),
		"Symlink":           reflect.ValueOf(os.Symlink),
		"TempDir":           reflect.ValueOf(os.TempDir),
		"Truncate":          reflect.ValueOf(os.Truncate),
		"Unsetenv":          reflect.ValueOf(os.Unsetenv),
		"UserCacheDir":      reflect.ValueOf(os.UserCacheDir),

		// type definitions
		"File":         reflect.ValueOf((*os.File)(nil)),
		"FileInfo":     reflect.ValueOf((*os.FileInfo)(nil)),
		"FileMode":     reflect.ValueOf((*os.FileMode)(nil)),
		"LinkError":    reflect.ValueOf((*os.LinkError)(nil)),
		"PathError":    reflect.ValueOf((*os.PathError)(nil)),
		"ProcAttr":     reflect.ValueOf((*os.ProcAttr)(nil)),
		"Process":      reflect.ValueOf((*os.Process)(nil)),
		"ProcessState": reflect.ValueOf((*os.ProcessState)(nil)),
		"Signal":       reflect.ValueOf((*os.Signal)(nil)),
		"SyscallError": reflect.ValueOf((*os.SyscallError)(nil)),
	}
}
