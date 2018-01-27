const mysql = require('mysql');
const log4js = require('log4js');
const getopts = require('getopts')
const jayson = require('jayson');
const fs = require('fs');

// args
const options = getopts(process.argv.slice(2), {
  default: {
    network: 'ropsten',
    single: false
  }
})

// files
const lockFile = '/var/tmp/explorer-' + options.network + '.lock'
const lastFile = '/var/tmp/explorer-' + options.network + '.last-block'
const errorLog = '/var/tmp/explorer-' + options.network + '.errors'

// logs
log4js.configure({
  appenders: { explorer: { type: 'file', filename: errorLog } },
  categories: { default: { appenders: ['explorer'], level: 'debug' } }
});
const logger = log4js.getLogger('explorer');

// database
const connection = mysql.createConnection({
  host     : 'database',
  user     : 'admin',
  password : '12dlql*41X',
  database : 'explorer'
});

// json rpc
const url = options.runType === 'mainnet' ? 'https://api.myetherapi.com/eth' : 'https://api.myetherapi.com/rop';
const client = jayson.client.https(url);

// create lock file
if (fs.existsSync(lockFile)) {
  logger.error('detected lockfile at ' + lockFile + ' ... exiting');
  return;
} else {
  fs.writeFile(lockFile, null, { flag: 'wx' }, function (err) {
    if (err) {
      logger.error('error while trying to create lock file');
      return;
    }
  });
}

// current block
if(options.block){
  fs.readFile(lastFile, 'utf8', function (err, text) {
    console.log(text);
  });
}