const parse = (coordinates, [depth, hposition, aim]) => {
    if (!coordinates.length)
        return [depth, hposition]

    const [command, number] = coordinates[0]

    switch (command) {
        case 'up':
            return parse(coordinates.slice(1), [depth, hposition, aim - number])

        case 'down':
            return parse(coordinates.slice(1), [depth, hposition, aim + number])

        default:
            return parse(coordinates.slice(1), [depth + (aim * number), hposition + number, aim])
    }
}

export const second = (coordinates) => {
    return parse(coordinates, [0, 0, 0]).reduce((a, b) => a * b, 1)
}