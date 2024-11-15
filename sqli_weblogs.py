import re

sql_injection_patterns = [
    r"SELECT.*FROM",                   # SELECT queries
    r"UNION.*SELECT",                  # UNION-based attacks
    r"INSERT.*INTO",                   # INSERT statements
    r"DELETE.*FROM",                   # DELETE statements
    r"UPDATE.*SET",                    # UPDATE statements
    r"DROP TABLE",                     # DROP TABLE statement
    r"OR 1=1",                         # Authentication bypass attempt
    r"' OR '1'='1",                    # Simple SQL injection
    r"--",                             # Comment symbol in SQL
    r";",                              # End of SQL statement
    r"/\*.*?\*/",                        # SQL inline comments
]

compiledpatterns = [re.compile(pattern, re.IGNORECASE) for pattern in sql_injection_patterns]

def detect_sqli(filename):
    with open(filename, "r") as file:
        for lineno, line in enumerate(file, 1):
            for pattern in compiledpatterns:
                if pattern.search(line):
                    print(f"potential sqli vulnerabilty in {lineno} : {line.strip()}")

filename = "sqli.logs"
detect_sqli(filename=filename)