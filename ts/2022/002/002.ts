declare var require: any

const fs = require('fs');
const readline = require('readline');

async function doMagic(file: string) {
    const fileStream = fs.createReadStream(file);

    const rl = readline.createInterface({
        input: fileStream,
        crlfDelay: Infinity
    });

    let score1: number = 0;
    let score2: number = 0;

    for await (const txt of rl) {
        switch (txt[0]) {
        case 'A':
            switch (txt[2]) {
            case 'X':
                score1++;
                score1 += 3;

                score2 += 3;
                break;
            case 'Y':
                score1 += 2;
                score1 += 6;

                score2 += 3;
                score2++;
                break;
            case 'Z':
                score1 += 3;

                score2 += 6;
                score2 += 2;
                break;
            };
            break;
        case 'B':
            switch (txt[2]) {
            case 'X':
                score1++;

                score2++;
                break;
            case 'Y':
                score1 += 2;
                score1 += 3;

                score2 += 3;
                score2 += 2;
                break;
            case 'Z':
                score1 += 3;
                score1 += 6;

                score2 += 6;
                score2 += 3;
                break;
            };
            break;
        case 'C':
            switch (txt[2]) {
            case 'X':
                score1++;
                score1 += 6;

                score2 += 2;
                break;
            case 'Y':
                score1 += 2;

                score2 += 3;
                score2 += 3;
                break;
            case 'Z':
                score1 += 3;
                score1 += 3;

                score2 += 6;
                score2++;
                break;
            };
            break;
        };
    };

    console.log(score1);
    console.log(score2);
};

doMagic('example.txt');
doMagic('input.txt');