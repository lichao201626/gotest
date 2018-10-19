
function* fib(max){
    var
        a=0,
        b=1,
        n=0;
    while(n < max){
        yield a;
        [a, b] = [b, a+b];
        n++;
    }
    return;
}

var f = fib(5);
console.log(f);
f.next();
console.log(f.next().value);
console.log(f.next().done);
f.next();
f.next();
f.next();
f.next();
f.next();
f.next();

for (x of fib(10)) {
    console.log(x);
}