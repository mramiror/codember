const fs = require('fs');
const readline = require('readline');

function isValid(pwd) {
    let state = 'start';
    let lastDigit = '0';
    let lastChar = 'a';

    for (const c of pwd) {
        if (c >= '0' && c <= '9') {
            if (state === 'char') {
                return false;
            }
            state = 'digit';
            if (c < lastDigit) {
                return false;
            }
            lastDigit = c;
        }
        else if (c >= 'a' && c <= 'z') {
            state = 'char';
            if (c < lastChar) {
                return false;
            }
            lastChar = c;
        }
        else {
            return false;
       }
    }
    return true;
}

async function main() {
    const fileStream = fs.createReadStream('log.txt', 'utf8');

    const rl = readline.createInterface({
        input: fileStream,
        crlfDelay: Infinity
    });

    let valid = 0;
    let invalid = 0;

    for await (const line of rl) {
        const pwd = line.trim();
        if (pwd === '') continue;

        if (isValid(pwd)) valid++;
        else invalid++;
    }
    console.log(`submit ${valid}true${invalid}false`);
}

main().catch(err => {
    console.error('Error:', err);
    process.exit(1);
});