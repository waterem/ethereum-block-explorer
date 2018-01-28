const fs = require('fs')
const getopts = require('getopts')

const config = require('./lib/config')
const functions = require('./lib/functions')

// args
const options = getopts(process.argv.slice(2), {
  default: {
    network: 'ropsten',
    single: false
  }
})

// init
config.initWithOptions(options)

// create lock file
/*
if (fs.existsSync(global.lockFile)) {
  return global.logger.error('detected lockfile at ' + global.lockFile + ' ... exiting')
} else {
  fs.writeFileSync(global.lockFile, '', function (err) {
    if (err) {
      return global.logger.error('error while trying to create lock file')
    }
  })
}
*/

// create last block file if not exists
if (!fs.existsSync(global.lastFile)) {
  fs.writeFileSync(global.lastFile, '0', function (err) {
    if (err) {
      return global.logger.error('error while trying to create last block file')
    }
  })
}

// current block
var currentBlock
var firstBlock = options.runType === 'mainnet' ? 0 : 0 // TODO change later
if (!options.block) {
  try {
    var text = fs.readFileSync(global.lastFile, 'utf-8')
    currentBlock = text && parseInt(text) >= firstBlock ? parseInt(text) + 1 : firstBlock
  } catch (err) {
    return global.logger.error('error while trying to read last block file' + err)
  }
} else {
  currentBlock = options.block
}
global.logger.debug('current block : ' + currentBlock)

// get the most recent block
const recentBlock = global.web3.eth.blockNumber
global.logger.debug('recentBlock block : ' + recentBlock)

//----------------

var db = require('mysql-promise')();

db.configure({
  host: 'database',
  user: 'admin',
  password: '12dlql*41',
  database: 'explorer'
});

db.query('select * from index_blocks', []).then(function () {
  return db.query('select * from index_blocks');
}).spread(function (rows) {
  console.log('Loook at all the foo');
});

console.log('end')

/*
let database_query = function (query) {
  return new Promise(function (resolve, reject) {
    let conn = mysql.createConnection({
      host: 'database',
      user: 'admin',
      password: '12dlql*41',
      database: 'explorer'
    })
    conn.connect()
    conn.query(query, (err, data) => (err ? reject(err) : resolve(data)))
    conn.end()
  })
}

(async function () {
  try {
    let rows = await database_query('select * from index_blocks')
    rows.forEach((row) => row_view(row))
  } catch (err) {
    //handleError(err)
    console.log(err)
  }
})
*/

//----------------

// loop through the blocks until we are current
/*
while (currentBlock <= recentBlock) {
  global.logger.debug('processing block ' + currentBlock + '...')

  // create block record
  functions.createBlock(currentBlock)

  console.log('0')

  currentBlock++

  break
}
*/

/*
// delete lock file
fs.unlinkSync(global.lockFile, function (err) {
  if (err) {
    return global.logger.error('error while trying to delete lock file')
  }
})
*/