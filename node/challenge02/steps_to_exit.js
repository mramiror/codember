const fs = require('fs');
const readline = require('readline');

function stepsToExit(nums) {
    let pos = 0;
    let steps = 0;
    let n = nums.length;

    while (pos >= 0 && pos < n) {
        let salto = nums[pos];
        nums[pos] += 1;
        pos += salto;
        steps += 1;
    }

    return steps;
}

async function main() {
    const fileStream = fs.createReadStream('trace.txt', 'utf8');

    const rl = readline.createInterface({
        input: fileStream,
        crlfDelay: Infinity
    });

    let totalSteps = 0;
    let stepsLastLine = 0;

    for await (const line of rl) {
        const l = line.trim();
        if (l === '') continue;

        let nums = l.split(' ').map(Number);

        let steps = stepsToExit(nums);

        totalSteps += steps;
        stepsLastLine = steps;

    }
    console.log(`submit ${totalSteps}-${stepsLastLine}`);
}

main().catch(err => {
    console.error('Error:', err);
    process.exit(1);
});