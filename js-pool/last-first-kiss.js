function first(arg) {
    return arg[0];
}

function last(arg) {
    return arg[arg.length - 1];
}

function kiss(arg) {
    const arr = [last(arg), first(arg)];
    return arr;
}