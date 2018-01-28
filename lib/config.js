const log4js = require('log4js')
const mysql = require('mysql2')
const Web3 = require('web3')

module.exports = {
  initWithOptions: function (options) {

    // files
    global.lockFile = '/var/tmp/explorer-' + options.network + '.lock'
    global.lastFile = '/var/tmp/explorer-' + options.network + '.last-block'
    global.errorLog = '/var/tmp/explorer-' + options.network + '.errors'

    // logs
    log4js.configure({
      appenders: {explorer: {type: 'file', filename: errorLog}},
      categories: {default: {appenders: ['explorer'], level: 'debug'}}
    })
    global.logger = log4js.getLogger('explorer')

    // database
    global.connection = mysql.createConnection({
      host: 'database',
      user: 'admin',
      password: '12dlql*41',
      database: 'explorer'
    })

    // web3
    const url = options.runType === 'mainnet' ? 'https://api.myetherapi.com/eth' : 'https://api.myetherapi.com/eth' // TODO change later
    global.web3 = new Web3(new Web3.providers.HttpProvider(url))
  }
}