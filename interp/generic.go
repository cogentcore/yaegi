package interp

import (
	"fmt"
	"strings"
	"sync/atomic"
)

// GenericFunc contains the code of a generic function.
// This is used in the `yaegi extract` command to represent generic functions
// instead of the actual value of the function, since you cannot get the
// [reflect.Value] of a generic function. This is then used to interpret the
// function when it is imported in yaegi.
type GenericFunc string

// adot produces an AST dot(1) directed acyclic graph for the given node. For debugging only.
// func (n *node) adot() { n.astDot(dotWriter(n.interp.dotCmd), n.ident) }

// genAST returns a new AST where generic types are replaced by instantiated types.
func genAST(sc *scope, root *node, types []*itype) (*node, bool, error) {
	typeParam := map[string]*node{}
	pindex := 0
	tname := ""
	rtname := ""
	recvrPtr := false
	fixNodes := []*node{}
	var gtree func(*node, *node) (*node, error)
	sname := root.child[0].ident + "["
	if root.kind == funcDecl {
		sname = root.child[1].ident + "["
	}

	// Input type parameters must be resolved prior AST generation, as compilation
	// of generated AST may occur in a different scope.
	for _, t := range types {
		sname += t.id() + ","
	}
	sname = strings.TrimSuffix(sname, ",") + "]"
	if trace {
		tracePrintln(root, "genAST", sname)
	}

	gtree = func(n, anc *node) (*node, error) {
		if trace {
			tracePrintln(n)
		}
		nod := copyNode(n, anc, false)
		switch n.kind {
		case funcDecl, funcType:
			nod.val = nod

		/*  todo: this is not path that gets hit in cfg for fun[int]() -- come back to it
		case callExpr:
			c0 := n.child[0]
			if c0.kind != indexExpr {
				break
			}
			fun := c0
			tracePrintTree(n, n, "generic start", fun)
			lt := []*itype{}
			for _, c := range c0.child[1:] {
				sym, _, _ := sc.lookup(c.ident)
				c.typ = sym.typ
				tracePrintln(c, "type", c.typ, "sym", sym, sym.typ)
				lt = append(lt, sym.typ)
			}
			tracePrintln(c0, "recursive gen:", fun.ident, lt)
			g, found, err := genAST(sc, fun, lt)
			if err != nil {
				tracePrintln(c0, "rgen failed:", found, err)
				break
			}
			n.child[0] = g
			if g.typ == nil {
				tracePrintln(g, "type is nil, setting to:", lt[0])
				g.typ = lt[0] // todo: what is actual type here?
			}
			return g, nil
		*/

		case identExpr:
			// Replace generic type by instantiated one.
			nt, ok := typeParam[n.ident]
			if !ok {
				break
			}
			nod = copyNode(nt, anc, true)
			nod.typ = nt.typ

		case indexExpr:
			tracePrintln(n, n.anc, "in index expr")
			// Catch a possible recursive generic type definition
			if root.kind == typeSpec {
				tracePrintln(root, "root kind == typeSpec")
				if root.child[0].ident == n.child[0].ident {
					tracePrintln(root.child[0], "same ident")
					nod := copyNode(n.child[0], anc, false)
					fixNodes = append(fixNodes, nod)
					return nod, nil
				}
			}
			if len(n.child) == 0 || n.child[0].typ == nil {
				tracePrintln(n, "no c0 with typ")
				break
			}
			t := n.child[0].typ
			for t.cat == linkedT {
				t = t.val
			}
			if t.cat != funcT {
				break
			}
			c1 := n.child[1]
			if !c1.isType(sc) {
				break
			}
			tracePrintln(t.node.anc, "generic genAST indexed funcT")
			g, _, err := genAST(sc, t.node.anc, []*itype{c1.typ})
			if err != nil {
				tracePrintln(g, "error", err)
				return n, err
			}
			// Replace generic func node by instantiated one.
			n.anc.child[childPos(n)] = g
			n.typ = g.typ
			tracePrintln(n, g, "gen type", g.typ, "par", n.anc, fmt.Sprintf("n: %p  g: %p par: %p", n, g, n.anc))
			return n, nil

		case fieldList:
			//  Node is the type parameters list of a generic function.
			if root.kind == funcDecl && n.anc == root.child[2] && childPos(n) == 0 && len(types) > 0 {
				// Fill the types lookup table used for type substitution.
				for _, c := range n.child {
					l := len(c.child) - 1
					for _, cc := range c.child[:l] {
						t, err := nodeType(c.interp, sc, c.child[l])
						if err != nil {
							return nil, err
						}
						if err := checkConstraint(types[pindex], t); err != nil {
							return nil, err
						}
						typeParam[cc.ident] = copyNode(cc, cc.anc, false)
						typeParam[cc.ident].ident = types[pindex].id()
						typeParam[cc.ident].typ = types[pindex]
						pindex++
					}
				}
				// Skip type parameters specification, so generated func doesn't look generic.
				return nod, nil
			}

			// Node is the receiver of a generic method.
			if root.kind == funcDecl && n.anc == root && childPos(n) == 0 && len(n.child) > 0 {
				rtn := n.child[0].child[1]
				// Method receiver is a generic type if it takes some type parameters.
				if rtn.kind == indexExpr || rtn.kind == indexListExpr || (rtn.kind == starExpr && (rtn.child[0].kind == indexExpr || rtn.child[0].kind == indexListExpr)) {
					if rtn.kind == starExpr {
						// Method receiver is a pointer on a generic type.
						rtn = rtn.child[0]
						recvrPtr = true
					}
					rtname = rtn.child[0].ident + "["
					for _, cc := range rtn.child[1:] {
						if pindex >= len(types) {
							return nil, cc.cfgErrorf("undefined type for %s", cc.ident)
						}
						it := types[pindex]
						typeParam[cc.ident] = copyNode(cc, cc.anc, false)
						typeParam[cc.ident].ident = it.id()
						typeParam[cc.ident].typ = it
						rtname += it.id() + ","
						pindex++
					}
					rtname = strings.TrimSuffix(rtname, ",") + "]"
				}
			}

			// Node is the type parameters list of a generic type.
			if root.kind == typeSpec && n.anc == root && childPos(n) == 1 {
				// Fill the types lookup table used for type substitution.
				tname = n.anc.child[0].ident + "["
				for _, c := range n.child {
					l := len(c.child) - 1
					for _, cc := range c.child[:l] {
						if pindex >= len(types) {
							return nil, cc.cfgErrorf("undefined type for %s", cc.ident)
						}
						it := types[pindex]
						t, err := nodeType(c.interp, sc, c.child[l])
						if err != nil {
							return nil, err
						}
						if err := checkConstraint(types[pindex], t); err != nil {
							return nil, err
						}
						typeParam[cc.ident] = copyNode(cc, cc.anc, false)
						typeParam[cc.ident].ident = it.id()
						typeParam[cc.ident].typ = it
						tname += it.id() + ","
						pindex++
					}
				}
				tname = strings.TrimSuffix(tname, ",") + "]"
				return nod, nil
			}
		}

		for _, c := range n.child {
			gn, err := gtree(c, nod)
			if err != nil {
				return nil, err
			}
			nod.child = append(nod.child, gn)
		}
		return nod, nil
	}

	if nod, found := root.interp.generic[sname]; found {
		if trace {
			tracePrintln(root, "found compiled version")
		}
		return nod, true, nil
	}

	r, err := gtree(root, root.anc)
	if err != nil {
		return nil, false, err
	}
	root.interp.generic[sname] = r
	r.param = append(r.param, types...)
	if tname != "" {
		for _, nod := range fixNodes {
			nod.ident = tname
		}
		r.child[0].ident = tname
	}
	if rtname != "" {
		// Replace method receiver type by synthetized ident.
		nod := r.child[0].child[0].child[1]
		if recvrPtr {
			nod = nod.child[0]
		}
		nod.kind = identExpr
		nod.ident = rtname
		nod.child = nil
	}
	// r.adot() // Used for debugging only.
	if trace {
		tracePrintln(r, "genAST", sname, "done")
	}
	return r, false, nil
}

