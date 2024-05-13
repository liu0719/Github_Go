import sqlite3

class ShoppingDecisionDatabase:
    def __init__(self):
        self.conn = sqlite3.connect('shopping_decision_data.db')
        self.cursor = self.conn.cursor()
        self.create_table()

    def create_table(self):
        create_table_sql = '''
        CREATE TABLE IF NOT EXISTS user_data (
            用户ID INTEGER PRIMARY KEY AUTOINCREMENT,
            年龄 INT,
            性别 CHAR,
            购买历史 TEXT,
            是否购买 INT
        );
        '''
        self.cursor.execute(create_table_sql)
        self.conn.commit()

    def insert_data(self, age, gender, purchase_history, decision):
        insert_sql = '''
        INSERT INTO user_data(年龄, 性别, 购买历史, 是否购买)
        VALUES (?, ?, ?, ?)
        '''
        self.cursor.execute(insert_sql, (age, gender, purchase_history, decision))
        self.conn.commit()

    def get_data(self):
        select_sql = '''
        SELECT * FROM user_data
        '''
        result = self.cursor.execute(select_sql).fetchall()
        return result
