package librarian

var stmtCreateRealtimeDataTable = "CREATE TABLE IF NOT EXISTS RealtimeData_%d (time DATETIME, rank INT, circulating_supply FLOAT, price FLOAT, volume24h FLOAT, market_cap FLOAT, rate1h FLOAT, rate24h FLOAT, rate7d FLOAT);"

var stmtInsertToRealtimeDataTable = "INSERT INTO RealtimeData_%d (time, rank, circulating_supply, price, volume24h, market_cap, rate1h, rate24h, rate7d) VALUES ('%s', %d, %f, %f, %f, %f, %f,%f, %f);"

var stmtCreateCoinsTable = "CREATE TABLE IF NOT EXISTS coins(ID INT, name VARCHAR(50), symbol VARCHAR(25), slug VARCHAR(50));"

var stmtInsertCoinToCoinsTable = "INSERT INTO coins (ID, name, symbol, slug) VALUES (%d, '%s', '%s', '%s');"