func copyNode(n, anc *node, recursive bool) *node {
	var i interface{}
	nindex := atomic.AddInt64(&n.interp.nindex, 1)
	nod := &node{
		debug:  n.debug,
		anc:    anc,
		interp: n.interp,
		index:  nindex,
		level:  n.level,
		nleft:  n.nleft,
		nright: n.nright,
		kind:   n.kind,
		pos:    n.pos,
		action: n.action,
		gen:    n.gen,
		val:    &i,
		rval:   n.rval,
		ident:  n.ident,
		meta:   n.meta,
	}
	nod.start = nod
	if recursive {
		for _, c := range n.child {
			nod.child = append(nod.child, copyNode(c, nod, true))
		}
	}
	return nod
}

func inferTypesFromCall(sc *scope, fun *node, args []*node) ([]*itype, error) {
	ftn := fun.typ.node
	// Fill the map of parameter types, indexed by type param ident.
	paramTypes := map[string]*itype{}
	for _, c := range ftn.child[0].child {
		typ, err := nodeType(fun.interp, sc, c.lastChild())
		if err != nil {
			return nil, err
		}
		for _, cc := range c.child[:len(c.child)-1] {
			paramTypes[cc.ident] = typ
		}
	}

	var inferTypes func(*itype, *itype) ([]*itype, error)
	inferTypes = func(param, input *itype) ([]*itype, error) {
		switch param.cat {
		case chanT, ptrT, sliceT:
			return inferTypes(param.val, input.val)

		case mapT:
			k, err := inferTypes(param.key, input.key)
			if err != nil {
				return nil, err
			}
			v, err := inferTypes(param.val, input.val)
			if err != nil {
				return nil, err
			}
			return append(k, v...), nil

		case structT:
			lt := []*itype{}
			for i, f := range param.field {
				nl, err := inferTypes(f.typ, input.field[i].typ)
				if err != nil {
					return nil, err
				}
				lt = append(lt, nl...)
			}
			return lt, nil

		case funcT:
			lt := []*itype{}
			for i, t := range param.arg {
				if i >= len(input.arg) {
					break
				}
				nl, err := inferTypes(t, input.arg[i])
				if err != nil {
					return nil, err
				}
				lt = append(lt, nl...)
			}
			for i, t := range param.ret {
				if i >= len(input.ret) {
					break
				}
				nl, err := inferTypes(t, input.ret[i])
				if err != nil {
					return nil, err
				}
				lt = append(lt, nl...)
			}
			return lt, nil

		case nilT:
			if paramTypes[param.name] != nil {
				return []*itype{input}, nil
			}

		case genericT:
			return []*itype{input}, nil
		}
		return nil, nil
	}

	types := []*itype{}
	for i, c := range ftn.child[1].child {
		typ, err := nodeType(fun.interp, sc, c.lastChild())
		if err != nil {
			return nil, err
		}
		if i >= len(args) {
			fmt.Println("generic functions do not yet support variadic args")
			return nil, nil
		}
		lt, err := inferTypes(typ, args[i].typ)
		if err != nil {
			return nil, err
		}
		types = append(types, lt...)
	}

	return types, nil
}

func checkConstraint(it, ct *itype) error {
	if len(ct.constraint) == 0 && len(ct.ulconstraint) == 0 {
		return nil
	}
	for _, c := range ct.constraint {
		if it.equals(c) || it.matchDefault(c) {
			return nil
		}
	}
	for _, c := range ct.ulconstraint {
		if it.underlying().equals(c) || it.matchDefault(c) {
			return nil
		}
	}
	return it.node.cfgErrorf("%s does not implement %s", it.id(), ct.id())
}
