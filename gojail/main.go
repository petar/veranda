package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func usage() {
	fatal("gojail push|pop [push_index]")
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		peek()
		return
	}
	switch flag.Arg(0) {
	case "push":
		var gopath string
		r := roots()
		switch len(r) {
		case 0:
			fatal("not within a gojail")
		case 1:
			gopath = r[0]
		default:
			if flag.NArg() != 2 {
				println("which gojail?\n")
				fatalChoice(r)
			}
			i, err := strconv.Atoi(flag.Arg(1))
			if err != nil {
				println("use an integral 0-based index\n")
				fatalChoice(r)
			}
			if i >= len(r) {
				println("index too big\n")
				fatalChoice(r)
			}
			gopath = r[i]
		}
		push(gopath)
	case "pop":
		pop()
	default:
		usage()
	}
}

func peek() {
	for {
		x := peekvar("GOPATH")
		if x == "" {
			break
		}
		y := peekvar("PATH")
		if y == "" {
			break
		}
		if !strings.HasPrefix(y, x) {
			break
		}
		fmt.Println(x, " • ", y)
	}
}

var osenv = make(map[string][]string)

func peekvar(k string) (chopped string) {
	v, ok := osenv[k]
	if !ok {
		v = filepath.SplitList(os.Getenv(k))
		osenv[k] = v
	}
	if len(v) == 0 {
		return ""
	}
	chopped, v = v[len(v)-1], v[:len(v)-1]
	osenv[k] = v
	return chopped
}

func pop() {
	x, x0 := popvar("GOPATH")
	if x0 == "" {
		fatal("nothing on GOPATH stack")
	}
	y, y0 := popvar("PATH")
	if y0 == "" {
		fatal("nothing on PATH stack")
	}
	if !strings.HasPrefix(y0, x0) {
		fatal("PATH and GOPATH don't match")
	}
	fmt.Println(x)
	fmt.Println(y)
}

func popvar(k string) (cmd, chopped string) {
	gp := filepath.SplitList(os.Getenv(k))
	if len(gp) == 0 {
		return "", ""
	}
	println("removing", gp[len(gp)-1], "from", k)
	chopped = gp[len(gp)-1]
	gp = gp[:len(gp)-1]
	return fmt.Sprintf("export %s=%s", k, join(gp...)), chopped
}

func push(v string) {
	pushvar("GOPATH", v)
	pushvar("PATH", path.Join(v, "bin"))
}

func pushvar(k, v string) {
	println("adding", v, "to", k)
	gp := filepath.SplitList(os.Getenv(k))
	gp = append(gp, v)
	fmt.Printf("export %s=%s\n", k, join(gp...))
}

func join(gg ...string) string {
	var w bytes.Buffer
	for i, g := range gg {
		if i > 0 {
			w.WriteRune(filepath.ListSeparator)
		}
		w.WriteString(g)
	}
	return string(w.Bytes())
}

func fatalChoice(r []string) {
	for i, j := range r {
		fmt.Fprintf(os.Stderr, "%d: %s\n", i, j)
	}
	os.Exit(1)
}

func roots() []string {
	var r []string
	wd, err := os.Getwd()
	pie(err)	
	p := strings.Split(path.Clean(wd), "/")
	if len(p) == 0 || p[0] != "" {
		fatal("working directory not absolute")
	}
	p = p[1:]
	for i, _ := range p {
		g := path.Join(p[:len(p)-i]...)
		g = "/" + g
		if isGoJail(g) {
			r = append(r, g)
		}
	}
	return r
}

func isGoJail(p string) bool {
	_, err := os.Stat(path.Join(p, ".gojail"))
	return err == nil
}

func pie(e error) {
	if e == nil {
		return
	}
	fatal(e)
}

func fatal(v ...interface{}) {
	fmt.Fprintln(os.Stderr, v...)
	os.Exit(1)
}
