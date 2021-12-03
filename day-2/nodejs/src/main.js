import { equal } from 'assert'
import { readFileSync } from 'fs'

import { first } from "./first.js"
import { second } from "./second.js"

const inputs = readFileSync('input.txt', 'utf8')
const coordinates = inputs.split('\n').map(coordinate => {
    const [command, number] = coordinate.split(' ')
    return [command, number | 0]
})

equal(first(coordinates), 1989265, "os valores devem ser iguais")
equal(second(coordinates), 2089174012, "os valores devem ser iguais")