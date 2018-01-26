var mysql = require('mysql');

var connection = mysql.createConnection({
  host     : 'database',
  user     : 'admin',
  password : '12dlql*41',
  database : 'explorer'
});

connection.connect(function(err) {
  if (err) {
    console.error('error connecting: ' + err.stack);
    return;
  }
  console.log('connected as id ' + connection.threadId);
});