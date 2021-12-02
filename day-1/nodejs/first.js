import { createReadStream } from 'fs';
import { createInterface } from 'readline';

const fileStream = createReadStream('input.txt');
const rl = createInterface({
    input: fileStream,
    crlfDelay: Infinity
});

async function count(lines) {
    let previous, n = 0

    for await (const line of lines) {
        const number = line | 0;

        if (previous < number) n++
        previous = number
    }

    return n
}

(async () => {
    console.log(await count(rl));
})();