const escapeStr = "\` \/ \\ \' \"";
const arr = [4, '2'];
const obj = {str: 'Jane', num: 19, bool: false, undef: undefined}
const nested = {
    arr: [4, undefined , '2'],
    obj: { str: 'John', num: 33, bool: true }
};

Object.freeze(nested);
Object.freeze(nested.arr);
Object.freeze(nested.obj);