const path = require("path")

console.log(path.sep)

const filepath = path.join('/schemas','githubActions.json')

console.log(filepath)

const base = path.basename(filepath)
console.log(base)

const absolute = path.resolve(__dirname,"schemas","githubActions.json")
console.log(absolute)