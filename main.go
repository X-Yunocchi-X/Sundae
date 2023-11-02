package sundae

var temp string = `
let temp =5
let a,b = 1,2
let a,b,c : int,str,bool = 1,"2",true
let add = fn(a,b:int,int)int{return a+b}
const d = add(a,b)
const exec= fn(f: fn()) {f()}
let pick= if a>b {a;io.println(a)} else {b;io.println(b)}
match a {
	a.type -> int : io.println(a)
	a.toString() -> "101" : {
		a = a .. " is a string" // concat
		io.println(a)
	}
	default : io.println("default")
//  for id in range (identifier) { expression }(returnValue)
let sum = for i in 0..10 (let a = 0) {
	a += i
}(a)
`
