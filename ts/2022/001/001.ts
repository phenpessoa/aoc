declare var require: any

const fs = require('fs');
const readline = require('readline');

async function doMagic(file:string) {
    const fileStream = fs.createReadStream(file);

    const rl = readline.createInterface({
        input: fileStream,
        crlfDelay: Infinity
    });

    let vals: Array<number> = [];
    let cur: number = 0;

    for await (const txt of rl) {
        if (txt.length == 0) {
            vals.push(cur)
            cur = 0
            continue
        }
        cur += +txt;
    }

    if (cur !== 0) {
        vals.push(cur)
    }

    vals.sort((a, b) => {
        return a-b;
    });

    console.log(vals[vals.length-1]);
    console.log(vals[vals.length-1]+vals[vals.length-2]+vals[vals.length-3]);
}

doMagic('example.txt');
doMagic('input.txt');