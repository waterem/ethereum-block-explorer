module.exports = {
  createBlock: function (blockIndex) {

    /*
    const block = await global.web3.eth.getBlock(blockIndex)
    const a = this.createTransaction(block.hash)
    console.log('complete!')
    */
    /*
    const block = global.web3.eth.getBlock(blockIndex)
    var p = new Promise(function (res) { res() })
    p = p.then(this.createTransaction(block.hash))
    p = p.then(this.createTransaction(block.parentHash))
    p.then(function () {
      console.log('complete!')
    })
    */
  },
  createTransaction: function (block_hash) {
    /*
    return new Promise(function (res, rej) {
      const sql = 'SELECT id FROM index_blocks WHERE hash=\'' + block_hash + '\' LIMIT 1'
      global.connection.query(sql, function (e, r, f) {
        console.log('r: ' + r)
      })
    })
    */
  }
}