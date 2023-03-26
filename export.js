const mysql = require('mysql2/promise');
const createCsvWriter = require('csv-writer').createArrayCsvWriter;

async function main() {
    const t1 = Date.now()
    const csvWriter = createCsvWriter({
        header: ['NAME', 'LANGUAGE'],
        path: 'node_output.csv'
    });
    await csvWriter.writeRecords([
        ['id', 'qty', 'price', 'total', 'column_0', 'column_1', 'column_2', 'column_3', 'column_4', 'column_5', 'column_6',' column_7', 'column_8', 'column_9', 'column_10']
    ])

    const connection = await mysql.createConnection({
        host: 'localhost',
        user: 'root',
        password: '123456',
        database: 'benchmark_test'
    });
    
    
    let after_id = 0;
    const limit = 10000;
    while (true) {
        const [rows] = await connection.execute(`SELECT * FROM products WHERE id > ${after_id} ORDER BY id ASC LIMIT ${limit}`);
        if (!rows.length) break;
        await csvWriter.writeRecords(rows.map(i => ([
            i.id,
            i.qty,
            i.price,
            i.qty * i.price,
            i.column_0,
            i.column_1,
            i.column_2,
            i.column_3,
            i.column_4,
            i.column_5,
            i.column_6,
            i.column_7,
            i.column_8,
            i.column_9,
            i.column_10,
        ])))

        after_id = rows[rows.length - 1].id
    }

    await connection.end();

    const t2 = Date.now();

    console.log(t2-t1);
}

main()