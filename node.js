//import os from "os"
const os = require('os')


console.log(__dirname)
console.log(__filename)




const currentOs={
    name:os.type(),
    release: os.release(),
    totalMem : os.totalmem(),
    freeMem: os.freemem(),
}
console.log(currentOs)
