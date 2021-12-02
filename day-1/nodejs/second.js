import { readFileSync } from 'fs'

const inputs = readFileSync('input.txt', 'utf8')
const numbers = inputs.split('\n').map(Number)

const windows = (arr, size) => {
    if (size > arr.length) {
        return arr;
    }

    let result = [];
    let lastWindow = arr.length - size;

    for (let i = 0; i <= lastWindow; i += 1) {
        result.push(arr.slice(i, i + size));
    }

    return result;
};

const sumArray = (arr) => arr.reduce((acc, cur) => acc + cur, 0)
const magic = (cur, idx, arr) => cur < (arr[idx + 1] ?? Number.MIN_SAFE_INTEGER)

console.log(
    windows(numbers, 3).map(sumArray).filter(magic).length
)