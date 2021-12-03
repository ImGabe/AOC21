const parse = (coordinates, [depth, hposition]) => {
    if (!coordinates.length)
        return [depth, hposition]

    const [command, number] = coordinates[0]

    switch (command) {
        case 'up':
            return parse(coordinates.slice(1), [depth - number, hposition])

        case 'down':
            return parse(coordinates.slice(1), [depth + number, hposition])

        default:
            return parse(coordinates.slice(1), [depth, hposition + number])
    }
}

export const first = (coordinates) => {
    return parse(coordinates, [0, 0]).reduce((a, b) => a * b, 1)
}