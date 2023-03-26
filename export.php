<?php
$t1 = floor(microtime(true) * 1000);
// Database credentials
$DB_USER = "root";
$DB_PASSWORD = "123456";
$DB_HOST = "localhost";
$DB_NAME = "benchmark_test";

// Create connection
$conn = new mysqli($DB_HOST, $DB_USER, $DB_PASSWORD, $DB_NAME);

// Check connection
if ($conn->connect_error) {
    die("Connection failed: " . $conn->connect_error);
}

$fp = fopen('output.csv', 'a');

fputcsv($fp, array('id', 'qty', 'price', 'total', 'column_0', 'column_1', 'column_2', 'column_3', 'column_4', 'column_5', 'column_6',' column_7', 'column_8', 'column_9', 'column_10'));
$after_id = 0;
$limit = 10000;
while (true) {
    $result = $conn->query("SELECT * FROM products WHERE id > $after_id ORDER BY id ASC LIMIT $limit");
    if ($result->num_rows == 0) {
        break;
    }
    while ($row = $result->fetch_assoc()) {
        fputcsv($fp, array(
            $row['id'],
            $row['qty'],
            $row['price'],
            $row['qty'] * $row['price'],
            $row['column_0'],
            $row['column_1'],
            $row['column_2'],
            $row['column_3'],
            $row['column_4'],
            $row['column_5'],
            $row['column_6'],
            $row['column_7'],
            $row['column_8'],
            $row['column_9'],
            $row['column_10'],
        ));
        $after_id = $row['id'];
    }
} 

fclose($fp);
$conn->close();

$t2 = floor(microtime(true) * 1000);

echo $t2-$t1;